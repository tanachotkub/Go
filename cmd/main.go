// main.go
package main

import (
	database "Go/config"
	_ "Go/docs"
	"Go/handlers"
	"Go/middlewares"
	"Go/models"
	"Go/routes"
	"Go/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Member API with Fiber and GORM
// @version 1.0

// @securityDefinitions.apiKey BearerAuth
// @in                         header
// @name                       Authorization
// @description                ใส่ Token ในรูปแบบ: Bearer <your_token>

// @host localhost:3000
// @BasePath /api
// @schemes   http https

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

	//เชื่อมต่อ Redis
	if err := database.ConnectRedis(); err != nil {
		log.Printf("Redis Connect Warning: %v", err)
	} else {
		log.Println("Redis Connected Successfully!")
	}

	//เรียกใช้การจัดกลุ่ม Migration จากแพ็กเกจ models
	if err := models.MigrateDB(db); err != nil {
		log.Printf("Could not migrate database: %v", err)
	}

	//Setup Dependency
	memberSrv := services.MemberService{DB: db}
	memberHdl := handlers.MemberHandler{Service: memberSrv}

	//Setup Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})
	app.Use(recover.New())

	//เพิ่ม Logger
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] | ${status} | ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	// ป้องกัน HTTP Header Attacks
	app.Use(helmet.New())

	//ใช้ Middleware จากโฟลเดอร์ใหม่
	middlewares.SetupCORS(app)

	//ใช้ Routes จากโฟลเดอร์ใหม่ (ส่ง Handler เข้าไปด้วย)
	routes.SetupRoutes(app, &memberHdl)

	//เริ่มรัน Server ตาม Port ใน .env
	log.Fatal(app.Listen(":" + port))
}
