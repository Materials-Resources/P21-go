package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type ExtensibilityObject struct {
	bun.BaseModel          `bun:"table:extensibility_object"`
	ExtensibilityObjectUid int32           `bun:"extensibility_object_uid,type:int,autoincrement,identity,pk"`
	ExtensibilityWindowUid int32           `bun:"extensibility_window_uid,type:int,unique"`
	ObjectType             string          `bun:"object_type,type:varchar(255),unique"`
	ObjectBaseClass        string          `bun:"object_base_class,type:varchar(255),unique"`
	EnableFlag             string          `bun:"enable_flag,type:char(1),default:('Y')"`
	RowStatusFlag          int32           `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated            mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
