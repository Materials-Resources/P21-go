package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type P21Dblevel struct {
	bun.BaseModel `bun:"table:p21_dblevel"`
	Version       *int16          `bun:"version,type:tinyint"`
	DateCreated   mssql.DateTime1 `bun:"date_created,type:datetime,default:(getdate())"`
}
