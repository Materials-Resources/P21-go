package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type DbDrivenMaintKey struct {
	bun.BaseModel       `bun:"table:db_driven_maint_key"`
	DbDrivenMaintKeyUid int32           `bun:"db_driven_maint_key_uid,type:int,autoincrement,identity,pk"`   // UID for this table.
	DbDrivenMaintUid    int32           `bun:"db_driven_maint_uid,type:int"`                                 // db_driven_maint record that this key is for.
	KeyColumnName       string          `bun:"key_column_name,type:varchar(255)"`                            // Key column name.
	DateCreated         mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
