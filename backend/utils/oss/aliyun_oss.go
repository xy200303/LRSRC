package oss

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"log"
	"mime/multipart"
	"strings"
	"time"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
)

var (
	aliYunOssClient  *oss.Client
	aliYunBucketName = ""
	aliYunRegin      = ""
)

// 初始化对象
func InitAliYunOss(sysConfigMap *system_types.SysConfigMap) {
	AccessKeyId := sysConfigMap.SysAliyunOssAk
	AccessKeySecret := sysConfigMap.SysAliyunOssSk
	aliYunRegin = sysConfigMap.SysAliyunOssRegion
	aliYunBucketName = sysConfigMap.SysAliyunOssBucket

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(AccessKeyId, AccessKeySecret)).
		WithRegion(aliYunRegin)
	// 创建OSS客户端
	aliYunOssClient = oss.NewClient(cfg)
}

// AliYunOssUploadFile 阿里云OSS上传文件
func AliYunOssUploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*system_types.File, error) {
	extension, err := utils.GetFileExtensionByName(fileHeader.Filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempFileName := utils.GenerateTempFileName("temp", "."+extension)

	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(aliYunBucketName),
		Key:    oss.Ptr(tempFileName),
		Body:   file,
	}
	defer file.Close()
	// 执行上传对象的请求
	result, err := aliYunOssClient.PutObject(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	md5File := strings.ToLower(strings.Replace(*result.ETag, "\"", "", -1))
	newFileName := utils.GetCurrentDate() + "/" + md5File + "." + extension
	err = AliYunOssRenameFile(tempFileName, newFileName)
	if err != nil {
		return nil, err
	}
	res := &system_types.File{
		FileID:    md5File,
		Extension: extension,
		FileType:  "aliyun",
		FileName:  newFileName,
	}
	return res, nil
}

// AliYunOssCopyFile 复制文件
func AliYunOssCopyFile(srcObjectName, destObjectName string) (*oss.CopyResult, error) {
	// 创建文件拷贝器
	c := aliYunOssClient.NewCopier()
	// 构建拷贝对象的请求
	copyRequest := &oss.CopyObjectRequest{
		Bucket:       oss.Ptr(aliYunBucketName), // 目标存储空间名称
		Key:          oss.Ptr(destObjectName),   // 目标对象名称
		SourceKey:    oss.Ptr(srcObjectName),    // 源对象名称
		SourceBucket: oss.Ptr(aliYunBucketName), // 源存储空间名称
		StorageClass: oss.StorageClassStandard,  // 指定存储类型为标准类型
	}
	// 执行拷贝对象的操作
	output, err := c.Copy(context.TODO(), copyRequest)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// AliYunOssRenameFile 在OSS中重命名对象
func AliYunOssRenameFile(srcObjectName, destObjectName string) error {
	_, err := AliYunOssCopyFile(srcObjectName, destObjectName)
	if err != nil {
		return err
	}
	err = AliYunOssDeleteFile(srcObjectName)
	if err != nil {
		return err
	}
	return nil
}

// AliYunOssDeleteFile 删除文件
func AliYunOssDeleteFile(fileName string) error {
	// 构建删除对象的请求
	deleteRequest := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(aliYunBucketName), // 存储空间名称
		Key:    oss.Ptr(fileName),         // 要删除的对象名称
	}
	// 执行删除对象的操作
	_, err := aliYunOssClient.DeleteObject(context.TODO(), deleteRequest)
	if err != nil {
		return err
	}
	return nil
}

// AliYunOssGetDownloadURL 根据用户名生成下载地址（预签名URL）
func AliYunOssGetDownloadURL(fileName string) (string, error) {
	result, err := aliYunOssClient.Presign(context.TODO(), &oss.GetObjectRequest{
		Bucket: oss.Ptr(aliYunBucketName),
		Key:    oss.Ptr(fileName),
	}, oss.PresignExpires(10*time.Minute))
	if err != nil {
		return "", err
	}
	return result.URL, err
}
