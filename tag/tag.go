package tag

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "b84d42076320a0:7b1f1e3e@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_ff75c840cfdb1d8?charset=utf8mb4&parseTime=True&loc=Local"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}

	DB.AutoMigrate(&Tag{})
}

func GetTags(c *fiber.Ctx) error {
	var tags []Tag
	DB.Find(&tags)
	return c.JSON(&tags)
}

func GetTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag Tag
	DB.Find(&tag, id)
	return c.JSON(tag)
}

func NewTag(c *fiber.Ctx) error {
	tag := new(Tag)
	if err := c.BodyParser(tag); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&tag)
	return c.JSON(&tag)
}

func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	var tag Tag
	DB.First(&tag, id)
	if tag.Name == "" {
		return c.Status(500).SendString("Tag not found")
	}

	DB.Delete(&tag)
	return c.SendString("Tag deleted")
}

func UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tag := new(Tag)
	DB.First(&tag, id)
	if tag.Name == "" {
		return c.Status(500).SendString("Tag not found")
	}
	if err := c.BodyParser(&tag); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&tag)
	return c.JSON(&tag)
}
