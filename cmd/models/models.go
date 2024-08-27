package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Signup request body model
type SignupRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Organisation string `json:"organisation"`
}

// User model
type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name         string             `bson:"name,omitempty" json:"name"`
	Username     string             `bson:"username,omitempty" json:"username"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	Organisation string             `bson:"organisation,omitempty" json:"organisation"`
	Role         string             `bson:"role,omitempty" json:"role"`
}

// this function converts struct to map
func (u *User) ToMap() map[string]string {
	return map[string]string{
		"id":           u.Id.Hex(),
		"name":         u.Name,
		"username":     u.Username,
		"organisation": u.Organisation,
		"role":         u.Role,
	}
}

// List of User model
type Users struct {
	Users []User `json:"users"`
}

// Login Request body model
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
