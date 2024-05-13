package bucket

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"rub_buddy/helper"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BucketInterface interface {
	UploadFile(file multipart.File, object string) (string, error)
	UploadFileHandler() echo.HandlerFunc
}

type Bucket struct {
	Client     *storage.Client
	ProjectID  string
	BucketName string
}

type ImageResponse struct {
	ImageUrl string `json:"image_url"`
}

func New(ProjectID string, BucketName string) (BucketInterface, error) {
	// Dev Mode
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials.json")

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Bucket{
		Client:     client,
		ProjectID:  ProjectID,
		BucketName: BucketName,
	}, nil
}

func (b *Bucket) UploadFile(file multipart.File, object string) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := b.Client.Bucket(b.BucketName).Object(object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", b.BucketName, object), nil
}

func (b *Bucket) UploadFileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Image is required", []interface{}{}))
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Image is required", []interface{}{}))
		}
		defer src.Close()

		image_url, err := b.UploadFile(src, "rubbuddy/"+uuid.New().String()+file.Filename)
		if err != nil {
			return err
		}

		var response = new(ImageResponse)
		response.ImageUrl = image_url
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Upload success", []interface{}{response}))
	}
}
