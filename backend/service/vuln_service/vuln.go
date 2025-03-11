package vuln_service

import (
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/vuln_types"
)

// ListVuln 查询漏洞数据
func ListVuln(req *vuln_types.VulnQuery) ([]*vuln_types.VulnResp, int64, error) {
	vuln, i, err := models.ListVuln(req)
	if err != nil {
		return nil, 0, err
	}
	return vuln, i, nil
}

// CreateVuln 添加漏洞数据
func CreateVuln(req *vuln_types.Vuln) error {
	err := models.CreateVuln(req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateVuln 更新数据
func UpdateVuln(req *vuln_types.Vuln) error {
	err := models.UpdateVuln(req)
	if err != nil {
		return err
	}
	return nil
}

// DeleteVuln 删除数据
func DeleteVuln(req *base_types.DeleteStringIds) error {
	err := models.DeleteVuln(req.Ids)
	if err != nil {
		return err
	}
	return nil
}

// 获取漏洞信息
func GetVuln(id string) (*vuln_types.VulnResp, error) {
	vuln, err := models.GetVuln(id)
	if err != nil {
		return nil, err
	}
	return vuln, nil
}

// SubmitVuln 添加漏洞数据
func SubmitVuln(req *vuln_types.Vuln) error {
	err := models.CreateVuln(req)
	if err != nil {
		return err
	}
	return nil
}

// AuditVuln 审核漏洞数据
func AuditVuln(req *vuln_types.VulnAuditReq) error {
	err := models.AuditVuln(req)
	if err != nil {
		return err
	}
	return nil
}
