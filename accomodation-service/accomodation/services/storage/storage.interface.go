package storage

type IStorageService interface {
	UploadImage(data []byte, name string) (string, error)
}
