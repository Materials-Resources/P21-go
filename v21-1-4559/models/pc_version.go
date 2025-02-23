package models

import "github.com/uptrace/bun"

type PcVersion struct {
	bun.BaseModel  `bun:"table:pc_version"`
	ProductCode    string  `bun:"product_code,type:varchar(10),pk"`  // Padlock table
	ProductVersion *string `bun:"product_version,type:varchar(100)"` // Padlock table
	ScriptName     *string `bun:"script_name,type:varchar(20)"`      // Padlock table
	DbVersion      *string `bun:"db_version,type:varchar(20)"`       // Padlock table
}
