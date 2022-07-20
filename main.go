package main

import (
	"adr/model"
	"adr/ui"
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

	db.AutoMigrate(&model.Project{}, &model.Record{}, &model.Comment{})

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

	api.Get("/projects", func(c *fiber.Ctx) error {
		var json []model.Project

		if err := db.Model(&model.Project{}).Order("updated_at desc").Find(&json).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Get("/project/:id", func(c *fiber.Ctx) error {
		var json model.Project

		if err := db.Model(&model.Project{}).Where("id = ?", c.Params("id")).First(&json).Error; err != nil {
			return err
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

	api.Get("/records", func(c *fiber.Ctx) error {
		var json []model.Record

		if err := db.Model(&model.Record{}).Order("created_at asc").Find(&json).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Get("/record/:project_id", func(c *fiber.Ctx) error {
		var json []model.Record

		if err := db.Model(&model.Record{}).Where("project_id = ?", c.Params("project_id")).Order("created_at asc").Find(&json).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Post("/record", func(c *fiber.Ctx) error {
		json := new(model.Record)
		if err := c.BodyParser(json); err != nil {
			return err
		}

		if err := validator.New().Struct(json); err != nil {
			return err
		}

		// TODO: valid project_id exisit

		// create or updte
		if err := db.Model(&model.Record{}).Where("id = ?", json.ID).Save(json).Error; err != nil {
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

	api.Delete("/record/:id", func(c *fiber.Ctx) error {
		if err := db.Delete(&model.Record{}, c.Params("id")).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
		})
	})

	api.Get("/comment", func(c *fiber.Ctx) error {
		var json []model.Comment

		if err := db.Model(&model.Comment{}).Order("created_at asc").Find(&json).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Get("/comment/:record_id", func(c *fiber.Ctx) error {
		var json []model.Comment

		if err := db.Model(&model.Comment{}).Where("record_id = ?", c.Params("record_id")).Order("created_at asc").Find(&json).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
			"data":  json,
		})
	})

	api.Post("/comment", func(c *fiber.Ctx) error {
		json := new(model.Comment)
		if err := c.BodyParser(json); err != nil {
			return err
		}

		if err := validator.New().Struct(json); err != nil {
			return err
		}

		// TODO: valid project_id exisit

		// create or updte
		if err := db.Model(&model.Comment{}).Where("id = ?", json.ID).Save(json).Error; err != nil {
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

	api.Delete("/comment/:id", func(c *fiber.Ctx) error {
		if err := db.Delete(&model.Comment{}, c.Params("id")).Error; err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"errNo": 0,
		})
	})

	app.Listen(":3000")
}
