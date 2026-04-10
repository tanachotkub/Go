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
