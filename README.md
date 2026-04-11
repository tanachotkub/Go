# 🚀 Go Fiber + GORM + Swagger (Member API)

โปรเจกต์เริ่มต้นสำหรับการสร้าง REST API ด้วยภาษา **Go** โดยใช้

* ⚡ **Fiber v2** เป็น Web Framework
* 🗄️ **GORM** สำหรับจัดการฐานข้อมูล MySQL
* 📄 **Swagger UI (Swaggo)** สำหรับสร้าง API Documentation อัตโนมัติ

---

# 📂 Project Structure

โปรเจกต์นี้ใช้แนวคิด **Modular Layered Architecture** เพื่อแยกหน้าที่แต่ละส่วนอย่างชัดเจน ทำให้ดูแลง่ายและขยายระบบได้สะดวก

```
.
├── cmd/                # Entry point: จุดเริ่มต้นของโปรแกรม
│   └── main.go         
├── handlers/           # Delivery Layer: รับ/ส่ง HTTP และเขียน Swagger Annotation
│   └── member_handler.go
├── services/           # Business Logic Layer: จัดการ Logic และติดต่อ Database
│   └── member_service.go
├── models/             # Data Models: โครงสร้าง Struct และ Schema
│   └── member.go
├── routes/             # Routing Layer: จัดการ Endpoint
│   └── routes.go
├── docs/               # Swagger Files (Auto-generated โดย swag CLI)
├── go.mod              # Dependency Management
└── .env                # (Optional) เก็บ Database Credentials
```

---

# 🛠️ Tech Stack

| Technology   | Description           |
| ------------ | --------------------- |
| **Go 1.2x+** | Programming Language  |
| **Fiber v2** | Web Framework         |
| **GORM**     | ORM สำหรับ MySQL      |
| **Swaggo**   | Swagger Documentation |

---

# 🚀 วิธีติดตั้งและรันโปรเจกต์

## 1️⃣ เตรียมความพร้อม (Prerequisites)

ติดตั้ง `swag CLI`

```powershell
go install github.com/swaggo/swag/cmd/swag@latest
```

ตรวจสอบว่าติดตั้งสำเร็จ:

```powershell
swag --version
```

---

## 2️⃣ ติดตั้ง Dependencies

```powershell
go mod tidy
```

---

## 3️⃣ สร้าง / อัปเดต Swagger Documentation

ทุกครั้งที่แก้ไข Comment ใน Handler ให้รันคำสั่งนี้ที่ Root Directory:

```powershell
# สำหรับ Windows (PowerShell)
& "$(go env GOPATH)\bin\swag" init -g cmd/main.go

# สำหรับ Windows (PowerShell)
swag init -g cmd/main.go
```

เมื่อสำเร็จจะมีโฟลเดอร์ `docs/` ถูกสร้างหรืออัปเดต

---

## 4️⃣ รัน Server

```powershell
go run cmd/main.go
```

Server จะทำงานที่:

```
http://localhost:3000
```

---

# 🛡️ การตั้งค่าที่สำคัญ (Key Fixes)

## ✅ 1. แก้ไขปัญหา CORS

เพื่อให้ Swagger UI เรียก API ได้โดยไม่เกิด `Fetch Error`

```go
app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders: "Origin, Content-Type, Accept",
}))
```

> ต้องวาง Middleware นี้ **ก่อนประกาศ Routes**

---

## ✅ 2. โครงสร้าง Database (MySQL)

```sql
CREATE TABLE `member` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NULL,
    `password` VARCHAR(255) NULL,
    PRIMARY KEY (`id`)
);
```

---

## ✅ 3. การเรียกใช้ Swagger UI

เปิด Browser ไปที่:

```
http://localhost:3000/swagger
```

---

# 📝 Troubleshooting

## ❌ undefined: fiber.H

หากพบ Error นี้:

```
undefined: fiber.H
```

ให้ใช้แทนด้วย:

```go
map[string]any{
    "message": "success",
}
```

เพื่อหลีกเลี่ยงปัญหา Cache ของ Go Compiler

---

## ❌ Swagger ขึ้น "Failed to fetch"

ให้ตรวจสอบว่า:

* ได้รัน `swag init` แล้ว
* มีการ Import ใน `main.go`:

```go
import _ "Go/docs"
```

---

## ❌ CORS Error

ตรวจสอบว่าได้เพิ่ม:

```go
app.Use(cors.New())
```

ไว้ก่อนบรรทัดประกาศ Routes

---

# 🎯 สรุป

โปรเจกต์นี้เป็น Template สำหรับสร้าง REST API ด้วย:

* ⚡ Fiber (Fast & Lightweight)
* 🗄️ GORM (ORM สำหรับ MySQL)
* 📄 Swagger (API Documentation อัตโนมัติ)

เหมาะสำหรับ:

* ระบบ Member Management
* Mini Project
* ระบบ Backend พื้นฐาน
* ใช้เป็น Boilerplate เริ่มต้นโปรเจกต์ใหม่

---

