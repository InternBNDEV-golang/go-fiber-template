package gateways

import (
	"go-fiber-template/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) GetAllUserData(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUsers()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h *HTTPGateway) CreateUser(ctx *fiber.Ctx) error {

	bodyData := entities.UserDataModel{}
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if bodyData.Username == "" || bodyData.Email == "" || bodyData.UserID == "" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if err := h.UserService.InsertNewUser(bodyData); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot insert new user account."})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}
