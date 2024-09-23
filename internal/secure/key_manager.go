package secure

import (
	"crypto/rsa"
	"github.com/google/uuid"
)

type KeyManager interface {
	// Генерация нового симметричного ключа
	GenerateAESKey() ([]byte, error)

	// Шифрование данных с использованием AES ключа
	EncryptDataAES(key []byte, data []byte) ([]byte, error)

	// Расшифрование данных с использованием ключа
	DecryptDataAES(key []byte, encryptedData []byte) ([]byte, error)

	// Генерация новой пары асимметричных ключей
	GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error)

	// Шифрование AES ключа при помощи публичного RSA ключа
	EncryptAESKey(publicKey *rsa.PublicKey, aesKey []byte) (string, error)

	// Расшифрование AES ключа при помощи приватного RSA ключа
	DecryptAESKey(privateKey *rsa.PrivateKey, encryptedAESKey []byte) ([]byte, error)

	// Получение ключа по ID пользователя
	GetAESKeyByID(userID uuid.UUID) ([]byte, error)

	// Сохранение ключа в менеджере
	SaveKey(keyID string, key []byte) error

	// Удаление ключа из менеджера
	DeleteKey(keyID string) error

	// Проверка разрешения доступа к ключу
	CheckPermission(keyID string, userID string, operation string) bool

	// Журналирование операций с ключами
	LogOperation(operation string, keyID string, userID string) error

	// Ротация симметричных ключей
	RotateSymmetricKeys() error

	// Ротация асимметричных ключей
	RotateAsymmetricKeys() error
}
