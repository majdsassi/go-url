# Go-Url

A modern, self-hosted URL shortener built with Go, Gin, GORM, and HTMX.
## 💻Demo 
![image](https://github.com/user-attachments/assets/ae72d36d-6c8c-4cb6-a95d-87838d9d5932)


## ✨ Features

- ✅ Create and manage shortened URLs
- 🚀 Fast redirects using Gin
- 🗃️ Persistent storage with GORM and MySQL 
- ⚡ Dynamic frontend updates via HTMX (no page reloads)
- 🔐 Unique shortcodes generation
- 📊 Link Visits

## 🏗️ Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - High-performance web framework
- [GORM](https://gorm.io/) - ORM for Go
- [HTMX](https://htmx.org/) - Modern frontend interaction without full-page reloads

## 🧪 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/majdsassi/go-url.git
cd go-url
```
### 2 . Install Dependencies 
```go
go mod tidy
```
### 3. Create your local variables
```env
DBURL="db_user:db_password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"
```
(if your db does require a password leave it empty and if ur db is remote change the tcp values with db_url:port ) 
### 5. Migrate  Your Database
```go
go run migration/migrate.go
```
### 6. Run Your App 
```go
go run main.go
```
OR Run it Using [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
```bash
CompileDaemon -command="./go-url" 
```
## 📄 License

MIT License. See [LICENSE](https://medium.com/@avinashvagh/github-licenses-explained-a-quick-guide-46d98ef4ca81)for details.
## 🤝 Contributing
Pull requests welcome. If you want to add features like analytics, expiration, or QR code generation—go for it!


