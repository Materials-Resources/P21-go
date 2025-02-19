package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PlProfileGroup struct {
	bun.BaseModel `bun:"table:pl_profile_group"`
	GroupSkey     int32            `bun:"group_skey,type:int,pk"`       // Padlock table.
	ProfileSkey   int32            `bun:"profile_skey,type:int,pk"`     // Padlock table.
	CreateUserId  *string          `bun:"create_user_id,type:char(10)"` // Padlock table.
	CreateDate    *mssql.DateTime1 `bun:"create_date,type:datetime"`    // Padlock table.
	MaintUserId   *string          `bun:"maint_user_id,type:char(10)"`  // This column is unused.
	MaintDate     *mssql.DateTime1 `bun:"maint_date,type:datetime"`     // This column is unused.
}
