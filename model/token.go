package model

type Token struct {
	Id    string `bson:"_id"`
	Email string `bson:"email"`
	Pass  string `bson:"pass"`
}
