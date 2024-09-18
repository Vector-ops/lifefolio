package request

import "github.com/google/uuid"

type InstRegisterIn struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}

type InstUpdateIn struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Location    string    `json:"location,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Password    string    `json:"password,omitempty"`
}
