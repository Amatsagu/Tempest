package tempest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tempest/misc"
	"time"
)

type Rest struct {
	Token                  string           `binding:"required" example:"Bot XYZABCEWQ"` // Discord Bot/App token. Remember to add "Bot" prefix.
	MaxRequestsBeforeSweep uint8            `default:"100"`
	GlobalRequestLimit     uint8            `default:"50"`
	globalRequests         uint8            `default:"0"`
	requestsSinceSweep     uint8            `default:"0"`
	lockedTo               int64            `default:"0"` // Timestamp (in ms) to when it's locked, 0 means there's no lock.
	locks                  map[string]int64 `default:"{}"`
	fails                  uint8            `default:"0"` // If request failed, try again up to 3 times (delay 250/500/750ms) - after 3rd failed attempt => panic
}

type rateLimitError struct {
	Global     bool    `json:"global"`
	Message    string  `json:"message"`
	RetryAfter float32 `json:"retry_after"`
}

func (rest Rest) Request(method string, route string, jsonPayload interface{}) []byte {
	rest.globalRequests++
	rest.requestsSinceSweep++

	if rest.locks == nil {
		rest.locks = make(map[string]int64, rest.MaxRequestsBeforeSweep)
	}

	now := time.Now().Unix()
	var offset uint8 = 0
	var req *http.Request

	if rest.globalRequests == rest.GlobalRequestLimit && now < rest.lockedTo {
		rest.globalRequests = 0
		offset += 8
	}

	expiresTimestamp, exists := rest.locks[route]
	if exists && expiresTimestamp > now {
		offset += uint8(expiresTimestamp - now)
	}

	if rest.requestsSinceSweep%rest.MaxRequestsBeforeSweep == 0 {
		rest.requestsSinceSweep = 0

		go func() {
			for key, value := range rest.locks {
				if now > value {
					delete(rest.locks, key)
				}
			}
		}()
	}

	if offset != 0 {
		time.Sleep(time.Second * time.Duration(offset))
	}

	if jsonPayload == nil {
		request, err := http.NewRequest(method, misc.DiscordAPIUrl+route, nil)
		if err != nil {
			log.Println("Failed to create new http request.")
			panic(err)
		}
		req = request
	} else {
		body, err := json.Marshal(jsonPayload)
		if err != nil {
			log.Println("Failed to parse provided payload. Make sure it's in JSON format.")
			panic(err)
		}

		request, err := http.NewRequest(method, misc.DiscordAPIUrl+route, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("Failed to create new %s request to %s.\n", method, route)
			panic(err)
		}
		req = request
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "DiscordApp https://github.com/Amatsagu/tempest")
	req.Header.Add("Authorization", rest.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		rest.fails++
		if rest.fails == 3 {
			log.Printf("Failed to make http request 3 times to %s :: %s in a row! Please check internet connection and app credentials.\n", method, route)
			panic(err)
		} else {
			time.Sleep(time.Millisecond * time.Duration(250*rest.fails))
			return rest.Request(method, route, jsonPayload) // Try again after potential internet connection failure.
		}
	}
	defer res.Body.Close()

	rest.fails = 0
	remaining, err := strconv.ParseFloat(res.Header.Get("x-ratelimit-remaining"), 32)
	if err == nil && remaining == 0 {
		resetAt, _ := strconv.ParseFloat(res.Header.Get("x-ratelimit-reset"), 64) // If first succeeded then there's no need to check this one.
		rest.locks[route] = int64(resetAt + 6)
	}

	if res.StatusCode == 204 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to parse response body (json) for result of %s :: %s\n", method, route)
		panic(err)
	}

	if res.StatusCode == 429 {
		rateErr := rateLimitError{}
		json.Unmarshal(body, &rateErr)
		time.Sleep(time.Second * time.Duration(rateErr.RetryAfter+2.5))
		return rest.Request(method, route, jsonPayload) // Try again after rate limit.
	}

	return body
}
