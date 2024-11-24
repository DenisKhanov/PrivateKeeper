package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/DenisKhanov/PrivateKeeper/internal/secure"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// TODO подумать над тем, чтобы объединить методы шифрования в один
func (d *ServiceData) AddLoginPasswordData(ctx context.Context, userID uuid.UUID, data models.LoginData) error {
	var encryptedData models.KeepData
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("error marshaling login/password data")
		return err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return err
	}
	logrus.Info(string(byteData))
	encrypted, err := secure.EncryptDataAES(decryptedKey, byteData)
	if err != nil {
		logrus.WithError(err).Error("error encrypting login/password data")
		return err
	}
	encryptedData = models.KeepData{
		DataType:      data.DataType,
		EncryptedData: encrypted,
		Info:          data.Info,
	}
	if err = d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err = d.repository.AddLoginPasswordData(ctx, tx, userID, encryptedData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (d *ServiceData) AddCardData(ctx context.Context, userID uuid.UUID, data models.CardData) error {
	var encryptedData models.KeepData
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("error marshaling bank card data")
		return err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return err
	}
	encrypted, err := secure.EncryptDataAES(decryptedKey, byteData)
	if err != nil {
		logrus.WithError(err).Error("error encrypting bank card data")
		return err
	}
	encryptedData = models.KeepData{
		DataType:      data.DataType,
		EncryptedData: encrypted,
		Info:          data.Info,
	}
	if err = d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err = d.repository.AddCardData(ctx, tx, userID, encryptedData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *ServiceData) AddTextData(ctx context.Context, userID uuid.UUID, data models.TextData) error {
	var encryptedData models.KeepData
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("error marshaling text data")
		return err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return err
	}
	encrypted, err := secure.EncryptDataAES(decryptedKey, byteData)
	if err != nil {
		logrus.WithError(err).Error("error encrypting text data")
		return err
	}
	encryptedData = models.KeepData{
		DataType:      data.DataType,
		EncryptedData: encrypted,
		Info:          data.Info,
	}
	if err = d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err = d.repository.AddTextData(ctx, tx, userID, encryptedData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *ServiceData) AddBinaryData(ctx context.Context, userID uuid.UUID, data models.BinaryData) error {
	var binaryInfo models.EncryptedBinaryData
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("error marshaling binary data")
		return err
	}
	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return err
	}
	encryptedContent, err := secure.EncryptDataAES(decryptedKey, byteData)
	if err != nil {
		logrus.WithError(err).Error("error encrypting binary data")
		return err
	}
	binaryInfo = models.EncryptedBinaryData{
		DataType:   data.DataType,
		ObjectName: data.ObjectName,
		Info:       data.Info,
	}
	if _, err = d.s3Repository.AddBinaryData(ctx, binaryInfo, encryptedContent); err != nil {
		return err
	}

	if err = d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err = d.repository.AddBinaryData(ctx, tx, userID, binaryInfo); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
