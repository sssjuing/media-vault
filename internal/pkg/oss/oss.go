package oss

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sort"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

var minioClient *minio.Client
var bucketName string

func init() {
	config := config.GetConfig()
	host := config.GetString("minio.host")
	port := config.GetString("minio.port")
	endpoint := fmt.Sprintf("%s:%s", host, port)
	username := config.GetString("minio.username")
	password := config.GetString("minio.password")
	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(username, password, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	minioClient = client
	bucketName = config.GetString("minio.bucket_name")
}

type File struct {
	Key          string    `json:"key"`
	Name         string    `json:"name"`
	Ftype        string    `json:"ftype"`
	Size         int64     `json:"size"`
	Url          *string   `json:"url"`
	LastModified time.Time `json:"last_modified"`
}

func GetFiles(folderName string) ([]*File, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	publicUrl := config.GetMinioPublicUrl()
	result := make([]*File, 0, 10)
	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: folderName, Recursive: false})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return nil, object.Err
		}
		ftype := "directory"
		var urlPtr *string
		if object.StorageClass == "STANDARD" {
			ftype = "file"
			urlPtr = new(string)
			*urlPtr = publicUrl + object.Key
		}
		result = append(result, &File{
			Key:          object.Key,
			Name:         object.Key,
			Ftype:        ftype,
			Size:         object.Size,
			Url:          urlPtr,
			LastModified: object.LastModified,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].LastModified.Equal(time.Time{}) {
			result[i].LastModified = time.Now()
		}
		if result[j].LastModified.Equal(time.Time{}) {
			result[j].LastModified = time.Now()
		}
		return result[i].LastModified.After(result[j].LastModified)
	})
	return result, nil
}

type UploadInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Url  string `json:"url"`
}

func UploadFile(file *multipart.File, objectPath, contentType string) (*UploadInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	publicUrl := config.GetMinioPublicUrl()
	defer cancel()
	info, err := minioClient.PutObject(ctx, bucketName, objectPath, *file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, err
	}
	return &UploadInfo{Path: objectPath, Size: info.Size, Url: publicUrl + objectPath}, nil
}

func UploadLocalFile(ctx context.Context, objectPath, filePath string) (*UploadInfo, error) {
	mt, _ := mimetype.DetectFile(filePath)
	info, err := minioClient.FPutObject(ctx, bucketName, objectPath, filePath, minio.PutObjectOptions{ContentType: mt.String()})
	if err != nil {
		return nil, err
	}
	return &UploadInfo{Path: objectPath, Size: info.Size}, nil
}
