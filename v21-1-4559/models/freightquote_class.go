package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type FreightquoteClass struct {
	bun.BaseModel        `bun:"table:freightquote_class"`
	FreightquoteClassUid int32           `bun:"freightquote_class_uid,type:int,autoincrement,identity,pk"`    // A unique identifier for each freightquote class
	FreightquoteClassId  float64         `bun:"freightquote_class_id,type:decimal(19,9)"`                     // The actual class identifier used by freightquote.com and users
	RowStatusFlag        int32           `bun:"row_status_flag,type:int"`                                     // Has the record been logically deleted?
	DateCreated          mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
