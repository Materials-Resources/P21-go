package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PlGroupDef struct {
	bun.BaseModel `bun:"table:pl_group_def"`
	GroupSkey     int32            `bun:"group_skey,type:int,pk"`              // Padlock table.
	Priority      *int32           `bun:"priority,type:int"`                   // Padlock table.
	SecAdminInd   int32            `bun:"sec_admin_ind,type:int"`              // Padlock table.
	GroupName     string           `bun:"group_name,type:varchar(100),unique"` // Padlock table.
	CreateUserId  *string          `bun:"create_user_id,type:char(10)"`        // Padlock table.
	CreateDate    *mssql.DateTime1 `bun:"create_date,type:datetime"`           // Padlock table.
	MaintUserId   *string          `bun:"maint_user_id,type:char(10)"`         // Padlock table.
	MaintDate     *mssql.DateTime1 `bun:"maint_date,type:datetime"`            // This column is unused.
}
