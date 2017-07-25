package http

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/nykxs/fitworld"
	"github.com/nykxs/fitworld/http/validator"
)

type userHandler struct {
	userService fitworld.UserService
}

// RegisterUserHandler will handle all required actions to expose http endpoints to interact with users.
func RegisterUserHandler(s *Server) {
	userHandler := &userHandler{
		userService: s.UserService,
	}

	usersGroup := s.Router.Group("/users")

	usersGroup.POST("/register", userHandler.Register)
	usersGroup.GET("/:id", userHandler.Get)

	// Endpoints related to the current user - should be registered.
	usersGroup.GET("/me", userHandler.Me)
	usersGroup.GET("/me/delete", userHandler.Delete)
}

// UserRegisterPayload can be Bind in the Register endpoint and defines parameters that are both received and validated
type UserRegisterPayload struct {
	FirstName string `json:"first_name" form:"firstname" query:"firstname"`
	LastName  string `json:"last_name" form:"firstname" query:"firstname"`
	Email     string `json:"email" form:"email" query:"email" valid:"required,email"`
	Password  string `json:"password" form:"password" query:"password" valid:"required"`
}

// Register endpoint validates parameters and calls the UserService trying to create a new user.
func (h *userHandler) Register(c echo.Context) error {
	payload := new(UserRegisterPayload)
	if err := c.Bind(payload); err != nil {
		return err
	}

	validator := httpvalidator.NewValidator()
	if err := validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := fitworld.UserRegister{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	}

	stored, err := h.userService.Register(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, stored)
}

// UserGetPayload can be Bind in the Get endpoint and defines parameters that are both received and validated
type UserGetPayload struct {
	ID string `json:"id" valid:"required"`
}

// Get endpoint validates parameters and calls the UserService trying to get the user associated with an id.
func (h *userHandler) Get(c echo.Context) error {
	payload := UserGetPayload{
		ID: c.Param("id"),
	}

	validator := httpvalidator.NewValidator()
	if err := validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.userService.GetByID(payload.ID)
	if err != nil {
		if err == fitworld.ErrUserNotFound {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// Me endpoint validates parameters and calls the UserService trying to get info for the current user.
func (h *userHandler) Me(c echo.Context) error {
	IDStored := c.Get(ContextKeyCurrentUser)
	ID, ok := IDStored.(string)
	if !ok || ID == "" {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	user, err := h.userService.GetByID(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, user)
}

// Delete endpoint validates parameters and calls the UserService trying to delete the current user.
func (h *userHandler) Delete(c echo.Context) error {
	IDStored := c.Get(ContextKeyCurrentUser)
	ID, ok := IDStored.(string)
	if !ok || ID == "" {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	err := h.userService.Delete(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
