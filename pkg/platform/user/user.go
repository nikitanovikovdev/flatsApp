package user

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
}
