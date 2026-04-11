package models

// Member represents the member table in database
type Member struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id" swaggerignore:"true"` // เพิ่มตรงนี้
	Username string `gorm:"column:username" json:"username" example:"johndoe" extensions:"x-order=1"`
	Password string `gorm:"column:password" json:"password" example:"secret123" extensions:"x-order=2"`
}

// TableName overrides the table name for GORM
func (Member) TableName() string {
	return "member"
}
