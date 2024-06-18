package main

import (
	"github.com/travellingsalesman/src/container"
	"github.com/travellingsalesman/src/server"
)

func main() {
	readyContainer, err := container.BuildContainer()
	if err != nil {
		panic(err)
	}

	if err = readyContainer.Invoke(server.StartServer); err != nil {
		panic(err)
	}
	return
}
