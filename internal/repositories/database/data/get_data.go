package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

//TODO возвращать в сервис кастомные ошибки

func (d *RepositoryData) GetLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.LoginData, error) {
	var loginData models.LoginData
	const sqlQuery = `
        SELECT
            logins_passwords.login,
            logins_passwords.password,
            metadata.website
        FROM
            data_units
        JOIN
            logins_passwords ON data_units.login_password_id = logins_passwords.id
        JOIN
            metadata ON data_units.metadata_id = metadata.id
        WHERE
            data_units.user_id = $1
            AND data_units.metadata_id = $2;
    `

	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID, metadataID).Scan(
		&loginData.Login,
		&loginData.Password,
		&loginData.Info,
	); err != nil {
		logrus.WithError(err).Error("Error getting login/password data.")
		return models.LoginData{}, err
	}
	logrus.Info("Success getting login/password data.")
	return loginData, nil
}

func (d *RepositoryData) GetCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.CardData, error) {
	var cardData models.CardData
	const sqlQuery = `
		SELECT 
    		bank_cards.number,
    		bank_cards.cvv, 
    		bank_cards.expiration_date, 
    		bank_cards.holder_name,
    		metadata.bank
		FROM 
		    data_units 
		JOIN 
		    bank_cards ON data_units.bank_card_id = bank_cards.id
		JOIN 
		    metadata ON data_units.metadata_id = metadata.id
		WHERE 
		    data_units.user_id = $1
		    AND data_units.metadata_id = $2;
	`
	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID, metadataID).Scan(
		&cardData.Number,
		&cardData.CVV,
		&cardData.ExpDate,
		&cardData.HolderName,
		&cardData.Info,
	); err != nil {
		logrus.WithError(err).Error("Error getting bank card data.")
		return models.CardData{}, err
	}
	logrus.Info("Success getting bank card data.")
	return cardData, nil
}

func (d *RepositoryData) GetTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.TextData, error) {
	var textData models.TextData
	const sqlQuery = `
		SELECT 
		    text_data.content,
		    metadata.text_data_description
		FROM 
		    data_units
		JOIN 
		    text_data ON data_units.text_data_id = text_data.id
		JOIN
		    metadata ON data_units.metadata_id = metadata.id
		WHERE 
		    data_units.user_id=$1 
		  	AND data_units.metadata_id = $2
	`
	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID, metadataID).Scan(
		&textData.Content,
		&textData.Info,
	); err != nil {
		logrus.WithError(err).Error("Error getting text data.")
		return models.TextData{}, err
	}
	logrus.Info("Success getting text data.")
	return textData, nil
}

func (d *RepositoryData) GetBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.BinaryData, error) {
	var binaryData models.BinaryData
	const sqlQuery = `
		SELECT 
		    binary_data.s3_object_name,
			metadata.binary_data_description
		FROM 
		    data_units
		JOIN 
		    binary_data ON data_units.binary_data_id = binary_data.id
		JOIN
		    metadata ON data_units.metadata_id = metadata.id
		WHERE 
		    data_units.user_id = $1
		    AND data_units.metadata_id = $2;
	`
	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID, metadataID).Scan(
		&binaryData.ObjectName,
		&binaryData.Info,
	); err != nil {
		logrus.WithError(err).Error("Error getting binary data.")
		return models.BinaryData{}, err
	}
	logrus.Info("Success getting binary data.")
	return binaryData, nil
}
