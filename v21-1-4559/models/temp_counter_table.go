package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type TempCounterTable struct {
	bun.BaseModel  `bun:"table:temp_counter_table"`
	Uid            int32           `bun:"uid,type:int,autoincrement,identity"`
	JobPriceHdrUid int32           `bun:"job_price_hdr_uid,type:int"`
	StartDate      mssql.DateTime1 `bun:"start_date,type:datetime"`
	EndDate        mssql.DateTime1 `bun:"end_date,type:datetime"`
}
