# 🍔 Order Food Online API (Go + Fiber)

A production-grade backend API for an **online food ordering platform**,  
built using [Go Fiber](https://github.com/gofiber/fiber) and compliant with **OpenAPI 3.1** specifications.

## 🧩 Overview

This project implements a REST API for a food ordering system with:

- 🧁 Product listing and retrieval  
- 🛒 Order creation with promo code validation  
- 🧾 File-based storage (JSON files for products and orders)  
- ☁️ AWS S3 Select–based coupon validation (no need to load huge files)  
- 🔑 API key–secured endpoints  
- ✅ OpenAPI 3.1 documentation  
- 🧠 Modular clean architecture with `handlers`, `store`, and `router`

## ⚙️ Tech Stack

| Component | Technology |
|------------|-------------|
| Language | Go 1.21+ |
| Framework | Fiber v2 |
| API Spec | OpenAPI 3.1 |
| Storage | JSON files |
| Promo Validation | AWS S3 Select |
| Auth | Header api_key |
| Testing | go test, testify/mock |

## 🔐 Authentication

Add header in all API requests:
```
api_key: apitest
```

## 🧠 Promo Code Validation

Validated via AWS S3 Select (server-side query).  
Must appear in **at least two** of these files:
- couponbase1.gz
- couponbase2.gz
- couponbase3.gz

## 🧰 Setup

```bash
git clone https://github.com/arpit2425/order-food-api.git
cd order-food-api
go mod tidy
export API_KEY=apitest
make run
```

## 🧪 Testing

```bash
make test
```

## 🧾 Example Endpoints

| Method | Endpoint | Description |
|--------|-----------|-------------|
| GET | /api/products | List all products |
| GET | /api/products/:id | Get product by ID |
| POST | /api/orders | Place a new order |

## 👨‍💻 Author

**Developer:** Arpit Trivedi  
📧 Email: arpittrivedi2425@gmail.com

### ❤️ Built with Go, Fiber, and caffeine.
