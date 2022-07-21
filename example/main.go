package main

import (
	"encoding/json"
	"tempest"
	"tempest/structs"
)

const Token string = "Bot XYZ"

func main() {
	rest := tempest.CreateRest(Token)
	res := rest.Request("GET", "/users/390394829789593601", nil)

	user := structs.User{}
	json.Unmarshal(res, &user)
}
