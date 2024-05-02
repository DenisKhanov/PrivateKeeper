package s3

import (
	"github.com/minio/minio-go/v7"
)

type KeeperMinio struct {
	Bucket string
	Client *minio.Client
}

func NewMinio(client *minio.Client, bucket string) (*KeeperMinio, error) {

	return &KeeperMinio{
		Client: client,
		Bucket: bucket,
	}, nil
}
