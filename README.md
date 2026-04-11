# 🚀 Go Fiber + GORM + Swagger (Member API)

REST API มาตรฐานมืออาชีพ พัฒนาด้วยภาษา **Go** พร้อมระบบรักษาความปลอดภัยและการจัดการโครงสร้างที่เป็นระเบียบ

---

## ✨ ฟีเจอร์หลัก

- ⚡ **Fiber v2** - Web Framework ที่รวดเร็วและเบาที่สุด
- 🗄️ **GORM** - ORM สำหรับจัดการ MySQL พร้อมระบบ Auto Migration
- 📄 **Swagger UI** - สร้าง API Documentation อัตโนมัติด้วย Swaggo
- 🔐 **Bcrypt** - การเข้ารหัสรหัสผ่านที่ปลอดภัยระดับสากล
- ⚙️ **Dotenv** - จัดการค่าคอนฟิกผ่านไฟล์ .env

---

## 📂 โครงสร้างโปรเจกต์

โปรเจกต์นี้ใช้แนวคิด **Modular Layered Architecture** แยกหน้าที่ชัดเจน (SOC)

```
.
├── cmd/                    # Entry point: จุดเริ่มต้นของโปรแกรม
│   └── main.go
├── handlers/               # Delivery Layer: รับ/ส่ง HTTP + Swagger Annotation
├── services/               # Business Logic Layer: Logic & Database Access
├── models/                 # Data Models: Struct & Auto Migration
├── routes/                 # Routing Layer: จัดการเส้นทาง API ทั้งหมด
├── middlewares/            # Middlewares: CORS & Data Filtering
├── docs/                   # Swagger Files (Auto-generated)
├── .env                    # Environment Variables (Database, Port)
├── .gitignore              # Git ignore rules
├── go.mod                  # Dependency Management
└── README.md               # Documentation
```

---

## 🛠️ Tech Stack & Libraries

| Library | Description |
|---------|-------------|
| **Fiber v2** | High-performance Web Framework |
| **GORM** | The fantastic ORM library for Golang |
| **Bcrypt** | Password hashing algorithm |
| **Godotenv** | Load environment variables from .env |
| **Swaggo** | Automatically generate RESTful API documentation |

---

## 🚀 การติดตั้งและเริ่มใช้งาน

### 1️⃣ เตรียมความพร้อม (Prerequisites)

ติดตั้ง swag CLI เพื่อใช้สร้างเอกสาร API:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2️⃣ ตั้งค่าสภาพแวดล้อม (.env)

สร้างไฟล์ `.env` ที่ Root Directory และกำหนดค่าดังนี้:

```env
DB_DSN=root:your_password@tcp(127.0.0.1:3306)/your_db_name?charset=utf8mb4&parseTime=True&loc=Local
PORT=3000
```

**ตัวอย่าง:**
```env
DB_DSN=admin:password123@tcp(localhost:3306)/member_db?charset=utf8mb4&parseTime=True&loc=Local
PORT=8080
```

### 3️⃣ ติดตั้ง Dependencies และเรียกใช้

```bash
# ติดตั้ง library
go mod tidy

# รันโปรเจกต์ (ระบบจะทำการ Auto Migrate ตารางให้ทันที)
go run cmd/main.go
```

---

## 📄 การจัดการ API Documentation (Swagger)

ทุกครั้งที่มีการเพิ่มหรือแก้ไข Comment ใน Handlers ให้รันคำสั่งอัปเดต Docs:

```bash
# Windows PowerShell
& "$(go env GOPATH)\bin\swag" init -g cmd/main.go

# Linux / macOS
swag init -g cmd/main.go
```

**เข้าดู API Docs ได้ที่:**
```
http://localhost:3000/swagger
```

---

## 🛡️ ฟีเจอร์ที่น่าสนใจ

### ✅ 1. Auto Migration

ไม่ต้องเขียน SQL เอง! เมื่อคุณแก้ไข Struct ใน `models/` ระบบจะปรับโครงสร้างตารางใน MySQL ให้ตรงกันอัตโนมัติเมื่อรันโปรแกรม พร้อมระบบ Log แจ้งเตือนสถานะการ Migration

```go
// models/member.go ตัวอย่าง
type Member struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string `gorm:"uniqueIndex"`
}
```

### ✅ 2. Security First

- **Password Hashing:** เข้ารหัสผ่านด้วย Bcrypt ก่อนบันทึก (ห้ามเก็บ Plain Text)
- **CORS Middleware:** แยกไฟล์ไว้ที่ `middlewares/` เพื่อความสะอาด และอนุญาตเฉพาะ Header ที่จำเป็น
- **Hidden Sensitive Data:** ใช้ MemberRequest (DTO) เพื่อซ่อนฟิลด์ที่ไม่จำเป็นในหน้า Swagger

### ✅ 3. Clean Routes

ย้ายการประกาศเส้นทาง API ทั้งหมดไปไว้ใน `routes/routes.go` ทำให้ `main.go` สั้นและอ่านง่ายขึ้น

---
