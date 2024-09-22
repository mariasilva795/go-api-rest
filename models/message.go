package models

type WebsockertMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
