package handlers

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/module/application"
)

type GetPublicationHandler struct {
	useCase application.GetPublicationUseCase
}

func NewGetPublicationHandler(useCases *application.GetPublicationUseCase) *GetPublicationHandler {

	return &GetPublicationHandler{useCase: *useCases}
}

func (uc *GetPublicationHandler) RunGetPublicationHandler(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	// 2. Validar que no venga vacío y tenga el prefijo "Bearer "
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		slog.Error("token no proporcionado o formato inválido no authorizado")
		return config.NewStatusUnauthorized(errors.New("token no proporcionado o formato inválido"))
	}
	tokenString := authHeader[7:]

	responseUc, errorUsc := uc.useCase.ExecuteGetPublicationUseCase(tokenString)

	if errorUsc != nil {
		return errorUsc
	}

	return config.ResponseOk(c, responseUc, "Ok")
}
