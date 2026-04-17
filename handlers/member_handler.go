package handlers

import (
	"Go/models"
	"Go/services"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type MemberHandler struct {
	Service services.MemberService
}

// GetMembers godoc
// @Summary      Get all members
// @Tags         Members
// @Produce      json
// @Success      200  {array}   models.Member
// @Security     BearerAuth
// @Router       /members [get]
func (h *MemberHandler) GetMembers(c *fiber.Ctx) error {
	members, err := h.Service.GetAllMembers()
	if err != nil {
		return c.Status(500).JSON(map[string]any{"error": err.Error()})
	}
	return c.JSON(members)
}

// GetMemberByID godoc
// @Summary      Get member by ID
// @Tags         Members
// @Param        id   path      int  true  "Member ID"
// @Success      200  {object}  models.Member
// @Security     BearerAuth
// @Router       /members/{id} [get]
func (h *MemberHandler) GetMemberByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	member, err := h.Service.GetMemberByID(id)
	if err != nil {

		return c.Status(404).JSON(map[string]any{"error": "Member not found"})
	}
	return c.JSON(member)
}

// CreateMember godoc
// @Summary      Create a new member
// @Description  สร้างสมาชิกใหม่โดยรับข้อมูลจาก JSON body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        member  body      models.RegisterRequest  true  "Member Data"
// @Success      201     {object}  models.Member
// @Failure      400     {object}  map[string]any "Invalid request body"
// @Failure      500     {object}  map[string]any "Could not create member"
// @Router /register [post]
func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {
	req := new(models.RegisterRequest)

	// 1. Body Parser
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(map[string]any{"error": "รูปแบบข้อมูลไม่ถูกต้อง"})
	}

	// 2. Data Validation (ตรวจสอบเบื้องต้น)
	if len(req.Username) < 4 {
		return c.Status(400).JSON(map[string]any{"error": "Username ต้องมีอย่างน้อย 4 ตัวอักษร"})
	}
	if len(req.Password) < 8 {
		return c.Status(400).JSON(map[string]any{"error": "Password ต้องมีอย่างน้อย 8 ตัวอักษร"})
	}

	// 3. ตรวจสอบ Username ซ้ำในฐานข้อมูล
	existingMember, _ := h.Service.GetMemberByUsername(req.Username)
	if existingMember != nil {
		return c.Status(409).JSON(map[string]any{"error": "Username นี้ถูกใช้งานแล้ว"})
	}

	// 4. เข้ารหัส Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(map[string]any{"error": "เกิดข้อผิดพลาดภายในระบบ"})
	}

	// 5. บันทึกข้อมูล
	member := models.Member{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := h.Service.CreateMember(&member); err != nil {
		return c.Status(500).JSON(map[string]any{"error": "ไม่สามารถสร้างสมาชิกได้"})
	}

	// ส่งคืนเฉพาะข้อมูลที่จำเป็น (ไม่ส่ง Password กลับ)
	return c.Status(201).JSON(map[string]any{
		"message":  "สร้างสมาชิกสำเร็จ",
		"id":       member.ID,
		"username": member.Username,
	})
}

// Login godoc
// @Summary      Login to get JWT Token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login  body      models.LoginRequest  true  "Login Credentials"
// @Success      200    {object}  models.LoginResponse
// @Router       /login [post]
func (h *MemberHandler) LoginMember(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(map[string]any{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	token, member, err := h.Service.LoginMember(req)
	if err != nil {
		return c.Status(401).JSON(map[string]any{"error": err.Error()})
	}

	return c.JSON(models.LoginResponse{
		Token:  token,
		Member: member,
	})
}

// UpdateMember godoc
// @Summary      Update a member
// @Description  แก้ไขข้อมูลสมาชิก (เช่น เปลี่ยน Password) ตาม ID
// @Tags         Members
// @Accept       json
// @Produce      json
// @Param        id      path      int                   true  "Member ID"
// @Param        member  body      models.Member  true  "New Member Data"
// @Success      200     {object}  models.Member
// @Security     BearerAuth
// @Router       /members/{id} [put]
func (h *MemberHandler) UpdateMember(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	req := new(models.Member)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(map[string]any{"error": "รูปแบบข้อมูลไม่ถูกต้อง"})
	}

	// 1. หาข้อมูลเดิมก่อน
	member, err := h.Service.GetMemberByID(id)
	if err != nil {
		return c.Status(404).JSON(map[string]any{"error": "ไม่พบสมาชิกที่ต้องการแก้ไข"})
	}

	// 2. อัปเดตค่า (ถ้ามีการส่ง Password มาใหม่ ให้เข้ารหัสด้วย)
	member.Username = req.Username
	if req.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		member.Password = string(hashed)
	}

	// 3. บันทึกลง DB
	if err := h.Service.UpdateMember(&member); err != nil {
		return c.Status(500).JSON(map[string]any{"error": "ไม่สามารถอัปเดตข้อมูลได้"})
	}

	return c.JSON(member)
}

// DeleteMember godoc
// @Summary      Delete a member
// @Description  ลบสมาชิกออกจากระบบตาม ID
// @Tags         Members
// @Param        id   path      int  true  "Member ID"
// @Success      204  {object}  nil
// @Security     BearerAuth
// @Router       /members/{id} [delete]
func (h *MemberHandler) DeleteMember(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	// ตรวจสอบก่อนว่ามีของไหม
	_, err := h.Service.GetMemberByID(id)
	if err != nil {
		return c.Status(404).JSON(map[string]any{"error": "ไม่พบสมาชิกที่ต้องการลบ"})
	}

	if err := h.Service.DeleteMember(id); err != nil {
		return c.Status(500).JSON(map[string]any{"error": "ไม่สามารถลบข้อมูลได้"})
	}

	// 3. เปลี่ยนจาก 204 เป็น 200 เพื่อให้ส่ง JSON Body กลับไปได้
	return c.Status(200).JSON(map[string]any{
		"message": "ลบสมาชิกเรียบร้อยแล้ว",
		"id":      id,
	})
}
