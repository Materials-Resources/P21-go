package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type OePickTicketUd struct {
	bun.BaseModel     `bun:"table:oe_pick_ticket_ud"`
	OePickTicketUdUid int32           `bun:"oe_pick_ticket_ud_uid,type:int,autoincrement,identity,pk"`
	PickTicketNo      float64         `bun:"pick_ticket_no,type:decimal(19,0),unique"`
	Picker            *string         `bun:"picker,type:varchar(40)"` // Picker
	DateCreated       mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string          `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
