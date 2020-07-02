package model

type Game struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Started bool   `json:"started"`
}
