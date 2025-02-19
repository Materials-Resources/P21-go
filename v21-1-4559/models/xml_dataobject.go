package models

import (
	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
)

type XmlDataobject struct {
	bun.BaseModel      `bun:"table:xml_dataobject"`
	XmlDataobjectUid   int32           `bun:"xml_dataobject_uid,type:int,pk"`        // Unique identifier for records within xml_dataobject table
	DataobjectName     string          `bun:"dataobject_name,type:varchar(40)"`      // Name of container of retrieved data
	BusinessObjectName string          `bun:"business_object_name,type:varchar(32)"` // Name of business object associated with the dataobject
	SetupEventName     string          `bun:"setup_event_name,type:varchar(32)"`     // Name of event to be triggered to set up import or export
	DateCreated        mssql.DateTime1 `bun:"date_created,type:datetime"`            // Indicates the date/time this record was created.
	DateLastModified   mssql.DateTime1 `bun:"date_last_modified,type:datetime"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy   string          `bun:"last_maintained_by,type:varchar(30)"`   // ID of the user who last maintained this record
}
