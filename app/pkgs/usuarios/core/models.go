package core

type Usuario struct {
	ID        string  `json:"id,omitempty"`
	Nome      string  `json:"nome"`
	Email      string  `json:"email"`
}