package main

import (
	"tempest"
)

const Token string = "Bot XYZ"

func main() {
	rest := tempest.Rest{Token: Token}
	rest.Request("GET", "/fff", map[string]string{"content": "fff"})
}
