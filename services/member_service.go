package services

import (
	database "Go/config"
	"Go/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type MemberService struct {
	DB *gorm.DB
}

//	func (s *MemberService) GetAllMembers() ([]models.Member, error) {
//		var members []models.Member
//		result := s.DB.Find(&members)
//		return members, result.Error
//	}
func (s *MemberService) GetAllMembers() ([]models.Member, error) {
	ctx := context.Background()
	cacheKey := "members:all"
	var members []models.Member

	// 1. ลองดึงข้อมูลจาก Redis
	val, err := database.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &members); err == nil {
			// ✅ จุดที่ 1: แจ้งว่าดึงข้อมูลจาก Redis สำเร็จ
			log.Println("⚡ Redis: Cache Hit (Fetched from Redis)")
			return members, nil
		}
	}

	// 2. ถ้าไม่เจอใน Cache (Cache Miss)
	log.Println("🐢 Redis: Cache Miss (Fetching from MySQL...)") // ✅ จุดที่ 2: แจ้งว่าต้องไปพึ่ง MySQL

	result := s.DB.Find(&members)
	if result.Error != nil {
		return nil, result.Error
	}

	// 3. เก็บลง Redis พร้อมตั้งเวลา
	data, _ := json.Marshal(members)
	database.RedisClient.Set(ctx, cacheKey, data, 10*time.Minute)
	log.Println("💾 Redis: Data cached successfully") // ✅ จุดที่ 3: ยืนยันว่าเก็บลง Cache แล้ว

	return members, nil
}

// func (s *MemberService) GetMemberByID(id int) (models.Member, error) {
// 	var member models.Member
// 	result := s.DB.First(&member, id)
// 	return member, result.Error
// }

func (s *MemberService) GetMemberByID(id int) (models.Member, error) {
	ctx := context.Background()
	// สร้าง Key แยกตาม ID เช่น member:5
	cacheKey := fmt.Sprintf("member:%d", id)
	var member models.Member

	// 1. ลองดึงข้อมูลจาก Redis ดูก่อน
	val, err := database.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &member); err == nil {
			log.Printf("⚡ Redis: Cache Hit for ID %d", id)
			return member, nil
		}
	}

	// 2. ถ้าไม่เจอใน Cache -> ไปดึงจาก MySQL
	log.Printf("🐢 Redis: Cache Miss for ID %d (Fetching from MySQL...)", id)
	result := s.DB.First(&member, id)
	if result.Error != nil {
		return member, result.Error
	}

	// 3. เก็บลง Redis แยกตาม ID
	data, _ := json.Marshal(member)
	database.RedisClient.Set(ctx, cacheKey, data, 10*time.Minute)
	log.Printf("| 💾 Redis: Cached member ID %d", id)

	return member, nil
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
