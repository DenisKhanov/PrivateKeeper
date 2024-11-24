package data

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	proto "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

func (d *GRPCData) UploadBinaryData(stream proto.KeeperDataV1_UploadBinaryDataServer) error {
	// Извлекаем UserID из контекста
	userID, ok := stream.Context().Value(domain.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Info("Could not extract UserID from ctx")
		return status.Error(codes.Internal, "could not find user ID in context")
	}

	// Инициализируем переменные для данных
	var binaryData models.EncryptedBinaryData
	reader, writer := io.Pipe() // Pipe для передачи данных между потоками

	// Получаем первый пакет данных (метаданные файла)
	req, err := stream.Recv()
	if err == io.EOF {
		return stream.SendAndClose(&proto.UploadBigBinaryResponse{
			Message: "Stream closed successfully",
		})
	}
	if err != nil {
		logrus.WithError(err).Error("Error receiving stream")
		return status.Error(codes.Unknown, "Error receiving stream")
	}

	// Заполняем метаданные
	binaryData.DataType = req.DataUnit.DataType.String()
	binaryData.ObjectName = req.DataUnit.GetBinaryData().GetObjectName()
	binaryData.Info = req.DataUnit.MetaInfo.GetBinaryDataDescription()

	// Запускаем горутину для обработки данных
	go func() {
		defer writer.Close() // Закрываем writer после записи всех данных

		// Чтение данных чанками
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Print("All data received")
				break
			}
			if err != nil {
				logrus.WithError(err).Error("Error receiving stream")
				return
			}

			// Получаем чанк данных
			chunk := req.GetDataUnit().GetBinaryData().GetContent()

			// Пишем чанк в pipe
			_, err = writer.Write(chunk)
			if err != nil {
				logrus.WithError(err).Error("Error writing to pipe")
				return
			}
		}
	}()

	// Обрабатываем данные и сохраняем их
	err = d.service.UploadBigData(stream.Context(), userID, binaryData, reader)
	if err != nil {
		logrus.WithError(err).Error("Failed to process and store data")
		return status.Errorf(codes.Internal, "Failed to process and store data: %v", err)
	}

	// Отправляем успешный ответ
	return stream.SendAndClose(&proto.UploadBigBinaryResponse{
		Message: "File uploaded successfully",
	})
}
