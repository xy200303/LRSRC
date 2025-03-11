package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CopyFile 复制文件
func CopyFile(srcFile string, destFile string) error {
	s, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	defer s.Close()
	_, err = io.Copy(f, s)
	return err
}
func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		// 可能是权限问题等其他错误
		fmt.Println("Error:", err)
		return false
	}
	return true
}

// CalFileMD5 计算文件的MD5哈希值，支持大文件
func CalFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 创建一个新的MD5哈希实例
	hash := md5.New()
	// 定义缓冲区大小
	const bufferSize = 1 * 1024 * 1024
	// 创建一个缓冲区
	buffer := make([]byte, bufferSize)
	// 逐块读取文件内容并写入哈希实例
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			// 将读取的内容写入哈希实例
			if _, err := hash.Write(buffer[:n]); err != nil {
				return "", err
			}
		}
		// 如果到达文件末尾或发生错误，则退出循环
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}
	// 获取哈希结果并转换为16进制字符串
	md5Sum := hash.Sum(nil)
	return fmt.Sprintf("%x", md5Sum), nil
}

// GenerateTempFileName 生成一个唯一的临时文件名
func GenerateTempFileName(prefix string, suffix string) string {
	// 获取当前时间戳
	timestamp := time.Now().Format("20060102150405")

	// 生成随机数
	randNum := rand.Intn(10000)

	// 生成 UUID
	uuidStr := uuid.New().String()

	// 组合文件名
	filename := fmt.Sprintf("%s_%s_%d_%s%s", prefix, timestamp, randNum, strings.Replace(uuidStr, "-", "", -1), suffix)

	return filename
}

// GetFileExtensionByName GetFileExtension 获取文件后缀
func GetFileExtensionByName(filePath string) (string, error) {
	// 获取文件的扩展名
	ext := filepath.Ext(filePath)
	if ext == "" {
		return "", fmt.Errorf("no extension found for file: %s", filePath)
	}
	// 去掉前面的点
	return ext[1:], nil // 如果你希望保留点，可以去掉 [1:]
}

func GetFileExtension(file *os.File) (string, error) {
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %v", err)
	}
	// 获取文件名
	fileName := fileInfo.Name()
	name, err := GetFileExtensionByName(fileName)
	if err != nil {
		return "", err
	}
	return name, nil
}

func main() {
	file, err := os.Open("C:\\Users\\34834\\Desktop\\WebProject\\xiaoyun\\uploads\\Apifox.exe")
	if err != nil {
		log.Fatalf("failed to open local file %v", err)
	}
	defer file.Close()
	extension, err := GetFileExtension(file)
	if err != nil {
		return
	}
	log.Print(extension)
}
