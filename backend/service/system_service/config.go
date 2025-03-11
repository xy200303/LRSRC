package system_service

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"xiaoyun/backend/config"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils"
	"xiaoyun/backend/utils/llm"
	"xiaoyun/backend/utils/oss"
)

var (
	// SysConfigMap 系统可设置配置
	SysConfigMap = &system_types.SysConfigMap{}
)

func init() {
	//初始化系统可设置参数
	//初始化数据
	var err error
	err = CreateSysConfigMap(&config.Config.SysConfigMap)
	if err != nil {
		log.Fatalf("Error creating sysconfig: %v", err)
	}
	SysConfigMap, err = GetSysConfigMap()
	if err != nil {
		log.Fatalf("Error getting sysconfig: %v", err)
	}
	//初始化OSS对象
	oss.InitOss(SysConfigMap)
	//初始化AI对象
	llm.InitOpenAi(SysConfigMap)
}

func GetSysConfig(sysKey string) (*system_types.SysConfig, error) {
	dict, err := models.GetSysConfig(sysKey)
	if err != nil {
		return nil, err
	}
	return dict, nil
}
func UpdateSysConfig(req *system_types.SysConfig) error {
	err := models.UpdateSysConfig(req)
	if err != nil {
		return err
	}
	return nil
}
func GetSysStatus() map[string]interface{} {
	cpuUsage, memUsage, diskUsage := utils.GetSystemUsage()
	return map[string]interface{}{
		"cpu":  cpuUsage,
		"mem":  memUsage,
		"disk": diskUsage,
	}
}

func GetSysConfigString(sysKey string) (string, error) {
	config, err := GetSysConfig(sysKey)
	if err != nil {
		return "", err
	}
	return config.Value, nil
}

func GetSysConfigByGroup(group string) ([]*system_types.SysConfig, error) {
	configs, err := models.GetSysConfigByGroup(group)
	if err != nil {
		return nil, err
	}
	return configs, nil
}
func GetSysConfigInt(sysKey string) (int64, error) {
	config, err := GetSysConfig(sysKey)
	if err != nil {
		return -1, err
	}
	intValue, err := strconv.ParseInt(config.Value, 10, 64)
	if err != nil {
		return -1, err
	}
	return intValue, nil
}

// findFieldByJSONTag 查找结构体中具有指定JSON标签的字段
func findFieldByJSONTag(t reflect.Type, jsonTag string) (reflect.StructField, bool) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagValue := field.Tag.Get("json"); tagValue == jsonTag {
			return field, true
		}
	}
	return reflect.StructField{}, false
}

// SysConfig2Map Map2Struct Map到结构体
func SysConfig2Map(arr []*system_types.SysConfig) (*system_types.SysConfigMap, error) {
	// 获取结构体的反射值
	v := &system_types.SysConfigMap{}
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to a struct")
	}
	val = val.Elem()
	// 遍历传入的map
	for _, sys := range arr {
		key := sys.SysKey
		value := sys.Value
		// 根据JSON标签查找对应的结构体字段
		field, ok := findFieldByJSONTag(val.Type(), key)
		if !ok {
			log.Println(fmt.Sprintf("field %s not found", key))
			continue
		}
		// 设置字段值
		fieldValue := val.FieldByName(field.Name)
		if !fieldValue.CanSet() {
			continue // 如果字段不可设置，则跳过
		}
		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(value)
		case reflect.Int:
			intValue, err := strconv.Atoi(value)
			if err == nil {
				fieldValue.SetInt(int64(intValue))
			} else {
				return nil, fmt.Errorf("failed to parse integer value for key %s: %v", key, err)
			}
		case reflect.Uint8:
			intValue, err := strconv.Atoi(value)
			if err == nil {
				fieldValue.SetUint(uint64(intValue))
			} else {
				return nil, fmt.Errorf("failed to parse integer value for key %s: %v", key, err)
			}
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err == nil {
				fieldValue.SetBool(boolValue)
			}
		default:
			return nil, fmt.Errorf("unsupported field type for key %s", key)
		}
	}
	return v, nil
}

func Map2SysConfig(data *system_types.SysConfigMap) ([]*system_types.SysConfig, error) {
	if data == nil {
		return nil, errors.New("input is nil")
	}
	var configs []*system_types.SysConfig
	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		jsonTag := field.Tag.Get("json")
		nameTag := field.Tag.Get("name")
		if jsonTag == "" {
			log.Println(fmt.Sprintf("field %s not found", field.Name))
			continue // Skip fields without json tag
		}
		var valueStr string
		switch value.Kind() {
		case reflect.Bool:
			valueStr = strconv.FormatBool(value.Bool())
		case reflect.String:
			valueStr = value.String()
		case reflect.Int:
			valueStr = fmt.Sprintf("%d", value.Int())
		case reflect.Uint8:
			valueStr = fmt.Sprintf("%d", value.Uint())
		case reflect.Int16:
			valueStr = fmt.Sprintf("%d", value.Int())
		default:
			log.Println(fmt.Sprintf("unsupported field type for key %s", field.Name))
			continue // Skip unsupported user_types
		}
		configs = append(configs, &system_types.SysConfig{
			SysKey: jsonTag,
			Value:  valueStr,
			Name:   nameTag,
		})
	}
	return configs, nil
}

// GetSysConfigMap 获取参数的Map形式
func GetSysConfigMap() (*system_types.SysConfigMap, error) {
	group, err := models.ListSysConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data, err := SysConfig2Map(group)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}

// UpdateSysConfigMap 初始化数据
func UpdateSysConfigMap(req *system_types.SysConfigMap) error {
	config, err := Map2SysConfig(req)
	if err != nil {
		return err
	}
	err = models.UpdateSysConfigBatch(config)
	if err != nil {
		return err
	}
	//刷新配置
	SysConfigMap, err = GetSysConfigMap()
	if err != nil {
		return err
	}
	//刷新数据
	//初始化OSS对象
	oss.InitOss(SysConfigMap)
	llm.InitOpenAi(SysConfigMap)
	return nil
}

// CreateSysConfigMap 创建配置文件
func CreateSysConfigMap(req *system_types.SysConfigMap) error {
	config, err := Map2SysConfig(req)
	if err != nil {
		return err
	}
	err = models.CreateSysConfigBatch(config)
	if err != nil {
		return err
	}
	return nil
}
