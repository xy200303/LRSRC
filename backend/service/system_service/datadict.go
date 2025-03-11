package system_service

import (
	"fmt"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
)

func ListDictType(page *base_types.BasePage) ([]*system_types.DictType, int64, error) {
	dict, total, err := models.ListDictType(page)
	if err != nil {
		return nil, total, err
	}
	return dict, total, nil
}

func GetDictData(dictType string) ([]*system_types.DictData, error) {
	dict, err := models.GetDictData(dictType)
	if err != nil {
		return nil, err
	}
	return dict, nil
}

func DeleteDictType(dictTypes []string) error {
	res, err := models.ListDictTypeByIds(dictTypes)
	if err != nil {
		return err
	}
	for _, r := range res {
		if r.Type == "system" {
			return fmt.Errorf("%s为系统内置字典类型，不可删除", r.DictType)
		}
	}
	err = models.DeleteDictType(dictTypes)
	return err
}

func DeleteDictData(dictTypes []int64) error {
	res, err := models.ListDictDataByIds(dictTypes)
	if err != nil {
		return err
	}
	for _, r := range res {
		if r.Type == "system" {
			return fmt.Errorf("字典编号%d为系统内置字典数据，不可删除", r.ID)
		}
	}
	err = models.DeleteDictData(dictTypes)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDictType(req *system_types.DictTypeReq) error {
	err := models.UpdateDictType(&system_types.DictType{
		Name:     req.Name,
		DictType: req.DictType,
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateDictData(req *system_types.DictDataReq) error {
	err := models.UpdateDictData(&system_types.DictData{
		BaseType: base_types.BaseType{
			ID: req.ID,
		},
		DictType:    req.DictType,
		Value:       req.Value,
		LabelName:   req.LabelName,
		IsDefault:   req.IsDefault,
		Type:        req.Type,
		ElTagEffect: req.ElTagEffect,
		ElTagType:   req.ElTagType,
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateDictType(req *system_types.DictTypeReq) error {
	err := models.CreateDictType(&system_types.DictType{
		Name:     req.Name,
		DictType: req.DictType,
	})
	if err != nil {
		return err
	}
	return nil
}
func CreateDictData(req *system_types.DictDataReq) error {
	err := models.CreateDictData(&system_types.DictData{
		DictType:  req.DictType,
		Value:     req.Value,
		LabelName: req.LabelName,
		IsDefault: req.IsDefault,
		Type:      req.Type,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetDictType(dictType string) (*system_types.DictType, error) {
	dict, err := models.GetDictType(dictType)
	if err != nil {
		return nil, err
	}
	return dict, nil
}

func ListDictData(dictType string, page *base_types.BasePage) ([]*system_types.DictData, int64, error) {
	data, total, err := models.ListDictData(dictType, page)
	if err != nil {
		return nil, total, err
	}
	return data, total, nil
}
