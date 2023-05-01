package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	storage_go "github.com/supabase-community/storage-go"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type StorageService struct {
	Client *storage_go.Client
}

func NewStorageService() *StorageService {
	storageUrl := os.Getenv("STORAGE_URL")
	serviceToken := os.Getenv("STORAGE_SERVICE_TOKEN")
	fmt.Println(storageUrl, serviceToken)
	client := storage_go.NewClient(storageUrl, serviceToken, nil)
	return &StorageService{
		Client: client,
	}
}

func (storageService *StorageService) UploadImage(data []byte, name string) (string, error) {
	bucket := storageService.getMainBucket()
	//reader := bytes.NewReader(data)
	//resp := storageService.Client.UploadFile(bucket.Id, name, reader)
	_, err := uploadImageToSupabase(data, name)
	if err != nil {
		fmt.Println(err)
	}
	url := storageService.Client.GetPublicUrl(bucket.Id, name)
	return url.SignedURL, nil
}

func uploadImageToSupabase(data []byte, name string) (string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	apiUrl := os.Getenv("STORAGE_URL") + "/object/" + os.Getenv("STORAGE_MAIN_BUCKET") + "/" + name

	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, bytes.NewReader(data)); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+os.Getenv("STORAGE_SERVICE_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	parsedBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res map[string]string
	err = json.Unmarshal(parsedBody, &res)
	if err != nil {
		return "", err
	}

	return res["Key"], nil
}

func (storageService *StorageService) getMainBucket() storage_go.Bucket {
	mainBucket := os.Getenv("STORAGE_MAIN_BUCKET")

	bucket, _ := storageService.Client.GetBucket(mainBucket)
	return bucket
}
