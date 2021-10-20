package post

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "b84d42076320a0:7b1f1e3e@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_ff75c840cfdb1d8?charset=utf8mb4&parseTime=True&loc=Local"

type Post struct {
	gorm.Model
	Title string  `json:"title"`
	Body  string  `json:"body"`
	Tags  []*Post `gorm:"many2many:post_tag"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}

	DB.AutoMigrate(&Post{})
}

func GetPosts(c *fiber.Ctx) error {
	var posts []Post
	DB.Find(&posts)
	return c.JSON(&posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post Post
	DB.Find(&post, id)
	return c.JSON(post)
}

func NewPost(c *fiber.Ctx) error {
	post := new(Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&post)
	return c.JSON(&post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post Post
	DB.First(&post, id)
	if post.Title == "" {
		return c.Status(500).SendString("Post not found")
	}

	DB.Delete(&post)
	return c.SendString("Post deleted")
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(Post)
	DB.First(&post, id)
	if post.Title == "" {
		return c.Status(500).SendString("Post not found")
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&post)
	return c.JSON(&post)
}
