package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type RibbonTabGroupXRibbonTab struct {
	bun.BaseModel               `bun:"table:ribbon_tab_group_x_ribbon_tab"`
	RibbonTabGroupXRibbonTabUid int32           `bun:"ribbon_tab_group_x_ribbon_tab_uid,type:int,autoincrement,identity,pk"`
	RibbonTabGroupUid           int32           `bun:"ribbon_tab_group_uid,type:int"`
	RibbonTabUid                int32           `bun:"ribbon_tab_uid,type:int"`
	SequenceNo                  int32           `bun:"sequence_no,type:int"`
	DateCreated                 mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
