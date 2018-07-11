package main

import (
	"github.com/rightjoin/aqua"
	imdb "imdb/movies/service"
)

func main() {
	service := aqua.NewRestServer()

	service.AddService(&imdb.Movies{})
	server.RunWith(conf.Int("server.port", 8090), true)
}
