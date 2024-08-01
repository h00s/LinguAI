package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key,omitempty"`
}
