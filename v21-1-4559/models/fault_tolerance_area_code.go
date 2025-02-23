package models

import "github.com/uptrace/bun"

type FaultToleranceAreaCode struct {
	bun.BaseModel `bun:"table:fault_tolerance_area_code"`
	FtaUid        int32   `bun:"fta_uid,type:int,autoincrement,identity,pk"`
	FtaCode       string  `bun:"fta_code,type:varchar(255)"`
	FtaTableName  *string `bun:"fta_table_name,type:varchar(255)"`
}
