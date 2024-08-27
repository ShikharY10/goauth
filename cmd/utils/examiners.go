package utils

import (
	"errors"

	"github.com/ShikharY10/goauth/cmd/models"
)

// Examines whether any fiels in empty or not. If empty it throws error.
func ExamineSignupRequestBody(request models.SignupRequest) error {
	if request.Name == "" {
		return errors.New("name not found")
	} else if request.Username == "" {
		return errors.New("username not found")
	} else if request.Organisation == "" {
		return errors.New("password not found")
	} else if request.Password == "" {
		return errors.New("organisation not found")
	}
	return nil
}

// Examines whether any fiels in empty or not. If empty it throws error.
func ExamineLoginRequestBody(request models.LoginRequest) error {
	if request.Username == "" {
		return errors.New("username is not found")
	} else if request.Password == "" {
		return errors.New("password not found")
	}
	return nil
}
