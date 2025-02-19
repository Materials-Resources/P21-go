package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type ItemIdChangeHistory struct {
	bun.BaseModel          `bun:"table:item_id_change_history"`
	InvMastUid             int32           `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	OldItemId              string          `bun:"old_item_id,type:varchar(40)"`                              // The old item id.
	NewItemId              string          `bun:"new_item_id,type:varchar(40)"`                              // The new item id.
	DateCreated            mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified       mssql.DateTime1 `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy       string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ItemIdChangeHistoryUid int32           `bun:"item_id_change_history_uid,type:int,autoincrement,identity,pk"`
}
