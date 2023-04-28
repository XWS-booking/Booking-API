package storage

import (
	"bytes"
	storage_go "github.com/supabase-community/storage-go"
	"os"
)

type StorageService struct {
	Client *storage_go.Client
}

func NewStorageService() *StorageService {
	storageUrl := os.Getenv("STORAGE_URL")
	serviceToken := os.Getenv("STORAGE_SERVICE_TOKEN")
	client := storage_go.NewClient(storageUrl, serviceToken, nil)
	return &StorageService{
		Client: client,
	}
}

func (storageService *StorageService) UploadImage(data []byte) error {
	bucket := storageService.getMainBucket()
	reader := bytes.NewReader(data)
	storageService.Client.UploadFile(bucket.Id, "", reader)
	return nil
}

func (storageService *StorageService) getMainBucket() storage_go.Bucket {
	mainBucket := os.Getenv("STORAGE_MAIN_BUCKET")

	bucket, _ := storageService.Client.GetBucket(mainBucket)
	return bucket
}
