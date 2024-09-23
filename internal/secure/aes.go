package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

// GenerateAESKey генерирует случайный ключ AES.
func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		logrus.WithError(err).Error("Error generating AES key")
		return nil, err
	}
	return key, nil
}

// EncryptDataAES шифрует данные с использованием ключа AES.
func EncryptDataAES(key, plaintext []byte) ([]byte, error) {
	// Создание блока AES с использованием ключа.
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.WithError(err).Error("Failed to create AES cipher")
		return nil, err
	}

	// Создание шифратора AES-GCM.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logrus.WithError(err).Error("Failed to create GCM")
		return nil, err
	}

	// Создание nonce (число, которое используется только один раз) для GCM.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logrus.WithError(err).Error("Failed to create nonce")
		return nil, err
	}

	// Шифрование данных с использованием AES-GCM.
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	// Возвращение зашифрованных данных вместе с nonce (которое нужно для расшифрования).
	return append(nonce, ciphertext...), nil
}

// DecryptDataAES расшифровывает данные с использованием ключа AES.
func DecryptDataAES(key, encryptedData []byte) ([]byte, error) {
	// Создание блока AES с использованием ключа.
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.WithError(err).Error("Failed to create AES cipher")
		return nil, err
	}

	// Создание шифратора AES-GCM.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logrus.WithError(err).Error("Failed to create GCM")
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, fmt.Errorf("encryptedData too short")
	}

	// Расшифрование данных с использованием AES-GCM.
	nonce, newData := encryptedData[:nonceSize], encryptedData[nonceSize:]
	fmt.Println(encryptedData)
	fmt.Println(newData)
	plaintext, err := aesGCM.Open(nil, nonce, newData, nil)
	if err != nil {
		logrus.WithError(err).Error("Failed to decrypt data")
		return nil, err
	}

	return plaintext, nil
}
