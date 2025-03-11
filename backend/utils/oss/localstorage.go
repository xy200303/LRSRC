package oss

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"xiaoyun/backend/config"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
)

// LocalStorageFile 存储文件到本地
func LocalStorageFile(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) (*system_types.File, error) {
	extension, err := utils.GetFileExtensionByName(fileHeader.Filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempFileName := utils.GenerateTempFileName("temp", "."+extension)
	// 创建目标文件路径
	tempFilePath := filepath.Join(os.TempDir(), tempFileName)
	err = c.SaveUploadedFile(fileHeader, tempFilePath)
	//写入文件
	if err != nil {
		log.Println(err)
		return nil, err
	}
	md5File, err := utils.CalFileMD5(tempFilePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newDateDir := utils.GetCurrentDate()
	newFileName := md5File + "." + extension
	err = os.MkdirAll(filepath.Join(config.Config.Server.UploadDir, newDateDir), os.ModePerm)
	if err != nil {
		fmt.Println("创建目录时出错:", err)
		return nil, err
	}
	newFile := filepath.Join(config.Config.Server.UploadDir, newDateDir, newFileName)
	// 使用 os.Rename 移动文件
	err = os.Rename(tempFilePath, newFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res := &system_types.File{
		FileID:    md5File,
		Extension: extension,
		FileType:  "local",
		FileName:  filepath.Join(newDateDir, newFileName),
	}
	return res, nil
}
