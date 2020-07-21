package utils

type ApiErr struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
