package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

//TODO оптимизировать запросы (join)?

func (d *RepositoryData) AddLoginPasswordData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.LoginData) error {

	var loginID int
	const sqlQuery = `INSERT INTO logins_passwords (login, password) VALUES ($1, $2) RETURNING id`
	err := tx.QueryRow(ctx, sqlQuery, data.Login, data.Password).Scan(&loginID)
	if err != nil {
		logrus.WithError(err).Error("Login/password don't save in database logins_passwords.")
		return err
	}
	var metadataID int
	const sqlQuery2 = `INSERT INTO metadata (data_type, website) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, sqlQuery2, data.DataType, data.Info).Scan(&metadataID)
	if err != nil {
		logrus.WithError(err).Error("Login/password don't save in database metadata.")
		return err
	}
	const sqlQuery3 = `INSERT INTO data_units (user_id, login_password_id, metadata_id)
	VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, sqlQuery3, userID, loginID, metadataID)
	if err != nil {
		logrus.WithError(err).Error("Login/password don't save in database data_units.")
		return err
	}
	return nil
}

func (d *RepositoryData) AddCardData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.CardData) error {
	var cardID int
	const sqlQuery = `INSERT INTO bank_cards (cvv,number,expiration_date,holder_name) VALUES ($1,$2,$3,$4) RETURNING id`
	err := tx.QueryRow(ctx, sqlQuery, data.CVV, data.Number, data.ExpDate, data.HolderName).Scan(&cardID)
	if err != nil {
		logrus.WithError(err).Error("Card don't save in database bank_cards.")
		return err
	}
	var metadataID int
	const sqlQuery2 = `INSERT INTO metadata (data_type, bank) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, sqlQuery2, data.DataType, data.Info).Scan(&metadataID)
	if err != nil {
		logrus.WithError(err).Error("Card don't save in database metadata.")
		return err
	}
	const sqlQuery3 = `INSERT INTO data_units (user_id, bank_card_id, metadata_id)
	VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, sqlQuery3, userID, cardID, metadataID)
	if err != nil {
		logrus.WithError(err).Error("Card don't save in database data_units.")
		return err
	}
	return nil
}

func (d *RepositoryData) AddTextData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.TextData) error {
	var textID int
	const sqlQuery = `INSERT INTO text_data (content) VALUES ($1) RETURNING id`
	err := tx.QueryRow(ctx, sqlQuery, data.Content).Scan(&textID)
	if err != nil {
		logrus.WithError(err).Error("Text don't save in database text_data.")
		return err
	}
	var metadataID int
	const sqlQuery2 = `INSERT INTO metadata (data_type, text_data_description) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, sqlQuery2, data.DataType, data.Info).Scan(&metadataID)
	if err != nil {
		logrus.WithError(err).Error("Text don't save in database metadata.")
		return err
	}
	const sqlQuery3 = `INSERT INTO data_units (user_id, text_data_id, metadata_id)
	VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, sqlQuery3, userID, textID, metadataID)
	if err != nil {
		logrus.WithError(err).Error("Text don't save in database data_units.")
		return err
	}
	return nil
}

func (d *RepositoryData) AddBinaryData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.BinaryData) error {
	var binaryID int
	const sqlQuery = `INSERT INTO binary_data (s3_object_name) VALUES ($1) RETURNING id`
	err := tx.QueryRow(ctx, sqlQuery, data.ObjectName).Scan(&binaryID)
	if err != nil {
		logrus.WithError(err).Error("Binary don't save in database binary_data.")
		return err
	}
	var metadataID int
	const sqlQuery2 = `INSERT INTO metadata (data_type, binary_data_description) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, sqlQuery2, data.DataType, data.Info).Scan(&metadataID)
	if err != nil {
		logrus.WithError(err).Error("Binary don't save in database metadata.")
		return err
	}
	const sqlQuery3 = `INSERT INTO data_units (user_id, binary_data_id, metadata_id)
	VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, sqlQuery3, userID, binaryID, metadataID)
	if err != nil {
		logrus.WithError(err).Error("Binary don't save in database data_units.")
		return err
	}
	return nil
}
