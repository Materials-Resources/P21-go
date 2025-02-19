package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type ProcessSystemParameters struct {
	bun.BaseModel              `bun:"table:process_system_parameters"`
	ProcessSystemParametersUid int16           `bun:"process_system_parameters_uid,type:tinyint,pk"`
	DayDefinition              int16           `bun:"day_definition,type:tinyint"`
	DateCreated                mssql.DateTime1 `bun:"date_created,type:datetime"`
	DateLastModified           mssql.DateTime1 `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
