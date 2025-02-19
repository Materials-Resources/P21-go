package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type CodeGroupP21 struct {
	bun.BaseModel        `bun:"table:code_group_p21"`
	CodeGroupUid         int32           `bun:"code_group_uid,type:int,pk"`
	CodeGroupNo          int16           `bun:"code_group_no,type:smallint"`
	GroupTypeNo          int16           `bun:"group_type_no,type:tinyint"`
	CodeGroupDescription string          `bun:"code_group_description,type:varchar(255)"`
	RowStatusFlag        string          `bun:"row_status_flag,type:char(1)"`
	DateCreated          mssql.DateTime1 `bun:"date_created,type:datetime"`
	DateLastModified     mssql.DateTime1 `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string          `bun:"last_maintained_by,type:varchar(30)"`
}
