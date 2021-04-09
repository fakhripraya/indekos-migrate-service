package migrate

import "time"

// DBTransaction will migrate a transaction table with the given specification into the database
type DBTransaction struct {
	ID             uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	TrxReferenceID uint      `gorm:"not null" json:"trx_reference_id"`
	TrxCategory    uint      `gorm:"not null" json:"trx_category"` // kategori transaksi (bayar kost, bayar perpanjang, dll)
	PaidOff        float64   `gorm:"not null" json:"paid_off"`
	MustPay        float64   `gorm:"not null" json:"must_pay"`
	IsActive       bool      `gorm:"not null;default:true" json:"is_active"`
	Created        time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy      string    `json:"created_by"`
	Modified       time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy     string    `json:"modified_by"`
}

// DBTransactionDetail will migrate a transaction detail table with the given specification into the database
type DBTransactionDetail struct {
	ID              uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	TrxID           uint      `gorm:"not null" json:"trx_id"`
	PaymentMethodID uint      `gorm:"not null" json:"payment_method_id"`
	Status          uint      `gorm:"not null" json:"status"`
	Payment         float64   `gorm:"not null" json:"Payment"`
	IsActive        bool      `gorm:"not null;default:true" json:"is_active"`
	Created         time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy       string    `json:"created_by"`
	Modified        time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy      string    `json:"modified_by"`
}

// DBTransactionVerification is an entity that directly communicate with the DBTransactionVerification table in the database
type DBTransactionVerification struct {
	ID          uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	ReferenceID uint      `gorm:"not null" json:"reference_id"`
	PictDesc    string    `gorm:"not null" json:"pict_desc"`
	URL         string    `gorm:"not null" json:"url"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Created     time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy   string    `json:"created_by"`
	Modified    time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy  string    `json:"modified_by"`
}

// DBTransactionTable set the migrated struct table name
func (masterPaymentMethod *MasterPaymentMethod) DBTransactionTable() string {
	return "dbTransaction"
}

// DBTransactionDetailTable set the migrated struct table name
func (masterPaymentMethod *MasterPaymentMethod) DBTransactionDetailTable() string {
	return "dbTransactionDetail"
}

// DBTransactionVerificationTable set the migrated struct table name
func (dbTransactionVerification *DBTransactionVerification) DBTransactionVerificationTable() string {
	return "dbTransactionVerification"
}
