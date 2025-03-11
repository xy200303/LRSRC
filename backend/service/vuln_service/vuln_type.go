package vuln_service

import (
	"fmt"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/vuln_types"
)

// ListVulnType 查询漏洞类型
func ListVulnType(req *vuln_types.VulnTypeQuery) ([]*vuln_types.VulnType, error) {
	result, err := models.ListVulnType(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateVulnType 更新漏洞类型
func UpdateVulnType(req *vuln_types.VulnType) error {
	err := models.UpdateVulnType(req)
	if err != nil {
		return err
	}
	return nil
}

// CreateVulnType 创建漏洞类型
func CreateVulnType(req *vuln_types.VulnType) error {
	err := models.CreateVulnType(req)
	if err != nil {
		return err
	}
	return nil
}

// DeleteVulnType 删除漏洞类型
func DeleteVulnType(req *base_types.DeleteIntIds) error {
	err := models.DeleteVulnType(req.Ids)
	if err != nil {
		return err
	}
	return nil
}

// BuildVulnTypeTree 构建选择树数据
func BuildVulnTypeTree() ([]*base_types.TreeData, error) {
	// 获取所有漏洞类型
	vulnTypes, err := models.ListVulnType(&vuln_types.VulnTypeQuery{})
	if err != nil {
		return nil, fmt.Errorf("failed to list vulnerability types: %w", err)
	}

	nodes := make(map[uint64]*base_types.TreeData)
	var rootNodes []*base_types.TreeData

	// 初始化所有节点
	for _, vt := range vulnTypes {
		nodes[vt.ID] = &base_types.TreeData{
			Label:    vt.TypeName,
			Value:    vt.ID,
			Children: []*base_types.TreeData{},
		}
	}

	// 构建树结构
	for _, vt := range vulnTypes {
		node := nodes[vt.ID]
		if parentNode := nodes[vt.ParentID]; vt.ParentID != 0 && parentNode != nil {
			parentNode.Children = append(parentNode.Children, node)
		} else {
			// 没有父节点或父节点不存在，则为根节点
			rootNodes = append(rootNodes, node)
		}
	}

	return rootNodes, nil
}
