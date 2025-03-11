package system_types

import "time"

// SysConfigMap 配置文件映射
type SysConfigMap struct {
	// 文件存储方式
	SysFileStorage string `json:"sys_file_storage" yaml:"sys_file_storage" name:"文件存储方式"`
	SysFileMaxSize int    `json:"sys_file_max_size" yaml:"sys_file_max_size" name:"文件最大大小"`

	// OSS配置
	SysAliyunOssAk       string `json:"sys_aliyun_oss_ak" yaml:"sys_aliyun_oss_ak" name:"阿里云OSS AccessKey"`
	SysAliyunOssSk       string `json:"sys_aliyun_oss_sk" yaml:"sys_aliyun_oss_sk" name:"阿里云OSS SecretKey"`
	SysAliyunOssBucket   string `json:"sys_aliyun_oss_bucket" yaml:"sys_aliyun_oss_bucket" name:"阿里云OSS Bucket名称"`
	SysAliyunOssRegion   string `json:"sys_aliyun_oss_regin" yaml:"sys_aliyun_oss_regin" name:"阿里云OSS 区域"`
	SysTencentOssAk      string `json:"sys_tencent_oss_ak" yaml:"sys_tencent_oss_ak" name:"腾讯云OSS AccessKey"`
	SysTencentOssSk      string `json:"sys_tencent_oss_sk" yaml:"sys_tencent_oss_sk" name:"腾讯云OSS SecretKey"`
	SysTencentOssBaseUrl string `json:"sys_tencent_oss_base_url" yaml:"sys_tencent_oss_base_url" name:"腾讯云OSS 基础URL"`
	SysHuaweiOssAk       string `json:"sys_huawei_oss_ak" yaml:"sys_huawei_oss_ak" name:"华为云OSS AccessKey"`
	SysHuaweiOssSk       string `json:"sys_huawei_oss_sk" yaml:"sys_huawei_oss_sk" name:"华为云OSS SecretKey"`
	SysHuaweiOssBucket   string `json:"sys_huawei_oss_bucket" yaml:"sys_huawei_oss_bucket" name:"华为云OSS Bucket名称"`
	SysHuaweiOssEndpoint string `json:"sys_huawei_oss_endpoint" yaml:"sys_huawei_oss_endpoint" name:"华为云OSS 终端节点"`

	// 邮箱设置
	SysSmtpHost         string `json:"sys_smtp_host" yaml:"sys_smtp_host" name:"SMTP 服务器地址"`
	SysSmtpPort         int    `json:"sys_smtp_port" yaml:"sys_smtp_port" name:"SMTP 服务器端口"`
	SysSmtpUsername     string `json:"sys_smtp_username" yaml:"sys_smtp_username" name:"SMTP 用户名"`
	SysSmtpPassword     string `json:"sys_smtp_password" yaml:"sys_smtp_password" name:"SMTP 密码"`
	SysSmtpSender       string `json:"sys_smtp_sender" yaml:"sys_smtp_sender" name:"SMTP 发件人"`
	SysSmtpCaptchaTitle string `json:"sys_smtp_captcha_title" yaml:"sys_smtp_captcha_title" name:"邮箱验证码标题"`

	// 注册和登录配置
	SysRegisterEnable bool `json:"sys_register_enable" yaml:"sys_register_enable" name:"是否启用注册功能"`

	// AI大模型配置
	SysAiType         string `json:"sys_ai_type" yaml:"sys_ai_type" name:"大模型类型"`
	SysAiBaseUrl      string `json:"sys_ai_base_url" yaml:"sys_ai_base_url" name:"大模型接口地址"`
	SysAiApiKey       string `json:"sys_ai_api_key" yaml:"sys_ai_api_key" name:"大模型API_KEY"`
	SysAiModel        string `json:"sys_ai_model" yaml:"sys_ai_model" name:"模型名称"`
	SysAiMaxPerTokens int    `json:"sys_ai_max_per_tokens" yaml:"sys_ai_max_per_tokens" name:"单词对话长度限制"`
}

// SysConfig 系统配置表
type SysConfig struct {
	SysKey    string    `gorm:"primaryKey" json:"sys_key"`
	Value     string    `json:"value"`
	Type      string    `gorm:"default:system" json:"type"`
	SysGroup  string    `json:"sys_group"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Config 结构体定义
type Config struct {
	Database struct {
		Type       string `yaml:"type"`
		Connection struct {
			Host     string `yaml:"host"`
			Port     int64  `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
			Charset  string `yaml:"charset"`
			File     string `yaml:"file"`
		} `yaml:"connection"`
	} `yaml:"database"`
	Server struct {
		Mode         string `yaml:"mode"`
		StartMode    string `yaml:"start_mode"`
		Host         string `yaml:"host"`
		Port         int64  `yaml:"port"`
		ReadTimeout  int64  `yaml:"read_timeout"`
		WriteTimeout int64  `yaml:"write_timeout"`
		WebPath      string `yaml:"web_path"`
		StaticUrl    string `yaml:"static_url"`
		AllowOrigins string `yaml:"allow_origins"`
		AllowMethods string `yaml:"allow_methods"`
		AllowHeaders string `yaml:"allow_headers"`
		UploadDir    string `yaml:"upload_dir"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"log"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	SysConfigMap SysConfigMap `yaml:"sys_config_map"`
}
