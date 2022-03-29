package models

type Player struct {
    Name string `bson:"first_name"`
	Record string
	IsPlaying bool
	Opponent string
	CurrentScore string
}