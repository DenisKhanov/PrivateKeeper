package data

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/DenisKhanov/PrivateKeeper/internal/secure"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"io"
)

func (d *ServiceData) UploadBigData(ctx context.Context, userID uuid.UUID, binaryInfo models.EncryptedBinaryData, content io.Reader) error {
	// Read and encrypt data in chunks
	fmt.Println("Service UploadBigData")
	buffer := bytes.NewBuffer(nil)
	reader := bufio.NewReader(content)
	for {
		chunk := make([]byte, 1024) // Размер блока может быть настроен
		n, err := reader.Read(chunk)
		if err != nil && err != io.EOF {
			logrus.WithError(err).Error("Failed to read data")
			return err
		}
		if n == 0 {
			break
		}
		buffer.Write(chunk[:n])
	}

	encryptedKey, err := d.repository.GetEncryptedKey(ctx, userID)
	if err != nil {
		return err
	}
	decryptedKey, err := secure.DecryptAESKey(d.privateKey, encryptedKey)
	if err != nil {
		return err
	}
	encryptedContent, err := secure.EncryptDataAES(decryptedKey, buffer.Bytes())
	if err != nil {
		logrus.WithError(err).Error("error encrypting binary data")
		return err
	}
	info, err := d.s3Repository.AddBinaryData(ctx, binaryInfo, encryptedContent)
	if err != nil {
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
	logrus.Infof("Successfully uploaded %s of size %d\n", binaryInfo.ObjectName, info.Size)
	return nil
}
