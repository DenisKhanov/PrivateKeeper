package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
)

func (d *ServiceData) GetDecodedLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.LoginData, error) {
	loginPasswordData, err := d.repository.GetLoginPasswordData(ctx, userID, metadataID)
	if err != nil {
		return models.LoginData{}, err
	}
	//TODO тут будет происходить расшифровка данных
	var decodedLoginPasswordData = models.LoginData{
		Login:    loginPasswordData.Login,
		Password: loginPasswordData.Password,
		Info:     loginPasswordData.Info,
	}
	return decodedLoginPasswordData, nil
}

func (d *ServiceData) GetDecodedBankCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.CardData, error) {
	bankCardData, err := d.repository.GetCardData(ctx, userID, metadataID)
	if err != nil {
		return models.CardData{}, err
	}
	//TODO тут будет происходить расшифровка данных
	var decodedBankCardData = models.CardData{
		CVV:        bankCardData.CVV,
		Number:     bankCardData.Number,
		ExpDate:    bankCardData.ExpDate,
		HolderName: bankCardData.HolderName,
		Info:       bankCardData.Info,
	}
	return decodedBankCardData, nil
}

func (d *ServiceData) GetDecodedTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.TextData, error) {
	textData, err := d.repository.GetTextData(ctx, userID, metadataID)
	if err != nil {
		return models.TextData{}, err
	}
	//TODO тут будет происходить расшифровка данных
	var decodedTextData = models.TextData{
		Content: textData.Content,
		Info:    textData.Info,
	}
	return decodedTextData, nil
}
func (d *ServiceData) GetDecodedBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.BinaryData, error) {
	binaryData, err := d.repository.GetBinaryData(ctx, userID, metadataID)
	if err != nil {
		return models.BinaryData{}, err
	}
	decodedData, err := d.s3Repository.GetBinaryData(ctx, binaryData.ObjectName)
	//TODO тут будет происходить расшифровка данных
	var decodedBinaryData = models.BinaryData{
		Content:    decodedData,
		ObjectName: binaryData.ObjectName,
		Info:       binaryData.Info,
	}
	return decodedBinaryData, err
}
