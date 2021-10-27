### CRUD api - Fiber - GORM
### Heroku app
- `https://golang-blog-api.herokuapp.com/api/post`
### Install
```
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
```
- Create database & name `blog`
- Run `go run main.go`
### Post
- GET - `http://localhost:3000/api/post`
- GET - `http://localhost:3000/api/post/:id`
- POST - `http://localhost:3000/api/post`
- PUT - `http://localhost:3000/api/post/:id`
- DELETE - `http://localhost:3000/api/post/:id`
### Tag
- GET - `http://localhost:3000/api/tag`
- POST - `http://localhost:3000/api/tag`
- PUT - `http://localhost:3000/api/tag/:id`
