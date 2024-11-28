package main

import "mosquito/mnet"

func main() {
	server := mnet.NewServer("Mosquito")
	server.Serve()
}
