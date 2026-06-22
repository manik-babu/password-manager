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

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, httpresponse.Error{
			Ok:      false,
			Message: "Invalid data",
			Details: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(400, httpresponse.Error{
			Ok:      false,
			Message: "Invalid data",
			Details: err.Error(),
		})
	}
	res, err := h.service.LoginUser(req)
	if err != nil {
		return c.JSON(500, httpresponse.Error{
			Ok:      false,
			Code:    500,
			Message: err.Error(),
			Details: err.Error(),
		})
	}
	return c.JSON(200, httpresponse.Success{
		Ok:      true,
		Message: "Login successful",
		Data:    res,
	})
}
func (h *handler) GetMe(c *echo.Context) error {
	// user, ok := c.Get("user").(auth.JWTClaims)
	// fmt.Println("User info:", user)
	// if !ok {
	// 	return c.JSON(401, httpresponse.Error{
	// 		Ok:      false,
	// 		Code:    401,
	// 		Message: "Can not get user info",
	// 	})
	// }
	user := c.Get("user")
	return c.JSON(200, httpresponse.Success{
		Ok:      true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}
