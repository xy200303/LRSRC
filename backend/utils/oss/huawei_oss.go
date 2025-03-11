package oss

import (
	"fmt"
	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"log"
	"mime/multipart"
	"strings"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
)

var (
	huaweiOssClient  *obs.ObsClient
	huaweiBucketName = ""
)

// 初始化对象
func InitHuaweiOss(sysConfigMap *system_types.SysConfigMap) {
	AccessKeyId := sysConfigMap.SysHuaweiOssAk
	AccessKeySecret := sysConfigMap.SysHuaweiOssSk
	huaweiBucketName = sysConfigMap.SysHuaweiOssBucket

	endPoint := sysConfigMap.SysHuaweiOssEndpoint
	var err error
	huaweiOssClient, err = obs.New(AccessKeyId, AccessKeySecret, endPoint, obs.WithSignature(obs.SignatureObs))
	if err != nil {
		log.Println(err)
	}
}

// HuaweiOssUploadFile 上传文件
func HuaweiOssUploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*system_types.File, error) {
	extension, err := utils.GetFileExtensionByName(fileHeader.Filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempFileName := utils.GenerateTempFileName("temp", "."+extension)
	input := &obs.PutObjectInput{}
	// 指定存储桶名称
	input.Bucket = huaweiBucketName
	input.Key = tempFileName
	input.Body = file
	output, err := huaweiOssClient.PutObject(input)
	if err != nil {
		return nil, err
	}
	md5File := strings.ToLower(strings.Replace(output.ETag, "\"", "", -1))
	newFileName := utils.GetCurrentDate() + "/" + strings.ToLower(strings.Replace(output.ETag, "\"", "", -1)) + "." + extension
	_, err = HuaweiOssRenameFile(tempFileName, newFileName)
	if err != nil {
		return nil, err
	}

	res := &system_types.File{
		FileID:    md5File,
		Extension: extension,
		FileType:  "huawei",
		FileName:  newFileName,
	}
	return res, nil
}

// HuaweiOssCopyFile 复制文件
func HuaweiOssCopyFile(srcObjectName, destObjectName string) (*obs.CopyObjectOutput, error) {
	input := &obs.CopyObjectInput{}
	// 指定源桶名称
	input.Bucket = huaweiBucketName
	// 指定源对象，此处以 example/objectname 为例。
	input.Key = destObjectName
	// 指定复制源桶名，此处以sourcebucketname为例。
	input.CopySourceBucket = huaweiBucketName
	// 指定复制源桶名下的指定源对象名，此处以sourceobjectkey为例。
	input.CopySourceKey = srcObjectName
	// 复制对象
	output, err := huaweiOssClient.CopyObject(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// HuaweiOssDeleteFile 删除文件
func HuaweiOssDeleteFile(fileName string) (*obs.DeleteObjectOutput, error) {
	input := &obs.DeleteObjectInput{}
	// 指定存储桶名称
	input.Bucket = huaweiBucketName
	// 指定删除对象，此处以 example/objectname 为例。
	input.Key = fileName
	// 删除对象
	output, err := huaweiOssClient.DeleteObject(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// HuaweiOssRenameFile 重命名文件
func HuaweiOssRenameFile(srcObjectName, destObjectName string) (*obs.DeleteObjectOutput, error) {
	_, err := HuaweiOssCopyFile(srcObjectName, destObjectName)
	if err != nil {
		return nil, err
	}
	output, err := HuaweiOssDeleteFile(srcObjectName)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// HuaweiOssGetDownloadUrl 获取文件下载地址
func HuaweiOssGetDownloadUrl(fileName string) (string, error) {
	putObjectInput := &obs.CreateSignedUrlInput{}
	putObjectInput.Method = obs.HttpMethodGet
	putObjectInput.Bucket = huaweiBucketName
	putObjectInput.Key = fileName
	putObjectInput.Expires = 3600
	// 生成上传对象的带授权信息的URL
	output, err := huaweiOssClient.CreateSignedUrl(putObjectInput)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return output.SignedUrl, nil
}
