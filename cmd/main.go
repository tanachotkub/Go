package main

import (
	_ "Go/docs"
	"Go/handlers"
	"Go/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // 1. Import ตัวนี้เข้าไป
	"github.com/gofiber/swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Member API with Fiber and GORM
// @version 1.0
// @host localhost:3000
// @BasePath /api
func main() {
	// 1. Connection Database
	dsn := "root:1234@tcp(127.0.0.1:3306)/kenys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// 2. Setup Dependency
	memberSrv := services.MemberService{DB: db}
	memberHdl := handlers.MemberHandler{Service: memberSrv}

	// 3. Setup Fiber
	app := fiber.New()

	// ใช้ CORS Middleware (วางไว้ก่อน Routes เสมอ)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // อนุญาตทุก Domain (เหมาะสำหรับช่วงพัฒนา)
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API Routes
	api := app.Group("/api")
	api.Get("/members", memberHdl.GetMembers)
	api.Get("/members/:id", memberHdl.GetMemberByID)

	log.Fatal(app.Listen(":3000"))
}
