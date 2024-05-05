package s3

import (
	"bytes"
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

// AddBinaryData uploads binary data to the specified MinIO bucket and returns the object's URL.
// Accepts a context (`ctx`), the bucket name (`bucketName`), the object name (`objectName`), and a byte slice (`data`) containing the binary data.
//
// First, it creates a `Reader` for the data and sets the `contentType` to "application/octet-stream".
// Then, it uses MinIO client's `PutObject` method to upload the data to the specified MinIO bucket.
// If the upload is successful, it logs the size of the uploaded data.
//
// Returns:
// - The URL of the uploaded object if the upload is successful.
// - An error if there is a problem during data upload.
func (d *KeeperMinio) AddBinaryData(ctx context.Context, data models.BinaryData) error {
	reader := bytes.NewReader(data.Content)
	//TODO разобраться с выставлением правильного контент типа

	opts := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
		UserMetadata: map[string]string{
			"Description": data.Info,
		},
	}
	info, err := d.Client.PutObject(ctx, d.Bucket, data.ObjectName, reader, int64(len(data.Content)), opts)
	if err != nil {
		logrus.WithError(err).Error("create file failed")
		return err
	}
	logrus.Infof("Successfully uploaded %s of size %d\n", data.ObjectName, info.Size)
	return nil
}
