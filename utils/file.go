package utils

import (
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// CreateUUID 创建 uuid
func CreateUUID() string {
	return uuid.NewString()
}

// GetFileExt 获取文件扩展名
func GetFileExt(fileHeader *multipart.FileHeader) string {
	// 使用 filepath.Ext 函数获取文件扩展名
	fileName := fileHeader.Filename
	fileExtension := filepath.Ext(fileName)
	return fileExtension
}

// DirExistOrNot 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return false
	}
	return true
}

// SaveFileLocal 保存文件到文件, 返回的是相对路径
func SaveFileLocal(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		global.Log.Error("SaveFileLocal error open()", zap.Error(err))
		return "", err
	}
	// 文件扩展名
	fileExt := GetFileExt(fileHeader)
	path, _ := os.Getwd()
	// 存放文件的目录
	dirPath := filepath.Join(path, global.Config.Server.FilePath)
	if !DirExistOrNot(dirPath) {
		CreateDir(dirPath)
	}
	id := CreateUUID()
	filePath := filepath.Join(dirPath, id+fileExt)
	fileName := filepath.Join(global.Config.Server.FilePath, id+fileExt)
	content, err := io.ReadAll(file)
	if err != nil {
		global.Log.Error("SaveFileLocal ReadAll error", zap.Error(err))
		return "", err
	}
	err = os.WriteFile(filePath, content, 0666)
	if err != nil {
		global.Log.Error("SaveFileLocal WriteFile error", zap.Error(err))
		return "", err
	}
	return strings.Replace(fileName, "\\", "/", -1), nil
}
