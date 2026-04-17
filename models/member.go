package models

// Member represents the member table in database
type Member struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id" swaggerignore:"true"` // เพิ่มตรงนี้
	Username string `gorm:"column:username" json:"username" example:"johndoe" extensions:"x-order=1"`
	Password string `gorm:"column:password" json:"-" example:"secret123" extensions:"x-order=2"`
}

type LoginRequest struct {
	Username string `json:"username" example:"admin" extensions:"x-order=1"`
	Password string `json:"password" example:"123456" extensions:"x-order=2"`
}

type RegisterRequest struct {
	Username string `json:"username" example:"admin" extensions:"x-order=1"`
	Password string `json:"password" example:"123456" extensions:"x-order=2"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Member Member `json:"member"`
}

// TableName overrides the table name for GORM
func (Member) TableName() string {
	return "member"
}
