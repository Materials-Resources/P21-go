package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type InstallmentFrequency10005 struct {
	bun.BaseModel            `bun:"table:installment_frequency_10005"`
	InstallmentFrequencyUid  int32           `bun:"installment_frequency_uid,type:int,pk"`       // Unique internal record identifier - not visible to the user
	InstallmentFrequencyName string          `bun:"installment_frequency_name,type:varchar(15)"` // Displayed frequency factor (i.e. Monthly, Weekly)
	DateCreated              mssql.DateTime1 `bun:"date_created,type:datetime"`                  // Indicates the date/time this record was created.
	DateLastModified         mssql.DateTime1 `bun:"date_last_modified,type:datetime"`            // Indicates the date/time this record was last modified.
	LastMaintainedBy         string          `bun:"last_maintained_by,type:varchar(30)"`         // ID of the user who last maintained this record
	InstallmentFrequencyCode string          `bun:"installment_frequency_code,type:char(3)"`     // Internal code for frequency factor
}
