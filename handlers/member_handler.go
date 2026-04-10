package handlers

import (
	"Go/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	Service services.MemberService
}

// GetMembers godoc
// @Summary      Get all members
// @Tags         Members
// @Produce      json
// @Success      200  {array}   models.Member
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
// @Router       /members/{id} [get]
func (h *MemberHandler) GetMemberByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	member, err := h.Service.GetMemberByID(id)
	if err != nil {

		return c.Status(404).JSON(map[string]any{"error": "Member not found"})
	}
	return c.JSON(member)
}
