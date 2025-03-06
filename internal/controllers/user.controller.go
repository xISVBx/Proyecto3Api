package controllers

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{
		userService: s,
	}
}

// Login autentica a un usuario y devuelve un token JWT.
// @Summary Iniciar sesión
// @Description Permite a los usuarios autenticarse en la plataforma.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dtos.LoginDto true "Credenciales de usuario"
// @Security BearerAuth
// @Success 200 {object} models.AppResponse{data=string}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 401 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/login [post]
func (ct UserController) Login(c *gin.Context) {
	var dto dtos.LoginDto
	err := c.ShouldBindJSON(&dto)

	if err != nil {
		for k, v := range c.Request.Header {
			fmt.Printf("%s: %s\n", k, v)
		}
		token := c.GetHeader("Authorization")
		fmt.Println("Este es el token")
		fmt.Println(token)
		if token == "" {
			c.JSON(http.StatusBadRequest, models.ResError("Email and Password are required, or provide a valid Token"))
			return
		}
	}

	token := c.GetHeader("Authorization")

	fmt.Println("Este es el token")

	for k, v := range c.Request.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	dto.Token = token

	jwt, aErr := ct.userService.LoginService(dto)

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(jwt)
	c.JSON(http.StatusOK, response)
}

// Register permite a los usuarios registrarse en la plataforma.
// @Summary Registro de usuario
// @Description Crea una nueva cuenta para el usuario con los datos proporcionados.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dtos.RegisterRequestDto true "Datos del usuario"
// @Success 201 {object} models.AppResponse{data=dtos.RegisterResponseDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/register [post]
func (ct UserController) Register(c *gin.Context) {
	var dto dtos.RegisterRequestDto
	err := c.Bind(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError("Invalid data"))
		return
	}

	newUser, aErr := ct.userService.RegisterService(dto)

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(newUser)
	c.JSON(http.StatusOK, response)
}
