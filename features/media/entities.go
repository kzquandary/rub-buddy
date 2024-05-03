package media

import "github.com/labstack/echo/v4"

type Media struct {
	ImageFileName string `form:"image_file_name"`
}

type MediaHandlerInterface interface {
	UploadMedia() echo.HandlerFunc
}
