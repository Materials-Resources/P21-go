package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PcMessageDetail struct {
	bun.BaseModel `bun:"table:pc_message_detail"`
	MessageSkey   int32            `bun:"message_skey,type:int,pk"`        // Padlock table
	LanguageSkey  int32            `bun:"language_skey,type:int,pk"`       // Padlock table
	MessageTitle  *string          `bun:"message_title,type:varchar(255)"` // Padlock table
	MessageText   *string          `bun:"message_text,type:varchar(255)"`  // Padlock table
	CreateUserId  *string          `bun:"create_user_id,type:char(10)"`    // Padlock table
	CreateDate    *mssql.DateTime1 `bun:"create_date,type:datetime"`       // Padlock table
	MaintUserId   *string          `bun:"maint_user_id,type:char(10)"`     // Padlock table
	MaintDate     *mssql.DateTime1 `bun:"maint_date,type:datetime"`        // Padlock table
}
