// Package models defines common models for the application.
package models

type KeepData struct {
	DataType      string `json:"data_type"`
	EncryptedData []byte `json:"encrypted_data"`
	Info          string `json:"info"`
}
type EncryptedBinaryData struct {
	DataType   string `json:"data_type"`
	ObjectName string `json:"name"`
	Info       string `json:"info"`
}
