package controllers

import (
	"Key_Value_Storage/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var Db *models.Store

func GetKeyValue(c *fiber.Ctx) error {

	var cKey map[string]string
	if err := c.BodyParser(&cKey); err != nil {
		fmt.Println(err)
	}
	return c.JSON(Db.Get(cKey["key"]))
}

func SetKeyValue(c *fiber.Ctx) error {
	var cKeyValue models.KeyValue

	if err := c.BodyParser(&cKeyValue); err != nil {
		fmt.Println(err)
		return err
	}
	Db.Set(cKeyValue.Key, cKeyValue.Value)
	return c.JSON(cKeyValue)
}

func GetAllKeyValue(c *fiber.Ctx) error {
	return c.JSON(Db.GetAll())
}

func FlushAllData(c *fiber.Ctx) error {
	Db.FlushAllData()
	return c.JSON(Db.GetAll())
}
