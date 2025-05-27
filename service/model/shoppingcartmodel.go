package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShoppingCartModel = (*customShoppingCartModel)(nil)

type (
	// ShoppingCartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShoppingCartModel.
	ShoppingCartModel interface {
		shoppingCartModel
	}

	customShoppingCartModel struct {
		*defaultShoppingCartModel
	}
)

// NewShoppingCartModel returns a model for the database table.
func NewShoppingCartModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShoppingCartModel {
	return &customShoppingCartModel{
		defaultShoppingCartModel: newShoppingCartModel(conn, c, opts...),
	}
}
