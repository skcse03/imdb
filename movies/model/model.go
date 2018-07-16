package model

//ListMovie : list move response
type ListMovie struct {
	MovieName   string `json:"movie_name" gorm:"column:movie_name"`
	ReleaseYear string `json:"release_year" gorm:"column:relase_year"`
	Actors      string `json:"actors" gorm:"column:actor_name"`
	Producer    string `json:"producer" gorm:"column:producer_name"`
	Plot        string `json:"plot" gorm:"column:plot"`
	ImagePath   string `json:"image_path" gorm:"coumn:image"`
}

//AddMovie
type AddMovie struct {
	Movie      Movie `json:"movie"`
	ActorsID   []int `json:"actors_id"`
	ProducerID int   `json:"producer_id"`
}

//Actor : Actor Details
type Actor struct {
	ActorID   int    `json:"actor_id"  gorm:"column:actor_id"`
	ActorName string `json:"actor_name" valid:"required" gorm:"column:name"`
	Sex       string `json:"sex" valid:"required" gorm:"column:sex"`
	DOB       string `json:"dob" valid:"required" gorm:"column:dob"`
	Bio       string `josn:"bio" gorm:"column:bio"`
}

//Producer
type Producer struct {
	ProducerID   int    `json:"producer_id" gorm:"column:producer_id"`
	ProducerName string `json:"producer_name" valid:"required"`
	Sex          string `json:"sex" valid:"required" gorm:"column:sex"`
	DOB          string `json:"dob" valid:"required" gorm:"column:dob"`
	Bio          string `josn:"bio" gorm:"column:bio"`
}

//Movie
type Movie struct {
	MovieName   string `json:"movie_name" valid:"required"`
	ReleaseYear string `json:"release_year"`
	Plot        string `json:"plot"`
	Poster      string `json:"poster"`
}

//UpdateMovie : update movie
type UpdateMovie struct {
	MovieName   string `json:"movie_name" valid:"required"`
	ReleaseYear string `json:"release_year" valid:"required"`
}
