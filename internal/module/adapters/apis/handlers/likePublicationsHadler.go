package handlers

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/module/adapters/apis/handlers/dto"
	"github.com/postredsocial/internal/module/application"
)

type LikePublicationHandler struct {
	useCase application.LikePublicationUseCase
}

func NewLikePublicationHandler(useCases *application.LikePublicationUseCase) *LikePublicationHandler {

	return &LikePublicationHandler{useCase: *useCases}
}

func (uc *LikePublicationHandler) RunLikePublicationHandler(c *fiber.Ctx) error {

	var data dto.DtoLikePublication

	if err := c.BodyParser(&data); err != nil {
		slog.Error(err.Error())
		return config.NewUnprStatusUnprocessableEntity(err)
	}
	authHeader := c.Get("Authorization")

	// 2. Validar que no venga vacío y tenga el prefijo "Bearer "
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		slog.Error("token no proporcionado o formato inválido no authorizado")
		return config.NewStatusUnauthorized(errors.New("token no proporcionado o formato inválido"))
	}
	tokenString := authHeader[7:]

	responseUc, errorUsc := uc.useCase.ExecuteLikePublicationeCase(data.PostId, data.UserId, tokenString)

	if errorUsc != nil {
		return errorUsc
	}

	return config.ResponseOk(c, responseUc, "Ok")
}
