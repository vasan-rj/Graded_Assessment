package model

type Account struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	PassKey  string `json:"passkey"`
}
