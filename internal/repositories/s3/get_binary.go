package s3

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"io"
)

// GetBinaryData retrieves binary data from the specified MinIO bucket and object name.
// Accepts a context (`ctx`), the bucket name (`bucketName`), and the object name (`objectName`).
//
// It first retrieves the object from the specified MinIO bucket using the MinIO client (`d.Client`).
// If an error occurs while getting the object, it logs the error and returns it.
//
// Then, it reads all the data from the object using `io.ReadAll`.
// If an error occurs during the read operation, it logs the error and returns it.
//
// Returns:
// - The binary data (`data`) retrieved from the object if the operation is successful.
// - An error if there is a problem during the object retrieval or read operation.
func (d *KeeperMinio) GetBinaryData(ctx context.Context, objectName string) (data []byte, err error) {
	object, err := d.Client.GetObject(ctx, d.Bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		logrus.WithError(err).Error("Failed to get object")
		return nil, err
	}
	defer object.Close()

	//TODO написать реализацию скачивания чанками
	data, err = io.ReadAll(object)
	if err != nil {
		logrus.WithError(err).Error("Failed to read object")
		return nil, err
	}
	return data, nil
}
