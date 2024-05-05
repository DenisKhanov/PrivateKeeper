// Package models defines common models and errors for the application.
package models

// User ...
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginData struct {
	DataType string `json:"data_type"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Info     string `json:"info"`
}

type CardData struct {
	DataType   string `json:"data_type"`
	CVV        string `json:"cvv"`
	Number     string `json:"number"`
	ExpDate    string `json:"exp_date"`
	HolderName string `json:"holder_name"`
	Info       string `json:"info"`
}

type TextData struct {
	DataType string `json:"data_type"`
	Content  string `json:"Content"`
	Info     string `json:"info"`
}

type BinaryData struct {
	DataType   string `json:"data_type"`
	ObjectName string `json:"name"`
	Content    []byte `json:"Content"`
	Info       string `json:"info"`
}

type Metadata struct {
	ID          int
	DataType    string
	Description string
}

// CTXKey is the type used as a context key for storing user ID.
type CTXKey string

// All constants used in project
const (
	TokenKey      CTXKey = "token"          // TokenKey is the specific key used in the context to store token.
	UserIDKey     CTXKey = "userID"         // UserIDKey is the specific key used in the context to store user ID.
	CertPEM       string = "cert.pem"       // CertPEM is the file name for TLS cert
	PrivateKeyPEM string = "privateKey.pem" // PrivateKeyPEM is the file name for TLS private key
)
