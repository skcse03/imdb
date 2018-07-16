package service

import (
	//"fmt"
	"github.com/rightjoin/aqua"
	"imdb/movies/util"
)

type Movies struct {
	aqua.RestService `prefix:"saurabh/imdb" root:"/" version:"1"`
	getMoviesList    aqua.POST `url:"/get/movies/list/"`
	addMovies        aqua.POST `url:"/new/movies/"`
	updateMovie      aqua.POST `url:"/update/"`
	getActor         aqua.GET  `url:"/get/actor/"`
	getProducer      aqua.GET  `url:"/get/producer/`
	addActor         aqua.POST `url:"/add/actor/"`
	addProducer      aqua.POST `url:"/add/producer/"`
}

//GetMoviesList : get movies list based of filter
func (mvs *Movies) GetMoviesList(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.GetMoviesList(j)
	return
}

//AddMovies : add new movies
func (mvs *Movies) AddMovies(j aqua.Aide) (err error) {
	err = util.AddMovies(j)
	return
}

// //UpdateMovie: update movie details
func (mvs *Movies) UpdateMovie(j aqua.Aide) (err error) {
	err = util.UpdateMovie(j)
	return
}

//AddActor: add new Actor
func (mvs *Movies) AddActor(j aqua.Aide) (err error) {
	err = util.AddActor(j)
	return
}

//AddProducer : add new Producer
func (mvs *Movies) AddProducer(j aqua.Aide) (err error) {
	err = util.AddProducer(j)
	return
}

//func GetActor: get Actors List
func (mvs *Movies) GetActor(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.GetActor(j)
	return
}

//func GetProducer: get Producers List
func (mvs *Movies) GetProducer(j aqua.Aide) (
	response interface{}, err error) {

	response, err = util.GetProducer(j)
	return
}
