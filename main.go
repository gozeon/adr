package main

import (
	"adr/model"
	"adr/ui"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLog "gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/adr?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: gLog.Default.LogMode(gLog.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Project{})

	app := fiber.New(fiber.Config{
		// https://docs.gofiber.io/guide/error-handling
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(200).JSON(fiber.Map{
				"errMsg": err.Error(),
				"errNo":  code,
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	// https://docs.gofiber.io/api/middleware/filesystem
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(ui.DistDirFS),
		MaxAge: 60 * 60,
	}))

	// https://docs.gofiber.io/guide/grouping
	api := app.Group("/api")

	api.Get("/project", func(c *fiber.Ctx) error {
		var json []model.Project

		err := db.Model(&model.Project{}).Order("updated_at desc").Find(&json)
		if err != nil {
			fmt.Println(err)
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Post("/project", func(c *fiber.Ctx) error {
		json := new(model.Project)
		if err := c.BodyParser(json); err != nil {
			return err
		}

		if err := validator.New().Struct(json); err != nil {
			return err
		}

		// create or updte
		if err := db.Model(&model.Project{}).Where("id = ?", json.ID).Save(json).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(json)
			} else {
				return err
			}
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Delete("/project/:id", func(c *fiber.Ctx) error {

		if err := db.Delete(&model.Project{}, c.Params("id")).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
		})
	})

	app.Listen(":3000")
}
