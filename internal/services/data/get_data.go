package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/DenisKhanov/PrivateKeeper/internal/secure"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (d *ServiceData) GetDecodedLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.LoginData, error) {
	var decryptedLoginData models.LoginData
	data, err := d.repository.GetLoginPasswordData(ctx, userID, metadataID)
	if err != nil {
		return models.LoginData{}, err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return models.LoginData{}, err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return models.LoginData{}, err
	}
	decrypted, err := secure.DecryptDataAES(decryptedKey, data.EncryptedData)
	if err != nil {
		return models.LoginData{}, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(decrypted, &decryptedLoginData); err != nil {
		logrus.WithError(err).Error("error unmarshalling decrypted login/password data")
		return models.LoginData{}, err
	}

	decryptedLoginData.DataType = data.DataType
	decryptedLoginData.Info = data.Info

	return decryptedLoginData, nil
}

func (d *ServiceData) GetDecodedBankCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.CardData, error) {

	var decryptedCardData models.CardData
	data, err := d.repository.GetCardData(ctx, userID, metadataID)
	if err != nil {
		return models.CardData{}, err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return models.CardData{}, err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return models.CardData{}, err
	}
	decrypted, err := secure.DecryptDataAES(decryptedKey, data.EncryptedData)
	if err != nil {
		return models.CardData{}, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(decrypted, &decryptedCardData); err != nil {
		logrus.WithError(err).Error("error unmarshalling decrypted card data")
		return models.CardData{}, err
	}
	decryptedCardData.DataType = data.DataType
	decryptedCardData.Info = data.Info
	return decryptedCardData, nil
}

func (d *ServiceData) GetDecodedTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.TextData, error) {
	var decryptedTextData models.TextData
	data, err := d.repository.GetTextData(ctx, userID, metadataID)
	if err != nil {
		return models.TextData{}, err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return models.TextData{}, err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return models.TextData{}, err
	}
	decrypted, err := secure.DecryptDataAES(decryptedKey, data.EncryptedData)
	if err != nil {
		return models.TextData{}, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(decrypted, &decryptedTextData); err != nil {
		logrus.WithError(err).Error("error unmarshalling decrypted text data")
		return models.TextData{}, err
	}
	decryptedTextData.DataType = data.DataType
	decryptedTextData.Info = data.Info
	return decryptedTextData, nil
}
func (d *ServiceData) GetDecodedBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.BinaryData, error) {
	binaryData, err := d.repository.GetBinaryData(ctx, userID, metadataID)
	if err != nil {
		return models.BinaryData{}, err
	}
	data, err := d.s3Repository.GetBinaryData(ctx, binaryData.ObjectName)
	if err != nil {
		return models.BinaryData{}, err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return models.BinaryData{}, err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return models.BinaryData{}, err
	}
	decrypted, err := secure.DecryptDataAES(decryptedKey, data)
	if err != nil {
		return models.BinaryData{}, err
	}
	var decryptedBinaryData = models.BinaryData{
		DataType: binaryData.DataType,
		Content:  decrypted,
		Info:     binaryData.Info,
	}
	return decryptedBinaryData, nil
}
