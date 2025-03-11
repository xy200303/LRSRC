package system_controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// GetDictType 获取字典类型
func GetDictType(c *gin.Context) {
	dictType := c.Query("dict_type")
	dict, err := system_service.GetDictType(dictType)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, dict)
}

// GetDictData 获取字典数据
func GetDictData(c *gin.Context) {
	dictType := c.Query("dict_type")
	dict, err := system_service.GetDictData(dictType)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, dict)
}

// ListDictType 获取所有的dictType
func ListDictType(c *gin.Context) {
	var page base_types.BasePage
	//确保json请求格式
	if err := c.ShouldBind(&page); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&page, err), nil)
		log.Print(err)
		return
	}
	dictType, total, err := system_service.ListDictType(&page)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, dictType, total, &page)
}

// ListDictData 获取所有的dictType
func ListDictData(c *gin.Context) {
	dictType := c.Query("dict_type")
	var page base_types.BasePage
	//确保json请求格式
	if err := c.ShouldBind(&page); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&page, err), nil)
		log.Print(err)
		return
	}
	data, total, err := system_service.ListDictData(dictType, &page)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, data, total, &page)
}

// DeleteDictType 删除数据字典类型
func DeleteDictType(c *gin.Context) {
	var ids base_types.DeleteStringIds
	//确保json请求格式
	if err := c.ShouldBind(&ids); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&ids, err), nil)
		log.Print(err)
		return
	}
	err := system_service.DeleteDictType(ids.Ids)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// 添加数据
func CreateDictType(c *gin.Context) {
	var data system_types.DictTypeReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err := system_service.CreateDictType(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateDictType 更新数据
func UpdateDictType(c *gin.Context) {
	var data system_types.DictTypeReq
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err := system_service.UpdateDictType(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// DeleteDictData 删除数据字典类型
func DeleteDictData(c *gin.Context) {
	var ids base_types.DeleteIntIds
	//确保json请求格式
	if err := c.ShouldBind(&ids); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&ids, err), nil)
		log.Print(err)
		return
	}
	err := system_service.DeleteDictData(ids.Ids)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// CreateDictData 添加数据
func CreateDictData(c *gin.Context) {
	var data system_types.DictDataReq
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err := system_service.CreateDictData(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateDictData 更新数据
func UpdateDictData(c *gin.Context) {
	var data system_types.DictDataReq
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err := system_service.UpdateDictData(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}
