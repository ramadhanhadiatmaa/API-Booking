package controllers

import (
	"apibooking/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var booking []models.Booking

	models.DB.Find(&booking)

	return c.Status(fiber.StatusOK).JSON(booking)

}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")

	var booking models.Booking

	if err := models.DB.Model(&booking).Where("id = ?", id).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Kamar dengan id " + id + " tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan",
		})
	}

	return c.JSON(booking)
}

func Create(c *fiber.Ctx) error {

	var booking models.Booking

	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	
	booking.Created = time.Now()
	booking.Updated = time.Now()

	var current = time.Now()
	var unix = current.Unix()
	var unixTime = int(unix)

	var nomor = booking.Ktp
	var ktp = strconv.Itoa(nomor)
	var uniq = strconv.Itoa(unixTime)

	booking.Id = ktp + uniq

	if err := models.DB.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data pesanan berhasil dibuat",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var booking models.Booking

	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	booking.Updated = time.Now()

	if models.DB.Where("id = ?", id).Updates(&booking).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diperbaharui",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var booking models.Booking

	if models.DB.Where("id = ?", id).Delete(&booking).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak dapat dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data kamar",
	})
}