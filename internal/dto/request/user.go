package request

import (
	"time"

	"github.com/vector-ops/lifefolio/ent/user"
)

type UserRegisterIn struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserUpdateIn struct {
	FirstName  string          `json:"first_name,omitempty"`
	LastName   string          `json:"last_name,omitempty"`
	DOB        time.Time       `json:"dob,omitempty"`
	BloodGroup user.BloodGroup `json:"blood_group,omitempty"`
	Weight     float32         `json:"weight,omitempty"`
	Height     float32         `json:"height,omitempty"`
}
