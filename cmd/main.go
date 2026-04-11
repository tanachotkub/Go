// main.go
package main

import (
	_ "Go/docs"
	"Go/handlers"
	"Go/middlewares" // Import เพิ่ม
	"Go/routes"      // Import เพิ่ม
	"Go/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // Library สำหรับ .env
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Member API with Fiber and GORM
// @version 1.0
// @host localhost:3000
// @BasePath /api
func main() {
	// 1. โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. ดึงค่าจาก .env
	dsn := os.Getenv("DB_DSN")
	port := os.Getenv("PORT")

	// 3. เชื่อมต่อ Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// 4. Setup Dependency
	memberSrv := services.MemberService{DB: db}
	memberHdl := handlers.MemberHandler{Service: memberSrv}

	// 5. Setup Fiber
	app := fiber.New()

	// 6. ใช้ Middleware จากโฟลเดอร์ใหม่
	middlewares.SetupCORS(app)

	// 7. ใช้ Routes จากโฟลเดอร์ใหม่ (ส่ง Handler เข้าไปด้วย)
	routes.SetupRoutes(app, &memberHdl)

	// 8. เริ่มรัน Server ตาม Port ใน .env
	log.Fatal(app.Listen(":" + port))
}
