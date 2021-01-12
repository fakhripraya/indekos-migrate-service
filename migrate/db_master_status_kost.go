package migrate

import "time"

// MasterStatusKost will migrate a master status kost table with the given specification into the database
type MasterStatusKost struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	StatusDesc string    `gorm:"not null" json:"status_desc"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterStatusKostTable set the migrated struct table name
func (masterStatusKost *MasterStatusKost) MasterStatusKostTable() string {
	return "dbMasterStatusKost"
}
