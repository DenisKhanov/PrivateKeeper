package s3

import (
	"bytes"
	"context"
	"fmt"
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
func (d *KeeperMinio) AddBinaryData(ctx context.Context, objectName string, data []byte) (string, error) {
	reader := bytes.NewReader(data)
	contentType := "application/octet-stream"
	info, err := d.Client.PutObject(ctx, d.Bucket, objectName, reader, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logrus.WithError(err).Error("create file failed")
		return "", err
	}
	logrus.Infof("Successfully uploaded %s of size %d\n", objectName, info.Size)
	objectURL := fmt.Sprintf("%s/%s/%s", d.Client.EndpointURL(), d.Bucket, objectName)
	return objectURL, nil
}
