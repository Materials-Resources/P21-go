package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type ReportCodeXCodeGroupP21 struct {
	bun.BaseModel              `bun:"table:report_code_x_code_group_p21"`
	ReportCodeXCodeGroupP21Uid int32           `bun:"report_code_x_code_group_p21_uid,type:int,autoincrement,identity,pk"` // Primary key
	ReportCodeGroupP21Uid      int32           `bun:"report_code_group_p21_uid,type:int"`                                  // Relationship to report_code_group_p21
	ReportCodeP21Uid           int32           `bun:"report_code_p21_uid,type:int"`                                        // Relationship to report_code_p21
	DateCreated                mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`                      // Date and time the record was originally created
	CreatedBy                  string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`                // User who created the record
	DateLastModified           mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`                // Date and time the record was modified
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`        // User who last changed the record
}
