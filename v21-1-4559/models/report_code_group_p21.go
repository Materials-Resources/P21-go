package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type ReportCodeGroupP21 struct {
	bun.BaseModel         `bun:"table:report_code_group_p21"`
	ReportCodeGroupP21Uid int32           `bun:"report_code_group_p21_uid,type:int,autoincrement,identity,pk"` // Primary key
	CodeGroupNo           int32           `bun:"code_group_no,type:int"`                                       // Code of the group
	GroupTypeNo           *int32          `bun:"group_type_no,type:int"`                                       // code of type of group
	CodeGroupDescription  string          `bun:"code_group_description,type:varchar(255)"`                     // Decription used to display to the user
	DateCreated           mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
