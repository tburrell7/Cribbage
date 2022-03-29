package models

type Player struct {
    Name string //`bson:"first_name"`
}

func NewPlayer(Name string) Player {
	return Player{Name}
}