package resp

import (
	"election-service/internal/core/models"

	"github.com/gofiber/fiber/v2"
)

type Response = models.Response

func Send(c *fiber.Ctx, resp Response) error {
	return c.Status(resp.Code).JSON(resp.Data)
}

func Created(data interface{}) Response {
	return Response{
		Code:    fiber.StatusCreated,
		Success: true,
		Message: "One or more new resources have been successfully created",
		Data:    &data,
	}
}

func OK(data interface{}) Response {
	return Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "The request has succeeded",
		Data:    &data,
	}
}

var (
	NoContent = Response{
		Code:    fiber.StatusNoContent,
		Success: true,
		Message: "A request has succeeded, but that the client doesn't need to navigate away from its current page",
	}

	BadRequestError = Response{
		Code:    fiber.StatusBadRequest,
		Success: false,
		Message: "The server cannot or will not process the request due to something that is perceived to be a client error.",
	}

	UnauthorizedError = Response{
		Code:    fiber.StatusUnauthorized,
		Success: false,
		Message: "The client request has not been completed because it lacks valid authentication credentials for the requested resource.",
	}

	ForbiddenError = Response{
		Code:    fiber.StatusForbidden,
		Success: false,
		Message: "The client request has not been completed because it lacks valid authentication credentials for the requested resource.",
	}

	NotFoundError = Response{
		Code:    fiber.StatusNotFound,
		Success: false,
		Message: "The server cannot find the requested resource.",
	}

	UnprocessableEntity = Response{
		Code:    fiber.StatusUnprocessableEntity,
		Success: false,
		Message: "The server understands the content type of the request entity, and the syntax of the request entity is correct, but it was unable to process the contained instructions.",
	}

	InternalServerError = Response{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
	}
)
