package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成随机Jwt secret
func GenerateRandomJwtSecret() string {
	randomBytes := make([]byte, 32)
	_, _ = rand.Read(randomBytes)
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// generateRandomChars 生成包含特殊字符的随机密码
func GenerateRandomChars(length int64, chartype int64) (string, error) {
	const numChars = "0123456789"
	const specChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	const lowerChars = "abcdefghijklmnopqrstuvwxyz"
	const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var chars string
	//const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// 纯数字
	if chartype == 1 {
		chars = numChars
	}
	// 纯小写字母
	if chartype == 2 {
		chars = lowerChars
	}
	// 纯大写字母
	if chartype == 3 {
		chars = upperChars
	}
	// 字母+数字
	if chartype == 4 {
		chars = numChars + lowerChars + upperChars
	}
	// 字母+数字+特殊符号
	if chartype == 5 {
		chars = numChars + lowerChars + upperChars + specChars
	}
	for i, b := range randomBytes {
		randomBytes[i] = chars[int(b)%len(chars)]
	}
	return string(randomBytes), nil
}

func GenJWTToken(username string, role uint8, expires int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role_id":  role,
		"exp":      expires,
	})
	return token.SignedString([]byte(secret))
}

func GenPasswordHash(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(passwordHash)
}

func IsHashEqual(hash string, passwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)); err != nil {
		return false
	}
	return true
}

func DecodeJWTToken(tokenString string, secret string) (jwt.MapClaims, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	// 提取 Claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func GenerateUniqueID() string {
	return uuid.New().String()
}

func GetSystemUsage() (int64, int64, int64) {
	cpuPercent, _ := cpu.Percent(0, false)
	cpuUsage := int64(cpuPercent[0])

	memInfo, _ := mem.VirtualMemory()
	memUsage := int64(memInfo.UsedPercent)

	diskUsageInfo, _ := disk.Usage("/")
	diskUsage := int64(diskUsageInfo.UsedPercent)
	return cpuUsage, memUsage, diskUsage
}

// 生成六位数验证码
func GenRandomNumber(length uint8) string {
	// 使用当前时间的Unix时间戳纳秒数作为随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建一个字符串builder来高效地拼接数字
	var sb strings.Builder
	// 生成指定长度的随机数字
	for i := uint8(0); i < length; i++ {
		// 生成一个0到9之间的随机整数
		num := rand.Intn(10)
		// 将整数转换为字符串并追加到builder中
		sb.WriteString(strconv.Itoa(num))
	}
	// 将builder转换为字符串并返回
	return sb.String()
}

// GetToken 提取Token
func GetToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	// 检查是否是 Bearer Token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}
	return tokenParts[1]
}
