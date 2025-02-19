package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type BusinessObjectKeyFields struct {
	bun.BaseModel              `bun:"table:business_object_key_fields"`
	BusinessObjectKeyFieldsUid int32           `bun:"business_object_key_fields_uid,type:int,autoincrement,identity,pk"` // UID for these records.
	BusinessObjectName         string          `bun:"business_object_name,type:varchar(100)"`                            // BO name (
	KeyFieldNames              string          `bun:"key_field_names,type:varchar(255)"`                                 // Key fields for this BO.
	DateCreated                mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                  string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified           mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
