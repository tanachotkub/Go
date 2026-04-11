// main.go
package main

import (
	_ "Go/docs"
	"Go/handlers"
	"Go/middlewares" // Import เพิ่ม
	"Go/models"
	"Go/routes" // Import เพิ่ม
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
	//โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//ดึงค่าจาก .env
	dsn := os.Getenv("DB_DSN")
	port := os.Getenv("PORT")

	//เชื่อมต่อ Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	//เรียกใช้การจัดกลุ่ม Migration จากแพ็กเกจ models
	if err := models.MigrateDB(db); err != nil {
		log.Printf("Could not migrate database: %v", err)
	}

	//Setup Dependency
	memberSrv := services.MemberService{DB: db}
	memberHdl := handlers.MemberHandler{Service: memberSrv}

	//Setup Fiber
	app := fiber.New()

	//ใช้ Middleware จากโฟลเดอร์ใหม่
	middlewares.SetupCORS(app)

	//ใช้ Routes จากโฟลเดอร์ใหม่ (ส่ง Handler เข้าไปด้วย)
	routes.SetupRoutes(app, &memberHdl)

	//เริ่มรัน Server ตาม Port ใน .env
	log.Fatal(app.Listen(":" + port))
}
