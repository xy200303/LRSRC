package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
	"xiaoyun/backend/config"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/service/file_service"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
	"xiaoyun/backend/utils/oss"
	"xiaoyun/backend/utils/resp"
)

// UploadFile 上传文件，支持大文件上传
func UploadFile(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	c.Request.Header.Set("X-Content-Type-Options", "nosniff")
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(system_service.SysConfigMap.SysFileMaxSize<<20))
	// 解析表单数据
	err = c.Request.ParseMultipartForm(50 << 20) // 10 MB 内存缓存
	if err != nil {
		log.Println(system_service.SysConfigMap)
		resp.Err(c, err)
		return
	}
	// 获取文件对象
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		resp.Err(c, err)
		return
	}
	fileSize := header.Size
	if fileSize/(1024*2014) > int64(system_service.SysConfigMap.SysFileMaxSize) {
		resp.Err(c, fmt.Errorf("文件大小超过限制的%sMB大小", system_service.SysConfigMap))
		return
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			resp.Err(c, err)
			return
		}
	}(file)
	fileName := header.Filename

	var res *system_types.File
	if system_service.SysConfigMap.SysFileStorage == "local" {
		res, err = oss.LocalStorageFile(c, file, header)
	}
	if system_service.SysConfigMap.SysFileStorage == "aliyun" {
		res, err = oss.AliYunOssUploadFile(file, header)
	}
	if system_service.SysConfigMap.SysFileStorage == "tencent" {
		res, err = oss.TencentOssUploadFile(file, header)
	}
	if system_service.SysConfigMap.SysFileStorage == "huawei" {
		res, err = oss.HuaweiOssUploadFile(file, header)
	}
	if err != nil {
		log.Println(err)
		resp.Err(c, err)
		return
	}
	res.CreatedBy = auth.User.Username
	res.Name = fileName
	createFile, err := file_service.SaveFile(res)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, createFile)
}

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	fileId := c.Query("id")
	if fileId == "" {
		resp.Err(c, fmt.Errorf("无效的文件ID"))
		return
	}
	file, err1 := file_service.GetFile(fileId)
	if err1 != nil {
		resp.Err(c, err1)
		return
	}
	var err error
	var url string
	if file.FileType == "local" {
		filePath := filepath.Join(config.Config.Server.UploadDir, file.FileName)
		if !utils.CheckFileExists(filePath) {
			log.Println(filePath)
			resp.Err(c, fmt.Errorf("文件不存在或者被删除"))
			return
		}
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(file.Name)))
		c.Header("Content-Type", "application/octet-stream")
		if file.Extension == "svg" {
			c.Header("Content-Type", "image/svg+xml")
		}
		c.File(filePath)
		return
	}
	if file.FileType == "aliyun" {
		url, err = oss.AliYunOssGetDownloadURL(file.FileName)
	}
	if file.FileType == "tencent" {
		url, err = oss.TencentOssGetDownloadUrl(file.FileName)
	}
	if file.FileType == "huawei" {
		url, err = oss.HuaweiOssGetDownloadUrl(file.FileName)
	}
	if err != nil {
		log.Println(err)
		resp.Err(c, err)
		return
	}
	if url != "" {
		log.Print(url)
		c.Redirect(http.StatusFound, url)
	}
	resp.Err(c, fmt.Errorf("无效的文件信息"))
}

// UploadAvatar 上传文件
func UploadAvatar(c *gin.Context) {

}
