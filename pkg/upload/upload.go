package upload

import (
	"context"
	"errors"
	"reservation/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(file interface{}, folder, filename string) (secureUrl string, err error) {
	cloudName := config.LoadConfig().CloudinaryCloudName
	apiKey := config.LoadConfig().CloudinaryAPIKey
	apiSecret := config.LoadConfig().CLoudinaryAPISecret
	folder = config.LoadConfig().CloudinaryFolder + "/" + folder
	request, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	response, err := request.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder:   folder,
		PublicID: filename,
	})
	if err != nil {
		err = errors.New("failed to upload file to cloudinary: " + err.Error())
		return
	}

	secureUrl = response.SecureURL

	return
}
