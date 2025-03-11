package models

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/natefinch/lumberjack"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os/user"
	"xiaoyun/backend/config"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/types/vuln_types"
	"xiaoyun/backend/utils"
)

var (
	db  *gorm.DB
	RDB *redis.Client
)

// 初始化SQL数据库
func InitSQL(gormLogger logger.Interface) {
	dsn := generateDSN(config.Config)
	switch config.Config.Database.Type {
	case "mysql":
		// 检测和创建数据库
		dbSQL, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.Database.Connection.User, config.Config.Database.Connection.Password, config.Config.Database.Connection.Host, config.Config.Database.Connection.Port))
		if err != nil {
			log.Fatalf("Error opening database connection:mysql %v", err)
		}
		defer dbSQL.Close()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
	case "postgres":
		dbSQL, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%d/", config.Config.Database.Connection.User, config.Config.Database.Connection.Password, config.Config.Database.Connection.Host, config.Config.Database.Connection.Port))
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
		defer dbSQL.Close()
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
	case "sqlite":
		var err error
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
	case "sqlserver":
		dbSQL, err := sql.Open("sqlserver", fmt.Sprintf("server=%s,%d;user_types id=%s;password=%s;", config.Config.Database.Connection.Host, config.Config.Database.Connection.Port, config.Config.Database.Connection.User, config.Config.Database.Connection.Password))
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
		defer dbSQL.Close()
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}
	default:
		log.Fatalf("Unsupported database type: %s", config.Config.Database.Type)
	}
}
func checkTablesExist(db *gorm.DB, tables []interface{}) bool {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			return false
		}
	}
	return true
}

// 初始化密码
func initAdminPassword(password string) {
	log.Println("Resetting admin password...")
	if password == "" {
		var err error
		password, err = utils.GenerateRandomChars(12, 5)
		if err != nil {
			password = "123456"
		}
	}
	err := CreateUser(&user_types.User{
		Username: "admin",
		Password: utils.GenPasswordHash(password),
		Email:    "admin@admin.com",
		Nickname: "管理员",
		IsAdmin:  true,
	})
	if err == nil {
		log.Println("Username: admin")
		log.Println("Password:", password)
	} else {
		fmt.Println(err)
	}
}

