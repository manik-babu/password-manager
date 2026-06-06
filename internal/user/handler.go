package user

import (
	"detrox/internal/httpresponse"
	"detrox/internal/user/dto"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, httpresponse.Error{
			Ok:      false,
			Code:    400,
			Message: "User creation failed",
			Details: err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, httpresponse.Error{
			Ok:      false,
			Code:    400,
			Message: "User creation failed",
			Details: err.Error(),
		})
	}
	// calling service layer here
	res, err := h.service.CreateUser(req)

	if err != nil {
		return c.JSON(500, httpresponse.Error{
			Ok:      false,
			Code:    500,
			Message: "User creation failed",
			Details: err.Error(),
		})
	}
	return c.JSON(201, httpresponse.Success{
		Ok:      true,
		Message: "User created",
		Data:    res,
	})

}
