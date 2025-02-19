package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type WeboeFile struct {
	bun.BaseModel     `bun:"table:weboe_file"`
	WeboeFileUid      int32           `bun:"weboe_file_uid,type:int,autoincrement,identity"`
	WeboeFileName     string          `bun:"weboe_file_name,type:varchar(255)"`
	WeboeOldFileName  string          `bun:"weboe_old_file_name,type:varchar(255)"`
	WeboeFileDesc     string          `bun:"weboe_file_desc,type:varchar(255)"`
	WeboeLastTransfer mssql.DateTime1 `bun:"weboe_last_transfer,type:datetime"`
	PartialTransfer   *bool           `bun:"partial_transfer,type:bit"`
	TransferOrder     *int16          `bun:"transfer_order,type:smallint"`
	FileSql           string          `bun:"file_sql,type:varchar(8000)"`
}
