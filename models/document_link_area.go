package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type DocumentLinkArea struct {
	bun.BaseModel       `bun:"table:document_link_area"`
	DocumentLinkAreaUid int32           `bun:"document_link_area_uid,type:int,pk"`
	DocumentLinkUid     int32           `bun:"document_link_uid,type:int"`
	DisplayAreaCd       int32           `bun:"display_area_cd,type:int"`
	RowStatusFlag       int32           `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated         mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
