package services

import (
	"Go/models"

	"gorm.io/gorm"
)

type MemberService struct {
	DB *gorm.DB
}

func (s *MemberService) GetAllMembers() ([]models.Member, error) {
	var members []models.Member
	result := s.DB.Find(&members)
	return members, result.Error
}

func (s *MemberService) GetMemberByID(id int) (models.Member, error) {
	var member models.Member
	result := s.DB.First(&member, id)
	return member, result.Error
}

func (s *MemberService) CreateMember(member *models.Member) error {
	result := s.DB.Create(&member)
	return result.Error
}

func (s *MemberService) GetMemberByUsername(username string) (*models.Member, error) {
	var member models.Member
	// หา record แรกที่ username ตรงกับที่ส่งมา
	result := s.DB.Where("username = ?", username).First(&member)

	if result.Error != nil {
		return nil, result.Error
	}
	return &member, nil
}

// UpdateMember อัปเดตข้อมูลสมาชิก
func (s *MemberService) UpdateMember(member *models.Member) error {
	return s.DB.Save(member).Error
}

// DeleteMember ลบสมาชิกตาม ID
func (s *MemberService) DeleteMember(id int) error {
	return s.DB.Delete(&models.Member{}, id).Error
}
