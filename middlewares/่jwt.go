package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
	// 1. ดึง Token จาก Header "Authorization"
	authHeader := c.Get("Authorization") // รูปแบบที่ส่งมาจะเป็น "Bearer <token>"

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing authorization header",
		})
	}

	// 2. ตัดคำว่า "Bearer " ออกเพื่อเอาเฉพาะตัว Token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// 3. ตรวจสอบความถูกต้องของ Token (Verify)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่าใช้วิธีการ Signing แบบที่คาดหวังไหม (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// 4. (Optional) เก็บข้อมูลจาก Token ลงใน Context เพื่อให้ Handler เรียกใช้ต่อได้
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])

	return c.Next() // ผ่าน! ไปยัง Handler ตัวถัดไปได้
}
