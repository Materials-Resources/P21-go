package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type DynachangeConfig struct {
	bun.BaseModel    `bun:"table:dynachange_config"`
	ConfigurationId  int32           `bun:"configuration_id,type:int,pk"`        // A company id that the dynachange is being applied too
	DynachangeId     float64         `bun:"dynachange_id,type:decimal(6,0),pk"`  // What is the unique identifier for this dynachange?
	DateCreated      mssql.DateTime1 `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified mssql.DateTime1 `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
