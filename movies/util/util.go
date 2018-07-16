package util

import (
	//"database/sql"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rightjoin/aqua"
	"imdb/movies/model"
)

//GetMoviesList
func GetMoviesList(j aqua.Aide) (interface{}, error) {
	var (
		response []model.ListMovie
		db       *gorm.DB
		err      error
	)
	if db, err = dbConn(); err == nil {
		sql := `Select mv.movies_id, 
				mv.name as movie_name,
		        mv.release_date ,
		        mv.image as image,
		        mv.plot as plot,
		        p.name as producer_name,
		        GROUP_CONCAT(DISTINCT a.name ORDER BY movies_id) AS actor_name
		        from producers as p 
		        join movies as mv on mv.fk_producer_id = p.producer_id
		        join actor_movie_map as amp on amp.fk_movies_id = mv.movies_id
                join actors as a on a.actor_id = amp.fk_actor_id group by (movies_id);`
		err = db.Raw(sql).Find(&response).Error
	}
	return response, err
}

//AddMovies : Add new entry for new insert of movie
func AddMovies(j aqua.Aide) (err error) {
	j.LoadVars()
	var reqPayload model.AddMovie
	if err = json.Unmarshal([]byte(j.Body), &reqPayload); err == nil {
		if _, err = govalidator.ValidateStruct(reqPayload); err == nil {
			if reqPayload.Movie.MovieName != "" {
				if err = upsertMovie(reqPayload.Movie, reqPayload.ProducerID); err == nil {
					fkMovieID, _ := getMovieID(reqPayload.Movie.MovieName)
					for i := 0; i < len(reqPayload.ActorsID); i++ {
						err = upsertActorMovieMap(
							reqPayload.ActorsID[i], fkMovieID)
					}
				}
			}
		}
	}
	return
}

//getMovieID: get movies ID
func getMovieID(name string) (mvID int, err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		idQry := `SELECT movies_id from movies where name = ?;`
		err = db.Raw(idQry, name).Debug().Row().Scan(&mvID)
	}
	return
}

//AddActor : add Actor
func AddActor(j aqua.Aide) (err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		j.LoadVars()
		var reqPayload model.Actor
		if err = json.Unmarshal([]byte(j.Body), &reqPayload); err == nil {
			if _, err = govalidator.ValidateStruct(reqPayload); err == nil {
				upsert := `INSERT INTO actors (name,sex,dob,bio) values (?, ?, ?, ?);`
				err = db.Debug().Exec(upsert, reqPayload.ActorName, reqPayload.Sex, reqPayload.DOB, reqPayload.Bio).Error
			}
		}
	}
	return
}

//AddProducer : add producer
func AddProducer(j aqua.Aide) (err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		j.LoadVars()
		var reqPayload model.Producer
		if err = json.Unmarshal([]byte(j.Body), &reqPayload); err == nil {
			if _, err = govalidator.ValidateStruct(reqPayload); err == nil {
				upsert := `INSERT INTO producers (name,sex,dob,bio) values (?, ?, ?, ?);`
				err = db.Debug().Exec(upsert, reqPayload.ProducerName, reqPayload.Sex, reqPayload.DOB, reqPayload.Bio).Error
			}
		}
	}
	return
}

//GetActor : get all actor for drop down option
func GetActor(j aqua.Aide) (response []model.Actor, err error) {
	var (
		db *gorm.DB
	)
	if db, err = dbConn(); err == nil {
		actorList := `SELECT actor_id, name, bio FROM actors;`
		err = db.Raw(actorList).Find(&response).Error
	}
	return
}

//GetProducer : get all producer for drop down option
func GetProducer(j aqua.Aide) (response []model.Producer, err error) {
	var (
		db *gorm.DB
	)
	if db, err = dbConn(); err == nil {
		producerList := `SELECT producer_id, name, bio FROM producers;`
		err = db.Raw(producerList).Find(&response).Error
	}
	return
}

//upsertProducer: update or insert producer
func upsertProducer(producer model.Producer) (err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		upsert := `INSERT INTO producers (name,sex,dob,bio) values (?,?,?,?) ON DUPLICATE KEY update name = ? ;`
		err = db.Debug().Exec(upsert, producer.ProducerName, producer.Sex, producer.Bio, producer.ProducerName).Error
	}
	return
}

//upsertActor
func upsertActor(actor model.Actor) (err error) {

	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		upsert := `INSERT INTO actors (name,sex,bio) values (?,?,?)
		 ON DUPLICATE KEY update name = ? ;`
		err = db.Debug().Exec(upsert, actor.ActorName, actor.Sex, actor.Bio, actor.ActorName).Error
	}
	return

}

//upsertMovie
func upsertMovie(movie model.Movie, producerID int) (err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		upsert := `INSERT INTO movies (name,fk_producer_id,release_date,plot,image) 
		values (?,?,?,?,?) ON DUPLICATE KEY update name = ? ,release_date = ? ;`
		err = db.Debug().Exec(upsert, movie.MovieName, producerID, movie.ReleaseYear, movie.Plot, movie.Poster, movie.MovieName, movie.ReleaseYear).Error
	}
	return
}

//upsertActorMovieMap
func upsertActorMovieMap(fkActorID int, fkMovieID int) (err error) {
	var db *gorm.DB

	if db, err = dbConn(); err == nil {
		upsert := `INSERT INTO actor_movie_map 
				   (fk_movies_id,fk_actor_id) values (?,?)
				   ON DUPLICATE KEY update fk_movies_id = ?, 
				   fk_actor_id = ?;`
		err = db.Debug().Exec(upsert, fkMovieID, fkActorID, fkMovieID, fkActorID).Error
	}
	return
}

//updateMovie : update Movie Details
func UpdateMovie(j aqua.Aide) (err error) {
	j.LoadVars()
	var reqPayload model.UpdateMovie
	if err = json.Unmarshal([]byte(j.Body), &reqPayload); err == nil {
		if _, err = govalidator.ValidateStruct(reqPayload); err == nil {
			err = updateMovieDetails(reqPayload)
		}
	}
	return
}

//updateMovieDetails
func updateMovieDetails(reqPayload model.UpdateMovie) (err error) {
	var db *gorm.DB
	if db, err = dbConn(); err == nil {
		update := `update movies  set release_date= ? , name = ? where name= ?`
		err = db.Exec(update, reqPayload.ReleaseYear, reqPayload.MovieName, reqPayload.MovieName).Error
	}
	return
}

//dbConn : make conenction to database
func dbConn() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:spatico@/deltax")
	return
}
