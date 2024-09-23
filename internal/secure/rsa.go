package secure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
)

// GenerateRSAKeys Генерация RSA ключей
func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.WithError(err).Error("failed to generate RSA key")
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// GenerateRSAKeysToFiles Генерация RSA ключей и запись их в файлы
func GenerateRSAKeysToFiles(privateKeyPath, publicKeyPath string) error {
	// Генерация RSA приватного ключа
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.WithError(err).Error("failed to generate RSA key")
		return err
	}

	// Кодирование приватного ключа в PEM формат
	privFile, err := os.Create(privateKeyPath)
	if err != nil {
		logrus.WithError(err).Error("failed to create private key file")
		return err
	}
	defer privFile.Close()

	privPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if err := pem.Encode(privFile, &privPEM); err != nil {
		logrus.WithError(err).Error("failed to write private key to file")
		return err
	}

	// Кодирование публичного ключа в PEM формат
	pubFile, err := os.Create(publicKeyPath)
	if err != nil {
		logrus.WithError(err).Error("failed to create public key file")
		return err
	}
	defer pubFile.Close()

	pubPEM, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal public key")
		return err
	}

	pubBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubPEM,
	}

	if err := pem.Encode(pubFile, &pubBlock); err != nil {
		logrus.WithError(err).Error("failed to write public key to file")
		return err
	}

	return nil
}

// LoadRSAPrivateKey Загрузка приватного RSA ключа из файла
func LoadRSAPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	logrus.Info(filePath)
	privateKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Errorf("Error reading private key file: %v", err)
		return nil, err
	}

	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logrus.WithError(err).Error("failed to parse RSA private key")
		return nil, err
	}
	return privateKey, nil
}

// LoadRSAPublicKey Загрузка публичного RSA ключа из файла
func LoadRSAPublicKey(filePath string) (*rsa.PublicKey, error) {
	pubKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Errorf("Error reading public key file: %v", err)
		return nil, err
	}

	block, _ := pem.Decode(pubKeyFile)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.WithError(err).Error("failed to parse RSA public key")
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPublicKey, nil
}

// EncryptAESKey Шифрование AES ключа с использованием RSA
func EncryptAESKey(publicKey *rsa.PublicKey, aesKey []byte) ([]byte, error) {
	logrus.Info(aesKey)
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, aesKey, nil)
	if err != nil {
		logrus.WithError(err).Error("Error encrypting AES key")
		return nil, err
	}
	return encryptedKey, nil
}

// DecryptAESKey Функция для расшифровки зашифрованного AES ключа с использованием RSA приватного ключа
func DecryptAESKey(privateKey *rsa.PrivateKey, encryptedAESKey []byte) ([]byte, error) {
	decryptedAESKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedAESKey, nil)
	if err != nil {
		logrus.WithError(err).Error("Error decrypting AES key")
		return nil, err
	}
	return decryptedAESKey, nil
}
