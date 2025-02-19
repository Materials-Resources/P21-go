package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PcProfileDef struct {
	bun.BaseModel `bun:"table:pc_profile_def"`
	ProfileSkey   int32            `bun:"profile_skey,type:int,pk"`              // Padlock table
	SysAdminInd   int32            `bun:"sys_admin_ind,type:int,default:(0)"`    // Padlock table
	ProfileName   string           `bun:"profile_name,type:varchar(100),unique"` // Padlock table
	CreateUserId  *string          `bun:"create_user_id,type:char(10)"`          // Padlock table
	CreateDate    *mssql.DateTime1 `bun:"create_date,type:datetime"`             // Padlock table
	MaintUserId   *string          `bun:"maint_user_id,type:char(10)"`           // Padlock table
	MaintDate     *mssql.DateTime1 `bun:"maint_date,type:datetime"`              // Padlock table
}
