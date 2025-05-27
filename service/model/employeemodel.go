package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EmployeeModel = (*customEmployeeModel)(nil)

type (
	// EmployeeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEmployeeModel.
	EmployeeModel interface {
		employeeModel
	}

	customEmployeeModel struct {
		*defaultEmployeeModel
	}
)

// NewEmployeeModel returns a model for the database table.
func NewEmployeeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) EmployeeModel {
	return &customEmployeeModel{
		defaultEmployeeModel: newEmployeeModel(conn, c, opts...),
	}
}
