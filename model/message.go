package model

type Message struct {
	Msg  []byte `json:"msg"`
	User []byte `json:"user"`
	Type int    `json:"type"`
}
