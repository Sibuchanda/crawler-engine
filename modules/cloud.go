package modules

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type PersistentMemory interface {
	Connect(endpoint, accessKey, secretAccessKey string) error
	UploadFile(bucketName, objectName, filePath, contentType string) error
	DownloadFile(bucketName, objectName, filePath string) error
}

// MinIO Type to mention MinIO instance
type MinIO struct {
	Client *minio.Client
}

// Connect Connectes to MinIO Instance
func (t *MinIO) Connect(endpoint, accessKey, secretAccessKey string) (err error) {
	t.Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: false,
	})
	return
}

// UploadFile Uploads File to MinIO Bucket
func (t *MinIO) UploadFile(bucketName, objectName, filePath, contentType string) (err error) {
	_, err = t.Client.FPutObject(context.Background(), bucketName,
		objectName, filePath, minio.PutObjectOptions{ContentType: contentType},
	)
	return
}

// DownloadFile Downloads File from MinIO Bucket
func (t *MinIO) DownloadFile(bucketName, objectName, filePath string) (err error) {
	err = t.Client.FGetObject(
		context.Background(),
		bucketName,
		objectName,
		filePath,
		minio.GetObjectOptions{},
	)
	return
}
