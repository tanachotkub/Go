package models

// Member represents the member table in database
type Member struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
	Username string `gorm:"column:username" json:"username" example:"johndoe"`
	Password string `gorm:"column:password" json:"password" example:"secret123"`
}

// TableName overrides the table name for GORM
func (Member) TableName() string {
	return "member"
}
