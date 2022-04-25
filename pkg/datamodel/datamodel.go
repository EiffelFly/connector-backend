package datamodel

import (
	"database/sql/driver"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	connectorPB "github.com/instill-ai/protogen-go/connector/v1alpha"
)

// BaseStatic contains common columns for all tables with static UUID as primary key
type BaseStatic struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BaseDynamic contains common columns for all tables with dynamic UUID as primary key generated when creating
type BaseDynamic struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseDynamic) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	db.Statement.SetColumn("ID", uuid)
	return nil
}

// ConnectorType is an alias type for Protobuf enum ConnectorType
type ConnectorType connectorPB.ConnectorType

// ConnectionType is an alias type for Protobuf enum ConnectionType
type ConnectionType connectorPB.ConnectionType

// ReleaseStage is an alias type for Protobuf enum ReleaseStage
type ReleaseStage connectorPB.ReleaseStage

// ConnectorDefinition is the data model of the connector_definition table
type ConnectorDefinition struct {
	BaseStatic
	Name                 string
	DockerRepository     string
	DockerImageTag       string
	DocumentationURL     string
	Icon                 string
	Tombstone            bool
	Public               bool
	Custom               bool
	ReleaseDate          *time.Time
	Spec                 datatypes.JSON `gorm:"type:jsonb"`
	ResourceRequirements datatypes.JSON `gorm:"type:jsonb"`
	ConnectorType        ConnectorType  `sql:"type:valid_connector_type"`
	ConnectionType       ConnectionType `sql:"type:valid_connection_type"`
	ReleaseStage         ReleaseStage   `sql:"type:valid_release_stage"`
}

// Connector is the data model of the connector table
type Connector struct {
	BaseDynamic
	OwnerID               uuid.UUID
	ConnectorDefinitionID uuid.UUID
	Name                  string
	Description           string
	Tombstone             bool
	Configuration         datatypes.JSON `gorm:"type:jsonb"`
	ConnectorType         ConnectorType  `sql:"type:valid_connector_type"`

	// Output-only field
	FullName string `gorm:"-"`
}

// Scan function for custom GORM type ConnectorType
func (c *ConnectorType) Scan(value interface{}) error {
	*c = ConnectorType(connectorPB.ConnectorType_value[value.(string)])
	return nil
}

// Value function for custom GORM type ConnectorType
func (c ConnectorType) Value() (driver.Value, error) {
	return connectorPB.ConnectorType(c).String(), nil
}

// Scan function for custom GORM type ConnectionType
func (c *ConnectionType) Scan(value interface{}) error {
	*c = ConnectionType(connectorPB.ConnectionType_value[value.(string)])
	return nil
}

// Value function for custom GORM type ConnectionType
func (c ConnectionType) Value() (driver.Value, error) {
	return connectorPB.ConnectionType(c).String(), nil
}

// Scan function for custom GORM type ConnectorType
func (r *ReleaseStage) Scan(value interface{}) error {
	*r = ReleaseStage(connectorPB.ReleaseStage_value[value.(string)])
	return nil
}

// Value function for custom GORM type ConnectorType
func (r ReleaseStage) Value() (driver.Value, error) {
	return connectorPB.ReleaseStage(r).String(), nil
}
