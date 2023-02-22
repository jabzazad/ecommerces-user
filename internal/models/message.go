package models

// Message message model
type Message struct {
	Code    int    `json:"code" gorm:"-"`
	Message string `json:"message" gorm:"-"`
}

// NewSuccessMessage new message success model
func NewSuccessMessage() *Message {
	return &Message{
		Code:    200,
		Message: "success",
	}
}
