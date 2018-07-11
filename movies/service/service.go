package service

import (
	"github.com/rightjoin/aqua"
	"github.com/tolexo/aero/conf"
	"imdb/movies/util"
)

type Movies struct {
	aqua.RestService `prefix:"saurabh/imdb" root:"/" version:"1"`
	getMoviesList    aqua.POST `url:"/get/movies/list/"`
	addMovies        aqua.GET  `url:"/new/movies/"`
	updateMovie      aqua.POST `url:"/update/"`
}

//GetMoviesList : get movies list based of filter
func (mvs *Movies) GetMoviesList(j aqua.Aide) {

	util.GetMoviesList(j)
}

//AddMovies : add new movies
func (mvs *Movies) AddMovies(j aqua.Aide) {

}

//UpdateMovie: update movie details
func (mvs *Movies) UpdateMovie(j aqua.Aide) {

}
