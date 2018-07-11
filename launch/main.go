package main

import (
	//"github.com/rightjoin/aero/conf"
	"github.com/rightjoin/aqua"
	imdb "imdb/movies/service"
)

// main function
func main() {
	server := aqua.NewRestServer()
	server.AddService(&imdb.Movies{})
	server.Run()
}
