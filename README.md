# 🛒 Full-Stack Shopping Cart Application

## A modern full-stack e-commerce shopping website built using:
| Layer      | Tech                             |
| ---------- | -------------------------------- |
| Frontend   | React (Vite), Axios, Context API |
| Backend    | Go (Gin), GORM, SQLite           |
| Deployment | Render                           |
| Auth       | JWT-based authentication         |
| Database   | SQLite (auto-migrated)           |

## 🚀 Live Demo
| Service       | URL                                                                                                           |
| ------------- | ------------------------------------------------------------------------------------------------------------- |
| Frontend      | 🔗 ***[https://your-frontend-url.onrender.com](https://your-frontend-url.onrender.com)***                     |
| Backend (API) | 🔗 **[https://shopping-cart-fullstack-gl1h.onrender.com](https://shopping-cart-fullstack-gl1h.onrender.com)** |


# ✨ Features
## 👤 User

 - Signup, Login & Logout (JWT)

 - Secure protected routes

## 🛍️ Products

 - Fetch all items

- Search products

- Add to cart from Home page

## 🛒 Cart

- View user cart

- Add items / Remove items

- Cart badge indicates number of items

- Checkout converts cart → order

## 📦 Orders

- View order history


# 📌 Tech Stack Highlights
| Category          | Tools                       |
| ----------------- | --------------------------- |
| Backend Framework | Gin                         |
| ORM               | GORM                        |
| Database          | SQLite                      |
| Authentication    | JWT in Authorization header |
| Frontend          | React + Context API         |
| Styling           | Custom CSS                  |
| Deployment        | Render (Backend & Frontend) |


# 🧪 Postman Collection

- A Postman API testing collection is included in this repository:

## 📁 postman_collection.json

Import it in Postman → click Collections → Import → Choose File.

The collection includes:

 - Signup

- Login

- Create item

- List items

- Add to cart

- Remove from cart

- List carts

- Create order

- Get orders

# 🛠️ Run the Project Locally
## 🔹 1. Clone the repo
```
 git clone https://github.com/Ashishkr8802/shopping-cart-fullstack.git
cd shopping-cart-fullstack
```
## 🔹 2. Run Backend
```
cd backend
go mod tidy
go run ./cmd/api
```
- Backend will start at: http://localhost:8080

- SQLite DB file will be generated automatically (e.g. gorm.db)

## 🔹 3. Run Frontend
```
cd shopping-frontend
npm install
npm run dev
```
Frontend will start at: http://localhost:5173

Ensure .env contains:
```
VITE_API_BASE_URL=http://localhost:8080
```
# 🔒 Authentication Guide
| Header Name     | Value     |
| --------------- | --------- |
| `Authorization` | `<token>` |

# 🧵 Available API Endpoints
## 👤 Users
| Method | Endpoint       | Description    |
| ------ | -------------- | -------------- |
| POST   | `/users`       | Create user    |
| GET    | `/users`       | List all users |
| POST   | `/users/login` | Login          |

## 🛍️ Items
| Method | Endpoint | Description             |
| ------ | -------- | ----------------------- |
| POST   | `/items` | Create item (admin use) |
| GET    | `/items` | List items              |

## 🛒 Cart (Protected)
| Method | Endpoint          | Description           |
| ------ | ----------------- | --------------------- |
| POST   | `/carts`          | Add item to cart      |
| GET    | `/carts`          | Get user's carts      |
| DELETE | `/cart-items/:id` | Remove item from cart |

## 📦 Orders (Protected)
| Method | Endpoint  | Description   |
| ------ | --------- | ------------- |
| POST   | `/orders` | Checkout cart |
| GET    | `/orders` | List orders   |

# 🙌 Credits

- Developed by Ashish Kumar
## If you like this project, don’t forget to ⭐ star the repo!



