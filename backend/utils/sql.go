package utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

func isDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}
	// 检查 MySQL 的特定错误码 1062（重复键）
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	return false
}
