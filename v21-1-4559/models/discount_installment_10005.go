package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type DiscountInstallment10005 struct {
	bun.BaseModel           `bun:"table:discount_installment_10005"`
	DiscountInstallmentUid  int32           `bun:"discount_installment_uid,type:int,pk"`                          // Unique internal record identifier - not visible to the user
	DiscountInstallmentName string          `bun:"discount_installment_name,type:varchar(20)"`                    // Displayed discounting factor (First, Last, All, None)
	DateCreated             mssql.DateTime1 `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        mssql.DateTime1 `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DiscountInstallmentCode string          `bun:"discount_installment_code,type:char(3)"`                        // Internal code for discounting factor
}
