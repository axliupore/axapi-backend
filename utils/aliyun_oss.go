package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/axliupore/axapi/axapi-backend/global"
	"go.uber.org/zap"
	"mime/multipart"
)

// 创建桶的实例
func newBucket() (*oss.Bucket, error) {
	client, err := oss.New(global.Config.AliyunOSS.Endpoint, global.Config.AliyunOSS.AccessKeyId, global.Config.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.Config.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// UploadFile 上传文件
func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	bucket, err := newBucket()
	if err != nil {
		global.Log.Error("func NewBucket() Failed", zap.Error(err))
		return "", err
	}
	file, err := fileHeader.Open()
	if err != nil {
		global.Log.Error("file open error", zap.Error(err))
		return "", err
	}
	// 获取扩展名
	fileExt := GetFileExt(fileHeader)

	objectKey := fmt.Sprintf("%s/%s%s", global.Config.AliyunOSS.BasePath, CreateUUID(), fileExt)
	if err := bucket.PutObject(objectKey, file); err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s", global.Config.AliyunOSS.BucketUrl, objectKey)
	return url, nil
}

// DeleteFile 删除文件
func DeleteFile(key string) error {
	bucket, err := newBucket()
	if err != nil {
		global.Log.Error("func NewBucket() Failed", zap.Error(err))
		return err
	}
	// 删除单个文件, objectName 表示删除 oss 文件所需的完整路径
	err = bucket.DeleteObject(key)
	if err != nil {
		global.Log.Error("func DeleteFile() Failed", zap.Error(err))
		return err
	}
	return nil
}
