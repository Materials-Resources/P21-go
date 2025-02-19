package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type PcLanguageDef struct {
	bun.BaseModel   `bun:"table:pc_language_def"`
	LanguageSkey    int32            `bun:"language_skey,type:int,pk"`           // Padlock table
	LanguageId      string           `bun:"language_id,type:varchar(20),unique"` // Padlock table
	LanguageComment *string          `bun:"language_comment,type:varchar(100)"`  // Padlock table
	CreateUserId    *string          `bun:"create_user_id,type:char(10)"`        // Padlock table
	CreateDate      *mssql.DateTime1 `bun:"create_date,type:datetime"`           // Padlock table
	MaintUserId     *string          `bun:"maint_user_id,type:char(10)"`         // Padlock table
	MaintDate       *mssql.DateTime1 `bun:"maint_date,type:datetime"`            // Padlock table
}
