package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
	"xiaoyun/backend/config"
)

// 初始化Redis数据库
func InitRedis() {
	database, err := strconv.Atoi(config.Config.Redis.DB)
	if err != nil {
		log.Fatalf("Failed to strconv: %v", err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + config.Config.Redis.Port,
		Password: config.Config.Redis.Password,
		DB:       database,
	})
	// 测试连接是否成功
	ctx := context.Background()
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}

// 存储键值
func SetRedisValue(key string, value string, expiration int) error {
	ctx := context.Background()
	err := RDB.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		log.Printf("Failed to set: %v", err)
	}
	return err
}
func GetRedisValue(key string) (string, error) {
	ctx := context.Background()
	val, err := RDB.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Failed to get: %v", err)
		return "", err
	}
	return val, nil
}

// GetRedisKeyTTL 获取指定键的剩余过期时间
func GetRedisKeyTTL(key string) (int64, error) {
	ctx := context.Background()
	ttl, err := RDB.TTL(ctx, key).Result()
	if err != nil {
		log.Printf("Failed to get TTL: %v", err)
		return -1, err
	}
	// ttl 返回的是 time.Duration 类型，需要转换为秒
	// 如果键不存在，ttl 会是 -2 * time.Second；如果未设置过期时间，ttl 会是 -1 * time.Second
	seconds := int64(ttl / time.Second)
	if seconds == -2 {
		log.Println("键不存在")
		return seconds, nil
	} else if seconds == -1 {
		log.Println("键没有设置过期时间")
		return seconds, nil
	}
	return seconds, nil
}

// DeleteRedisValue 删除键值
func DeleteRedisValue(key string) error {
	ctx := context.Background()
	err := RDB.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Failed to delete: %v", err)
	}
	return err
}

// 定义Redis前缀
func setRedisKeyWithExpire(key string, value string, expire int) error {
	return SetRedisValue(key, value, expire)
}

func getRedisKey(key string) (string, error) {
	return GetRedisValue(key)
}

func delRedisKey(key string) error {
	return DeleteRedisValue(key)
}

func RedisSetRegisterEmailCode(email string, code string, duration int) error {
	return setRedisKeyWithExpire("email_code_register_"+email, code, duration)
}

func RedisGetRegisterEmailCode(email string) (string, error) {
	return getRedisKey("email_code_register_" + email)
}

func RedisDelRegisterEmailCode(email string) error {
	return delRedisKey("email_code_register_" + email)
}

func RedisSetForgetEmailCode(email string, code string, duration int) error {
	return setRedisKeyWithExpire("email_code_forget_"+email, code, duration)
}

// RedisGetForgetEmailCodeTTL 获取忘记密码验证码TTL
func RedisGetForgetEmailCodeTTL(email string) (int64, error) {
	return GetRedisKeyTTL("email_code_forget_" + email)
}

// RedisGetRegisterEmailCodeTTL 获取用户注册验证码TTL
func RedisGetRegisterEmailCodeTTL(email string) (int64, error) {
	return GetRedisKeyTTL("email_code_register_" + email)
}

func RedisGetForgetEmailCode(email string) (string, error) {
	return getRedisKey("email_code_forget_" + email)
}

func RedisDelForgetEmailCode(email string) error {
	return delRedisKey("email_code_forget_" + email)
}

func RedisSetUserToken(username string, token string) error {
	return setRedisKeyWithExpire("token_"+username, token, 2*60*60)
}

func RedisGetUserToken(username string) (string, error) {
	return getRedisKey("token_" + username)
}

func RedisDelUserToken(username string) error {
	return delRedisKey("token_" + username)
}
func RedisSetJwtSecret(jwtSecret string) error {
	return setRedisKeyWithExpire("jwt_secret", jwtSecret, 0)
}

func RedisGetJwtSecret() (string, error) {
	return getRedisKey("jwt_secret")
}
