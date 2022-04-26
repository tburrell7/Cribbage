package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Score struct {
	//Id 		primitive.ObjectID 	`bson:"_id,omitempty"`
	LeftScore 	int	`bson:"leftscore"`
	RightScore	int	`bson:"rightscore"`
}