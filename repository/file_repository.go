package repository

import (
	"io"
	"itdp-group3-backend/config"
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
	path config.MediaPath
}

func (f *fileRepository) Save(file multipart.File, fileName string) (string, error) {
	fileLocation := filepath.Join(f.path.Path, fileName)
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
	path := f.path.PathFeed + "img-feed-" + uuid.New().String() + "." + fileName
	err := ctx.SaveUploadedFile(file, path)
	return path, err
}

func NewFileRepository(path config.MediaPath) FileRepository {
	return &fileRepository{
		path: path,
	}
}