// 初始化数据库
func initDatabase() {
	log.Println("Initializing Database...")
	err := db.AutoMigrate(
		&user_types.User{},
		&system_types.DictType{},
		&system_types.DictData{},
		&system_types.SysConfig{},
		&system_types.File{},
		&article_types.ArticleGroup{},
		&article_types.Article{},
		&article_types.ArticleTag{},
		&article_types.ArticleComment{},
		&vuln_types.VulnType{},
		&vuln_types.Vuln{},
	)
	if err != nil {
		log.Fatal(err)
	}
	//初始化数据表
	{
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_status",
			Name:     "系统通用状态",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_user_status",
			Name:     "账号状态",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_file_storage",
			Name:     "文件存储类型",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_dict_type",
			Name:     "字典类型",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_login_type",
			Name:     "登录方式",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_gender_type",
			Name:     "性别",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_ai_type",
			Name:     "AI类型",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_ai_model",
			Name:     "模型名称",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_article_type",
			Name:     "文章类型",
			Type:     "system",
		})

		err = CreateDictType(&system_types.DictType{
			DictType: "sys_vuln_attribute",
			Name:     "漏洞属性",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_vuln_level",
			Name:     "漏洞等级",
			Type:     "system",
		})
		err = CreateDictType(&system_types.DictType{
			DictType: "sys_vuln_status",
			Name:     "漏洞审核状态",
			Type:     "system",
		})
	}
	{
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_ai_type",
			Value:     "openai",
			LabelName: "OpenAi",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_article_type",
			Value:     "1",
			LabelName: "原创",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_article_type",
			Value:     "2",
			LabelName: "转载",
			IsDefault: true,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_ai_model",
			Value:     "gpt-4o-mini",
			LabelName: "gpt-4o-mini",
			IsDefault: true,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_gender_type",
			Value:     "0",
			LabelName: "未知",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_gender_type",
			Value:     "1",
			LabelName: "男",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_gender_type",
			Value:     "2",
			LabelName: "女",
			IsDefault: false,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_login_type",
			Value:     "user",
			LabelName: "用户登录",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_login_type",
			Value:     "muna",
			LabelName: "厂商登录",
			IsDefault: true,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_user_status",
			Value:     "1",
			LabelName: "正常",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_user_status",
			Value:     "2",
			LabelName: "已停用",
			IsDefault: false,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_dict_type",
			Value:     "system",
			LabelName: "系统内置",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_dict_type",
			Value:     "user",
			LabelName: "用户字典",
			IsDefault: false,
			Type:      "system",
		})

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_status",
			Value:     "1",
			LabelName: "正常",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_file_storage",
			Value:     "local",
			LabelName: "本地存储",
			IsDefault: true,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_file_storage",
			Value:     "aliyun",
			LabelName: "阿里云OSS存储",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_file_storage",
			Value:     "tencent",
			LabelName: "腾讯云OSS存储",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_file_storage",
			Value:     "huawei",
			LabelName: "华为云OSS存储",
			IsDefault: false,
			Type:      "system",
		})

		//漏洞类型

		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_attribute",
			Value:     "eve",
			LabelName: "事件型",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_attribute",
			Value:     "uni",
			LabelName: "通用型",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_level",
			Value:     "risk",
			LabelName: "高危",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_level",
			Value:     "medium",
			LabelName: "中危",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_level",
			Value:     "low",
			LabelName: "低危",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_status",
			Value:     "accepted",
			LabelName: "审核通过",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_status",
			Value:     "rejected",
			LabelName: "审核驳回",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_status",
			Value:     "under_review",
			LabelName: "审核中",
			IsDefault: false,
			Type:      "system",
		})
		err = CreateDictData(&system_types.DictData{
			DictType:  "sys_vuln_status",
			Value:     "repaired",
			LabelName: "已经修复",
			IsDefault: false,
			Type:      "system",
		})
	}

	//其他数据
	{
		vulnerabilities := []*vuln_types.VulnType{
			//基本类型
			{BaseType: base_types.BaseType{ID: 1}, TypeName: "Web漏洞", ParentID: 0, Desc: "基本数据类型"},
			{BaseType: base_types.BaseType{ID: 2}, TypeName: "App漏洞", ParentID: 0, Desc: "基本数据类型"},
			{BaseType: base_types.BaseType{ID: 3}, TypeName: "Iot漏洞", ParentID: 0, Desc: "基本数据类型"},
			{BaseType: base_types.BaseType{ID: 4}, TypeName: "工控漏洞", ParentID: 0, Desc: "基本数据类型"},
			{BaseType: base_types.BaseType{ID: 5}, TypeName: "操作系统漏洞", ParentID: 0, Desc: "基本数据类型"},

			//Web漏洞
			{TypeName: "XSS", ParentID: 1, Desc: "XSS（跨站脚本攻击）允许攻击者在受害者用户的浏览器中执行恶意JavaScript代码。"},
			{TypeName: "配置错误", ParentID: 1, Desc: "安全配置错误包括未正确配置的安全设置，如默认密码、不必要的服务启用等。"},
			{TypeName: "弱口令", ParentID: 1, Desc: "使用容易被猜测或暴力破解的简单密码可能导致账户被攻破。"},
			{TypeName: "入侵事件", ParentID: 1, Desc: "指未经授权的访问行为，可能涉及数据窃取、系统破坏等。"},
			{TypeName: "疑似被黑", ParentID: 1, Desc: "存在迹象表明系统可能已经被攻击者控制或渗透。"},
			{TypeName: "文件上传", ParentID: 1, Desc: "允许用户上传任意文件可能导致恶意文件被执行或覆盖现有文件。"},
			{TypeName: "信息泄露", ParentID: 1, Desc: "敏感数据暴露，由于缺乏加密或不当的存储方式，敏感数据可能会被泄露。"},
			{TypeName: "存在后门", ParentID: 1, Desc: "系统中存在的隐蔽入口，允许攻击者绕过正常的认证机制进入系统。"},
			{TypeName: "逻辑漏洞", ParentID: 1, Desc: "应用程序中的逻辑错误可能导致非预期的行为，例如权限提升或数据篡改。"},
			{TypeName: "代码执行", ParentID: 1, Desc: "允许攻击者在服务器上执行任意代码，通常通过注入或反序列化漏洞实现。"},
			{TypeName: "命令执行", ParentID: 1, Desc: "命令注入漏洞允许攻击者通过应用程序执行任意操作系统命令。"},
			{TypeName: "SQL注入", ParentID: 1, Desc: "SQL注入是一种通过将恶意SQL语句插入到查询中，从而操纵数据库的行为。"},
			{TypeName: "解析漏洞", ParentID: 1, Desc: "解析漏洞通常出现在文件解析器中，允许攻击者通过构造特殊的输入触发异常行为。"},
			//App漏洞
			{TypeName: "代码执行", ParentID: 2, Desc: "允许攻击者在应用程序中执行任意代码，通常通过注入或反序列化漏洞实现。"},
			{TypeName: "逻辑漏洞", ParentID: 2, Desc: "应用程序中的逻辑错误可能导致非预期的行为，例如权限提升或数据篡改。"},
			{TypeName: "信息泄漏", ParentID: 2, Desc: "敏感数据暴露，由于缺乏加密或不当的存储方式，敏感数据可能会被泄露。"},
			{TypeName: "数据验证不完备类漏洞", ParentID: 2, Desc: "应用程序未能正确验证用户输入，可能导致各种安全问题。"},
			{TypeName: "内存管理类漏洞", ParentID: 2, Desc: "内存管理不当可能导致缓冲区溢出、释放后使用等问题，进而引发安全漏洞。"},
			{TypeName: "溢出漏洞", ParentID: 2, Desc: "缓冲区溢出漏洞允许攻击者覆盖内存中的数据，从而控制程序行为。"},
			{TypeName: "拒绝服务", ParentID: 2, Desc: "拒绝服务攻击旨在使应用程序不可用，导致合法用户无法访问服务。"},
			{TypeName: "文件读写", ParentID: 2, Desc: "允许攻击者读取或写入任意文件，可能导致敏感数据泄露或系统破坏。"},
			{TypeName: "命令注入", ParentID: 2, Desc: "命令注入漏洞允许攻击者通过应用程序执行任意操作系统命令。"},
			{TypeName: "其他", ParentID: 2, Desc: "未分类的其他类型漏洞。"},
			//IOT漏洞
			{TypeName: "内存破坏", ParentID: 3, Desc: "内存管理不当可能导致缓冲区溢出、释放后使用等问题，进而引发安全漏洞。"},
			{TypeName: "弱口令", ParentID: 3, Desc: "使用容易被猜测或暴力破解的简单密码，增加了设备被非法访问的风险。"},
			{TypeName: "目录遍历", ParentID: 3, Desc: "攻击者可能通过操纵文件路径来访问服务器上的任意文件，导致敏感信息泄露。"},
			{TypeName: "信息泄露", ParentID: 3, Desc: "由于缺乏加密或不当的存储方式，敏感数据可能会被泄露。"},
			{TypeName: "存在后门", ParentID: 3, Desc: "预留在系统中的特殊入口，允许未经授权的用户绕过正常的安全机制进行访问。"},
			{TypeName: "逻辑漏洞", ParentID: 3, Desc: "应用程序中的逻辑错误可能导致非预期的行为，例如权限提升或数据篡改。"},
			{TypeName: "配置错误", ParentID: 3, Desc: "不正确的配置可能导致暴露调试接口或其他安全隐患，使得设备易受攻击。"},
			{TypeName: "代码执行", ParentID: 3, Desc: "允许攻击者在物联网设备中执行任意代码，通常通过注入或反序列化漏洞实现。"},
			{TypeName: "命令执行", ParentID: 3, Desc: "命令注入漏洞允许攻击者通过应用程序执行任意操作系统命令。"},
			{TypeName: "权限提升", ParentID: 3, Desc: "利用漏洞获取比分配给用户的更多权限，从而可以访问或操作更多的资源。"},
			{TypeName: "逻辑漏洞", ParentID: 3, Desc: "重复项，请参考上述逻辑漏洞描述。"},
			{TypeName: "拒绝服务", ParentID: 3, Desc: "拒绝服务攻击旨在使物联网设备不可用，导致合法用户无法访问服务。"},
			{TypeName: "硬编码漏洞", ParentID: 3, Desc: "在代码中直接写入的不可变值（如用户名、密码），增加了安全风险。"},
			{TypeName: "溢出漏洞", ParentID: 3, Desc: "缓冲区溢出漏洞允许攻击者覆盖内存中的数据，从而控制程序行为。"},
			{TypeName: "其他", ParentID: 3, Desc: "未分类的其他类型漏洞。"},
			//工控漏洞
			{TypeName: "拒绝服务", ParentID: 4, Desc: "拒绝服务攻击旨在使设备不可用，导致合法用户无法访问服务。"},
			{TypeName: "任意文件操作", ParentID: 4, Desc: "允许攻击者读取或写入任意文件，可能导致敏感数据泄露或系统破坏。"},
			{TypeName: "命令执行", ParentID: 4, Desc: "命令注入漏洞允许攻击者通过应用程序执行任意操作系统命令。"},
			{TypeName: "内存破坏", ParentID: 4, Desc: "内存管理不当可能导致缓冲区溢出、释放后使用等问题，进而引发安全漏洞。"},
			{TypeName: "输入验证", ParentID: 4, Desc: "应用程序未能正确验证用户输入，可能导致各种安全问题。"},
			{TypeName: "代码执行", ParentID: 4, Desc: "允许攻击者在设备中执行任意代码，通常通过注入或反序列化漏洞实现。"},
			{TypeName: "目录遍历", ParentID: 4, Desc: "攻击者可能通过操纵文件路径来访问服务器上的任意文件，导致敏感信息泄露。"},
			{TypeName: "远程控制", ParentID: 4, Desc: "未经授权的远程访问或控制，使得攻击者可以远程操控设备。"},
			{TypeName: "权限提升", ParentID: 4, Desc: "利用漏洞获取比分配给用户的更多权限，从而可以访问或操作更多的资源。"},
			{TypeName: "未授权访问", ParentID: 4, Desc: "由于身份验证或授权机制的缺陷，导致未经授权的用户能够访问系统资源。"},
			{TypeName: "信息泄露", ParentID: 4, Desc: "由于缺乏加密或不当的存储方式，敏感数据可能会被泄露。"},
			{TypeName: "溢出漏洞", ParentID: 4, Desc: "缓冲区溢出漏洞允许攻击者覆盖内存中的数据，从而控制程序行为。"},
			{TypeName: "其他", ParentID: 4, Desc: "未分类的其他类型漏洞。"},
			//操作系统漏洞
			{TypeName: "远程执行", ParentID: 5, Desc: "允许攻击者通过网络在目标操作系统上执行任意代码，通常利用服务或应用程序中的漏洞。"},
			{TypeName: "拒绝服务", ParentID: 5, Desc: "拒绝服务攻击旨在使操作系统不可用，导致合法用户无法访问服务。"},
			{TypeName: "权限提升", ParentID: 5, Desc: "利用漏洞获取比分配给用户的更多权限，从而可以访问或操作更多的资源。"},
			{TypeName: "代码执行", ParentID: 5, Desc: "允许攻击者在操作系统中执行任意代码，通常通过注入或反序列化漏洞实现。"},
			{TypeName: "溢出漏洞", ParentID: 5, Desc: "缓冲区溢出漏洞允许攻击者覆盖内存中的数据，从而控制程序行为。"},
			{TypeName: "虚拟逃逸", ParentID: 5, Desc: "虚拟机逃逸是一种攻击技术，攻击者通过利用虚拟机中的漏洞来突破虚拟环境，进而影响宿主机或其他虚拟机。"},
			{TypeName: "其他", ParentID: 5, Desc: "未分类的其他类型漏洞。"},
		}

		for _, vuln := range vulnerabilities {
			if err = CreateVulnType(vuln); err != nil {
				fmt.Println(err)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
func init() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Config.Log.File,
		MaxSize:    10,   // megabytes
		MaxBackups: 7,    // number of old log files to retain
		MaxAge:     30,   // days
		Compress:   true, // whether to compress the log files
	}
	gormLogger := logger.New(
		log.New(lumberjackLogger, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true, // 忽略记录未找到错误
		},
	)
	switch config.Config.Log.Level {
	case "silent":
		gormLogger = gormLogger.LogMode(logger.Silent)
	case "error":
		gormLogger = gormLogger.LogMode(logger.Error)
	case "warn":
		gormLogger = gormLogger.LogMode(logger.Warn)
	case "info":
		gormLogger = gormLogger.LogMode(logger.Info)
	}
	InitSQL(gormLogger)
	InitRedis()
	//初始化数据库，检查数据库是否存在
	tables := []interface{}{
		&user_types.User{},
		&system_types.DictType{},
		&system_types.DictData{},
		&system_types.SysConfig{},
		&system_types.File{},
		&article_types.ArticleGroup{},
		&article_types.Article{},
		&article_types.ArticleTag{},
		&article_types.ArticleComment{},
		&vuln_types.VulnType{},
		&vuln_types.Vuln{},
	}
	fmt.Println(config.Config.Database)
	allTablesExist := checkTablesExist(db, tables)
	if !allTablesExist {
		initDatabase()
	}
	res := db.Where("username = 'admin'").First(&user.User{}).RowsAffected
	if res == 0 {
		initAdminPassword("123456")
	}
}

// generateDSN 根据配置生成 DSN
func generateDSN(config system_types.Config) string {
	switch config.Database.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.Database.Connection.User,
			config.Database.Connection.Password,
			config.Database.Connection.Host,
			config.Database.Connection.Port,
			config.Database.Connection.Name,
			config.Database.Connection.Charset,
		)
	case "postgres":
		return fmt.Sprintf("user_types=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Database.Connection.User,
			config.Database.Connection.Password,
			config.Database.Connection.Name,
			config.Database.Connection.Port,
		)
	case "sqlite":
		return config.Database.Connection.File // SQLite 使用文件路径
	case "sqlserver":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			config.Database.Connection.User,
			config.Database.Connection.Password,
			config.Database.Connection.Host,
			config.Database.Connection.Port,
			config.Database.Connection.Name,
		)
	default:
		return ""
	}
}
