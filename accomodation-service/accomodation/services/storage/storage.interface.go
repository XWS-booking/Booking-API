package storage

type IStorageService interface {
	UploadImage(data []byte) error
}
