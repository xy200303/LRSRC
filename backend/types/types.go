package types

import (
	"gorm.io/gorm"
	"time"
)

// 系统配置表
type XqSystemConfig struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	UserRegister    bool      `json:"user_register"`
	UserDisplay     string    `json:"user_display"`
	MaxAttempts     int64     `json:"max_attempts"`
	LockoutDuration int64     `json:"lockout_duration"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

// Jwt配置表
type XqJwtConfig struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	JwtSecret  string    `json:"jwt_secret"`
	JwtExpires int64     `json:"jwt_expires"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 邮箱配置表
type XqEmailConfig struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	EmailHost     string    `json:"email_host"`
	EmailPort     int64     `json:"email_port"`
	EmailUser     string    `json:"email_user"`
	EmailPassword string    `json:"email_password"`
	EmailSender   string    `json:"email_sender"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

// 通知配置表
type XqNoticeConfig struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Type       int64     `json:"type"`
	Secret     string    `json:"secret"`
	Webhook    string    `json:"webhook"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// UserIntegral 用户积分
type UserIntegral struct {
	UserName   string    `json:"username"` // 用户名
	Type       string    `json:"type"`     // 类型
	Value      int32     `json:"value"`    // 值
	Reason     string    `json:"reason"`   // 原因
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 漏洞类型
type VulnType struct {
	gorm.Model
	SubType  string `json:"sub_type"`  // 子类型名称
	TypeName string `json:"type_name"` // 类型名称
	CateName string `json:"cate_name"` // 类别名称
}

// 漏洞审核
type VulnAudit struct {
	gorm.Model
	Type          string `json:"type"`
	VulnID        string `json:"vuln_id"`
	Status        uint8  `json:"status"`
	AuditOpinions string `json:"audit_opinions"`
	Username      string `json:"username"`
}

// 厂商信息
type Muna struct {
	gorm.Model
	Name    string `json:"name"`                 // 名称
	Desc    string `json:"desc"`                 // 描述
	Avatar  string `json:"avatar"`               // 头像
	Type    uint8  `json:"type"`                 // 类型或分类
	Capital uint32 `json:"capital"`              // 资本
	Domain  string `gorm:"unique" json:"domain"` // 主域名
}

// 厂商用户信息
type ManuUser struct {
	ManuID     uint      `gorm:"column:manu_id" json:"manu_id"`          // 厂商ID
	Email      string    `gorm:"column:email;unique;index" json:"email"` // 邮箱，唯一
	Password   string    `gorm:"column:password" json:"password"`        // 密码
	Status     uint8     `gorm:"column:status" json:"status"`            // 状态
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 漏洞评论信息
type VulnComment struct {
	VulnID     string    `gorm:"column:vuln_id" json:"vuln_id"`   // 漏洞ID
	Content    string    `gorm:"column:content" json:"content"`   // 评论内容
	Username   string    `gorm:"column:username" json:"username"` // 用户名称
	Status     uint8     `gorm:"column:status" json:"status"`     // 状态
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 文章评论信息
type ArticleComment struct {
	ArticleID  uint      `gorm:"column:article_id" json:"article_id"` // 文章ID
	Content    string    `gorm:"column:content" json:"content"`       // 评论内容
	Username   string    `gorm:"column:username" json:"username"`     // 用户名称
	Type       uint8     `gorm:"column:type" json:"type"`             // 评论类型
	Status     uint8     `gorm:"column:status" json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 组信息
type Group struct {
	gorm.Model
	Name     string `json:"name"`      // 名称
	ParentID uint   `json:"parent_id"` // 父组ID
	Desc     string `json:"desc"`      // 描述
}

// 等级信息
type Level struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`                 // 名称
	Desc        string `gorm:"column:desc" json:"desc"`                 // 描述
	MinIntegral uint32 `gorm:"column:min_integral" json:"min_integral"` // 最小积分
	MaxIntegral uint32 `gorm:"column:max_integral" json:"max_integral"` // 最大迭代次数
}

// 团队用户信息
type TeamUser struct {
	TeamID     uint      `json:"team_id"`   // 团队ID
	Username   string    `json:"username"`  // 用户ID
	UserType   uint8     `json:"user_type"` // 用户类型
	Status     uint8     `json:"status"`    // 状态
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 团队信息
type Team struct {
	gorm.Model
	Avatar   string `gorm:"column:avatar" json:"avatar"`     // 头像
	Team     string `gorm:"column:team" json:"team"`         // 团队名称
	Name     string `gorm:"column:name" json:"name"`         // 名称
	Desc     string `gorm:"column:desc" json:"desc"`         // 描述
	Username string `gorm:"column:username" json:"username"` // 用户名称
}

// 团队申请信息
type TeamApp struct {
	gorm.Model
	Status      uint8  `gorm:"column:status" json:"status"`             // 状态
	Username    string `gorm:"column:username" json:"username"`         // 用户名称
	AppReason   string `gorm:"column:app_reason" json:"app_reason"`     // 申请原因
	ReplyReason string `gorm:"column:reply_reason" json:"reply_reason"` // 回复原因
}

// 证书模板
type CertTemp struct {
	gorm.Model
	Username string `json:"username"`
	UserType uint8  `json:"user_type"`
	Type     uint8  `json:"type"`
	FileID   string `json:"file_id"`
}

// 用户证书
type UserCert struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	TempID     uint      `json:"temp_id"`
	FileID     string    `json:"file_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 文章操作
type ArticleOperation struct {
	ArticleID  string    `json:"article_id"`
	Username   string    `json:"username"`
	Type       string    `json:"type"`
	Value      int32     `json:"value"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 漏洞类型表
type XqVulnType struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// Vulnerability 漏洞表
type XqVulnerability struct {
	ID                     string    `gorm:"primaryKey" json:"id"`
	UserID                 uint64    `json:"user_id"`
	CVE                    string    `json:"cve"`
	NVD                    string    `json:"nvd"`
	EDB                    string    `json:"edb"`
	CNNVD                  string    `json:"cnnvd"`
	CNVD                   string    `json:"cnvd"`
	VulnName               string    `json:"vuln_name"`
	VulnTypeID             uint64    `json:"vuln_type_id"`
	VulnType               string    `json:"vuln_type"`
	VulnLevel              string    `json:"vuln_level"`
	CVSS                   float64   `json:"cvss"`
	Description            string    `json:"description"`
	AffectedProduct        string    `json:"affected_product"`
	AffectedProductVersion string    `json:"affected_product_version"`
	FofaQuery              string    `json:"fofa_query"`
	ZoomEyeQuery           string    `json:"zoomeye_query"`
	QuakeQuery             string    `json:"quake_query"`
	HunterQuery            string    `json:"hunter_query"`
	GoogleQuery            string    `json:"google_query"`
	ShodanQuery            string    `json:"shodan_query"`
	CensysQuery            string    `json:"censys_query"`
	GreynoiseQuery         string    `json:"greynoise_query"`
	Poc                    string    `json:"poc"`
	PocType                string    `json:"poc_type"`
	Exp                    string    `json:"exp"`
	ExpType                string    `json:"exp_type"`
	RepairSuggestion       string    `json:"repair_suggestion"`
	AttachmentID           string    `json:"attachment_id"`
	AttachmentName         string    `json:"attachment_name"`
	Submitter              string    `json:"submitter"`
	IsPublic               bool      `json:"is_public"`
	Status                 int64     `json:"status"`
	ReviewComments         string    `json:"review_comments"`
	CreateTime             time.Time `json:"create_time"`
	UpdateTime             time.Time `json:"update_time"`
}

//

// Lockip 锁定IP表
type XqLockip struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	ClientIP     string     `json:"client_ip"`
	LockoutUntil *time.Time `json:"lockout_until"`
	Status       int64      `json:"status"`
	CreateTime   time.Time  `json:"create_time"`
	UpdateTime   time.Time  `json:"update_time"`
}

// 附件表
type XqAttachment struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserID     uint64    `json:"user_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Data       []byte    `json:"data"`
	Status     int64     `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 用户ranking明细表
type XqRankingDetail struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	UserID     uint64    `json:"user_id"`
	VulnID     string    `json:"vuln_id"`
	Ranking    int64     `json:"ranking"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 评分规则表
type XqScoreRule struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Type        int64     `json:"type"`
	Rule        string    `json:"rule"`
	Score       float64   `json:"score"`
	Coefficient float64   `json:"coefficient"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

// 验证码表
type XqVerifyCode struct {
	ID          uint64 `gorm:"primaryKey"`
	Email       string
	Code        string
	CreateTime  time.Time
	UpdateTime  time.Time
	ExpiredTime time.Time
}

type SystemConfigData struct {
	EmailConfig  XqEmailConfig  `json:"emailconf"`
	JwtConfig    XqJwtConfig    `json:"jwtconf"`
	NoticeConfig XqNoticeConfig `json:"noticeconf"`
	SysConfig    XqSystemConfig `json:"sysconf"`
}
