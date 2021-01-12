package migrate

import "time"

// MasterRole will migrate a master role table with the given specification into the database
type MasterRole struct {
	ID         uint      `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	RoleName   string    `gorm:"unique;not null" json:"role_name"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterRoleTable set the migrated struct table name
func (masterRole *MasterRole) MasterRoleTable() string {
	return "dbMasterRole"
}
