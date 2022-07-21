package tempest

import "time"

type ClientOptions struct {
	Token string `binding:"required"`
	/** Your app/bot's user id. */
	ApplicationId uint64 `binding:"required"`
	/** Hash like key used to verify incoming payloads from Discord. */
	PublicKey string `binding:"required"`
	/** Settings related to cooldown on commands. Leave this object undefined to disable cooldowns. */
	cooldown ClientOptionsCooldown
}

type ClientOptionsCooldown struct {
	/** The cooldown between command usage in milliseconds. */
	duration uint8
}

type Client struct {
	Rest Rest
}

func (client Client) GetLatency() int64 {
	startedAt := time.Now().Unix()
	client.Rest.Request("GET", "/gateway", nil)
	return time.Now().Unix() - startedAt
}

func CreateClient(options ClientOptions) Client {
	rest := Rest{
		Token:                  options.Token,
		MaxRequestsBeforeSweep: 100,
		GlobalRequestLimit:     50,
		globalRequests:         0,
		requestsSinceSweep:     0,
		lockedTo:               0,
		locks:                  make(map[string]int64, 100),
		fails:                  0,
	}

	return Client{Rest: rest}
}
