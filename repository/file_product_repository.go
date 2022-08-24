package repository

import (
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileProductRepository interface {
	CreateProductImage(files []*multipart.FileHeader, ctx *gin.Context) ([]string, error)
}

type fileProductRepository struct {
	path string
}

func (f *fileProductRepository) CreateProductImage(files []*multipart.FileHeader, ctx *gin.Context) ([]string, error) {
	var detailMediaProducts []string

	path := f.path + `\` + uuid.New().String() + `\`

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, errors.New("failed when getting file")
	}

	for _, file := range files {
		newFileName := strings.Split(file.Filename, ".")
		if len(newFileName) != 2 {
			return nil, errors.New("Unrecognize file extension")
		}

		newPath := path + "img-product-" + uuid.New().String() + "." + newFileName[1]

		if err := ctx.SaveUploadedFile(file, newPath); err != nil {
			return nil, errors.New("failed while saving file")
		}

		detailMediaProducts = append(detailMediaProducts, newPath)
	}

	return detailMediaProducts, nil
}

func NewFileProductRepository(path string) FileProductRepository {
	return &fileProductRepository{
		path: path,
	}
}
