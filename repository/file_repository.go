package repository

import (
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileRepository interface {
	SaveSingleFile(file multipart.File, ctx *gin.Context, folderName string) (string, error)
	SaveMultipleFiles(file *multipart.FileHeader, ctx *gin.Context, folderName string) (string, error)
}

type fileRepository struct {
	path           string
	pathFeed       string
	pathClientFeed string
}

func (f *fileRepository) SaveSingleFile(file multipart.File, ctx *gin.Context, folderName string) (string, error) {
	cld, _ := cloudinary.NewFromParams("ihdiannaja", "954945529412874", "7mFstMRVYEOlO784FGNo09mfk_4")

	newFileName := "img" + uuid.New().String()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: newFileName,
		Folder:   folderName,
	})

	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

func (f *fileRepository) SaveMultipleFiles(file *multipart.FileHeader, ctx *gin.Context, folderName string) (string, error) {
	cld, _ := cloudinary.NewFromParams("ihdiannaja", "954945529412874", "7mFstMRVYEOlO784FGNo09mfk_4")
	pathHold := "img" + uuid.New().String()
	fileInput, err := file.Open()
	if err != nil {
		return "", err
	}
	result, err := cld.Upload.Upload(ctx, fileInput, uploader.UploadParams{
		PublicID: pathHold,
		Folder:   folderName,
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
