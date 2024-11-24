package domain

// validate
const (
	MinLoginLength    = 6
	MaxLoginLength    = 12
	MinPasswordLength = 8
	MaxPasswordLength = 16
)

const BinaryData = "BINARY_DATA"

// CTXKey is the type used as a context key for storing user ID.
type CTXKey string

// All constants used in project
const (
	TokenKey  CTXKey = "token"  // TokenKey is the specific key used in the context to store token.
	UserIDKey CTXKey = "userID" // UserIDKey is the specific key used in the context to store user ID.
)

// TLS
const (
	CertPEM       string = "cert.pem"       // CertPEM is the file name for TLS cert
	PrivateKeyPEM string = "privateKey.pem" // PrivateKeyPEM is the file name for TLS private key
)
