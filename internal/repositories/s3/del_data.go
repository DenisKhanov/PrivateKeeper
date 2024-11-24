package s3

import (
	"context"
	errors3 "github.com/DenisKhanov/PrivateKeeper/internal/errors"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

func (d *KeeperMinio) DelData(ctx context.Context, objectName string) error {
	if err := d.Client.RemoveObject(ctx, d.Bucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		logrus.WithError(err).Error("Failed to remove object from minio.")
		return errors3.ErrDelData
	}
	logrus.Info("Object removed successfully")
	return nil
}
