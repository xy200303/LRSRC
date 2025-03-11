package oss

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
)

var (
	TencentOssClient *cos.Client
	BaseUrl          = ""
	TencentSecretId  = ""
	TencentSecretKey = ""
)

func InitTencentOss(sysConfigMap *system_types.SysConfigMap) {
	BaseUrl = sysConfigMap.SysTencentOssBaseUrl
	TencentSecretKey = sysConfigMap.SysTencentOssAk
	TencentSecretId = sysConfigMap.SysHuaweiOssSk

	u, _ := url.Parse("https://" + BaseUrl)
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	TencentOssClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  TencentSecretId,
			SecretKey: TencentSecretKey,
		},
	})
}

// TencentOssUploadFile 上传文件
func TencentOssUploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*system_types.File, error) {
	extension, err := utils.GetFileExtensionByName(fileHeader.Filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempFileName := utils.GenerateTempFileName("temp", "."+extension)
	output, err := TencentOssClient.Object.Put(context.Background(), tempFileName, file, nil)
	if err != nil {
		return nil, err
	}
	md5File := strings.ToLower(strings.Replace(output.Header.Get("ETag"), "\"", "", -1))
	newFileName := utils.GetCurrentDate() + "/" + strings.ToLower(strings.Replace(output.Header.Get("ETag"), "\"", "", -1)) + "." + extension
	_, err = TencentOssRenameFile(tempFileName, newFileName)
	if err != nil {
		return nil, err
	}
	res := &system_types.File{
		FileID:    md5File,
		Extension: extension,
		FileType:  "tencent",
		FileName:  newFileName,
	}
	return res, nil
}

// 复制文件
func TencentOssCopyFile(srcObjectName, destObjectName string) (*cos.ObjectCopyResult, *cos.Response, error) {
	sourceURL := fmt.Sprintf(BaseUrl+"/%s", srcObjectName)
	// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	// opt := &cos.ObjectCopyOptions{}
	result, res, err := TencentOssClient.Object.Copy(context.Background(), destObjectName, sourceURL, nil)
	if err != nil {
		return result, res, err
	}
	return result, res, nil
}

// TencentOssDeleteFile 删除文件
func TencentOssDeleteFile(fileNme string) (*cos.Response, error) {
	res, err := TencentOssClient.Object.Delete(context.Background(), fileNme)
	if err != nil {
		return res, err
	}
	return res, nil
}

// 重命名文件
func TencentOssRenameFile(srcObjectName, destObjectName string) (*cos.Response, error) {
	_, res, err := TencentOssCopyFile(srcObjectName, destObjectName)
	if err != nil {
		return nil, err
	}
	res, err = TencentOssDeleteFile(srcObjectName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TencentOssGetDownloadUrl 获取文件下载地址
func TencentOssGetDownloadUrl(fileName string) (string, error) {
	//ourl := TencentOssClient.Object.GetObjectURL(fileName)
	//return ourl.String()
	presignedURL, err := TencentOssClient.Object.GetPresignedURL(context.Background(), http.MethodGet, fileName, TencentSecretId, TencentSecretKey, time.Hour, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
