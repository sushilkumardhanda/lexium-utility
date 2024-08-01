package datahandler

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

