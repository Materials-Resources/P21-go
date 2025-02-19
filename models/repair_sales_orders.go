package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type RepairSalesOrders struct {
	bun.BaseModel        `bun:"table:repair_sales_orders"`
	RepairSalesOrdersUid int32           `bun:"repair_sales_orders_uid,type:int,autoincrement,identity,pk"`
	RepairSalesOrdersRun int32           `bun:"repair_sales_orders_run,type:int,unique"`
	OrderNo              string          `bun:"order_no,type:varchar(255),unique"`
	DateRun              mssql.DateTime1 `bun:"date_run,type:datetime"`
}
