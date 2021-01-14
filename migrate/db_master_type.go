package migrate

import "time"

// MasterKostType will migrate a master kost type table with the given specification into the database
type MasterKostType struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	TypeDesc   string    `gorm:"not null" json:"type_desc"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterKostTypeTable set the migrated struct table name
func (masterKostType *MasterKostType) MasterKostTypeTable() string {
	return "dbMasterKostType"
}
