package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type Inventoryissuestestrundetail struct {
	bun.BaseModel `bun:"table:InventoryIssuesTestRunDetail"`
	Run           int32           `bun:"run,type:int"`
	Testuid       int32           `bun:"TestUID,type:int"`
	Instances     int32           `bun:"instances,type:int"`
	DateStarted   mssql.DateTime1 `bun:"date_started,type:datetime"`
	DateCompleted mssql.DateTime1 `bun:"date_completed,type:datetime"`
	DateDiffMs    int32           `bun:"date_diff_ms,type:int"`
	TotalTestUid  *int32          `bun:"total_test_uid,type:int"`
}
