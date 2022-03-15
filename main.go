package main

import (
	"github.com/ivanauliaa/shortener-go/router"
	"github.com/ivanauliaa/shortener-go/utils"
)

func main() {
	server := router.ServerRouter()
	server.Run(utils.PORT)
}
