// Package models defines common models for the application.
package models

import "time"

type LoginData struct {
	DataType string `json:"_"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Info     string `json:"_"`
}

type CardData struct {
	DataType   string `json:"_"`
	CVV        string `json:"cvv"`
	Number     string `json:"number"`
	ExpDate    string `json:"exp_date"`
	HolderName string `json:"holder_name"`
	Info       string `json:"_"`
}

type TextData struct {
	DataType string `json:"_"`
	Content  string `json:"Content"`
	Info     string `json:"_"`
}

type BinaryData struct {
	DataType   string `json:"_"`
	ObjectName string `json:"_"`
	Content    []byte `json:"Content"`
	Info       string `json:"_"`
}

type Metadata struct {
	ID          int
	DataType    string
	Description string
	CreatedAt   time.Time
}
