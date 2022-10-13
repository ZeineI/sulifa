package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type HexId struct {
	ID primitive.ObjectID `bson:"_id"`
}

type RegisterReq struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type LoginReq struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type CreateRoomReq struct {
	Username string `bson:"username"`
}

type JoinRoomReq struct {
	Username string `bson:"username"`
	RoomID   string `bson:"roomID"`
}
