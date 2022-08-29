package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
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
	cld, _ := cloudinary.NewFromParams("ihdiannaja", "954945529412874", "7mFstMRVYEOlO784FGNo09mfk_4")
	pathHold := "img-feed" + uuid.New().String()
	fileInput, err := file.Open()
	if err != nil {
		return "", err
	}
	result, err := cld.Upload.Upload(ctx, fileInput, uploader.UploadParams{
		PublicID: pathHold,
		Folder:   "Post Feed",
	})
	return result.SecureURL, err
}

func NewFileRepository(path string, pathFeed string, pathClientFeed string) FileRepository {
	return &fileRepository{
		path:           path,
		pathFeed:       pathFeed,
		pathClientFeed: pathClientFeed,
	}
}
