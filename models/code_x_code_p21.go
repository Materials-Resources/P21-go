package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type CodeXCodeP21 struct {
	bun.BaseModel    `bun:"table:code_x_code_p21"`
	CodeXCodeUid     int32           `bun:"code_x_code_uid,type:int,pk"`
	SourceCodeNo     int32           `bun:"source_code_no,type:int"`
	TargetCodeNo     int32           `bun:"target_code_no,type:int"`
	SequenceNumber   int32           `bun:"sequence_number,type:int,default:(1)"`
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
