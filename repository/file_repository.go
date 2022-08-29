package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileRepository interface {
	Save(file multipart.File, fileName string) (string, error)
	SavefromCtx(file *multipart.FileHeader, fileName string, ctx *gin.Context) (string, error)
}

type fileRepository struct {
	path           string
	pathFeed       string
	pathClientFeed string
}

func (f *fileRepository) Save(file multipart.File, fileName string) (string, error) {
	fileLocation := filepath.Join(f.path, fileName)
	out, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func (f *fileRepository) SavefromCtx(file *multipart.FileHeader, fileName string, ctx *gin.Context) (string, error) {
	pathHold := "img-feed-" + uuid.New().String() + "." + fileName
	path := f.pathFeed + pathHold
	pathClient := f.pathClientFeed + pathHold
	err := ctx.SaveUploadedFile(file, path)
	return pathClient, err
}

func NewFileRepository(path string, pathFeed string, pathClientFeed string) FileRepository {
	return &fileRepository{
		path:           path,
		pathFeed:       pathFeed,
		pathClientFeed: pathClientFeed,
	}
}
