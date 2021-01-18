package migrate

import "time"

// DBTransactionRoomBook will migrate a transaction room book table with the given specification into the database
type DBTransactionRoomBook struct {
	ID           uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	BookerID     uint      `gorm:"not null" json:"booker_id"`
	KostID       uint      `gorm:"not null" json:"kost_id"`
	RoomID       uint      `gorm:"not null" json:"room_id"`
	RoomDetailID uint      `gorm:"not null" json:"room_detail_id"`
	PeriodID     uint      `gorm:"not null" json:"period_id"`
	Status       uint      `gorm:"not null" json:"status"`
	BookCode     string    `gorm:"not null" json:"book_code"`
	BookDate     time.Time `gorm:"not null" json:"book_date"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	Created      time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy    string    `json:"created_by"`
	Modified     time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy   string    `json:"modified_by"`
}

// DBTransactionRoomBookMember will migrate a transaction room book member table with the given specification into the database
type DBTransactionRoomBookMember struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomBookID uint      `gorm:"not null" json:"room_book_id"`
	MemberName string    `gorm:"not null" json:"member_name"`
	Phone      string    `json:"phone"`
	Gender     bool      `gorm:"not null" json:"gender"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBTransactionRoomBookTable set the migrated struct table name
func (dbTransactionRoomBook *DBTransactionRoomBook) DBTransactionRoomBookTable() string {
	return "dbTransactionRoomBook"
}

// DBTransactionRoomBookMemberTable set the migrated struct table name
func (dbTransactionRoomBookMember *DBTransactionRoomBookMember) DBTransactionRoomBookMemberTable() string {
	return "dbTransactionRoomBookMember"
}
