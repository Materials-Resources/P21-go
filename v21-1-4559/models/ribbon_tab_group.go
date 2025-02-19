package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type RibbonTabGroup struct {
	bun.BaseModel     `bun:"table:ribbon_tab_group"`
	RibbonTabGroupUid int32           `bun:"ribbon_tab_group_uid,type:int,autoincrement,identity,pk"`
	TabGroupId        string          `bun:"tab_group_id,type:varchar(255)"`
	TabGroupText      string          `bun:"tab_group_text,type:varchar(255)"`
	Description       string          `bun:"description,type:varchar(255)"`
	DateCreated       mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
