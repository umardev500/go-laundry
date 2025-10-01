package upload

import (
	"context"
	"mime/multipart"
)

type Service interface {
	SaveFile(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	GetFileURL(filePath string, isSecure bool) string
}
