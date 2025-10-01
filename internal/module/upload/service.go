package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/upload"
)

type serviceImpl struct {
	basePath upload.BasePath
	cfg      *config.Config
}

// GetFileURL implements upload.Service.
func (s *serviceImpl) GetFileURL(filePath string, isSecure bool) string {
	scheme := "http"
	if isSecure {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s:%s/%s", scheme, s.cfg.Server.Host, s.cfg.Server.Port, filePath)
}

// SaveFile implements upload.Service.
func (s *serviceImpl) SaveFile(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	dest := filepath.Join(string(s.basePath), folder, filename)

	// Ensure folder exists
	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return "", err
	}

	// Save file
	if err := saveMultipartFile(file, dest); err != nil {
		return "", err
	}

	return dest, nil
}

func saveMultipartFile(file *multipart.FileHeader, dest string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return err
}

func NewService(basePath upload.BasePath, cfg *config.Config) upload.Service {
	return &serviceImpl{
		basePath: basePath,
		cfg:      cfg,
	}
}
