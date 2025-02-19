package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type FreightquotePkgType struct {
	bun.BaseModel           `bun:"table:freightquote_pkg_type"`
	FreightquotePkgTypeUid  int32           `bun:"freightquote_pkg_type_uid,type:int,autoincrement,identity,pk"` // The unique identifier for this table
	FreightquotePkgTypeId   string          `bun:"freightquote_pkg_type_id,type:varchar(40)"`                    // the freightquote.com package type identifier
	FreightquotePkgTypeDesc *string         `bun:"freightquote_pkg_type_desc,type:varchar(255)"`                 // the description for the package type id
	RowStatusFlag           int32           `bun:"row_status_flag,type:int"`                                     // Has the record been logically deleted?
	DateCreated             mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
