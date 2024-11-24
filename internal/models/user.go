// Package models defines common models for the application.
package models

// User ...
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
