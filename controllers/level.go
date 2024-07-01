package controllers

import (
	"fiber-admin/initializers"
	"fiber-admin/models"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func HandleErr(code int64, err string, c fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(Response{
		Code: code,
		Msg:  err,
		Data: nil,
	})
}

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetLevel(c fiber.Ctx) error {
	db := initializers.DB
	// levels := new([]models.Level)
	var levels []models.Level
	if err := db.Find(&levels).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	return c.Status(http.StatusOK).JSON(Response{
		Code: 200,
		Msg:  "Success",
		Data: levels,
	})
}

type ByIdReq struct {
	Id uint `json:"id" validate:"required"`
}

func GetLevelByID(c fiber.Ctx) error {

	var req ByIdReq

	if err := c.Bind().Body(&req); err != nil {
		return HandleErr(401, err.Error(), c)
	}

	db := initializers.DB
	var level models.Level
	if err := db.Debug().Take(&level, req.Id).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	return c.Status(http.StatusOK).JSON(Response{
		Code: 200,
		Msg:  "Success",
		Data: level,
	})
}

func DeleteLevelByID(c fiber.Ctx) error {
	var req ByIdReq

	if err := c.Bind().Body(&req); err != nil {
		return HandleErr(401, err.Error(), c)
	}
	db := initializers.DB
	var level models.Level
	if err := db.Debug().Delete(&level, req.Id).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	return c.Status(http.StatusOK).JSON(Response{
		Code: 200,
		Msg:  "Success",
		Data: level,
	})
}

func CreateLevel(c fiber.Ctx) error {
	db := initializers.DB

	// levels := new(models.Level) // level 是一个指向 models.Level 零值的指针
	var levels models.Level // levels 是一个 models.Level 类型的变量，初始化为零值

	if err := c.Bind().Body(&levels); err != nil {
		return HandleErr(401, err.Error(), c)
	}

	if err := db.Create(&levels).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	return c.Status(http.StatusOK).JSON(Response{
		Code: 200,
		Msg:  "Success",
		Data: levels,
	})
}

func UpdateLevelByID(c fiber.Ctx) error {
	var req ByIdReq

	if err := c.Bind().Body(&req); err != nil {
		return HandleErr(401, err.Error(), c)
	}
	db := initializers.DB
	var level models.Level
	if err := db.First(&level, req.Id).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	if err := c.Bind().Body(level); err != nil {
		return HandleErr(401, err.Error(), c)
	}

	if err := db.Save(&level).Error; err != nil {
		return HandleErr(401, err.Error(), c)
	}

	return c.Status(http.StatusOK).JSON(Response{
		Code: 200,
		Msg:  "Success",
		Data: level,
	})
}
