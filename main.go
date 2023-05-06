package main

import (
	"ntrip/server"
)

func main() {
	go server.Start("0.0.0.0", "30451")

	select {}
}
