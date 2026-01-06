package database

import (
	"errors"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/module/domains/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServicesDatabaseAdapter struct {
	dbgorm *gorm.DB
}

func NewServicesDatabase() *ServicesDatabaseAdapter {
	dsn := os.Getenv("URL_POSTGRESS_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Like{})
	if err != nil {
		slog.Error("Error en la migración", "error", err)
	}
	slog.Info("Conectado a la base de datos")

	return &ServicesDatabaseAdapter{
		dbgorm: db,
	}
}

func (sd *ServicesDatabaseAdapter) NewPublication(message string, accountId string) (entities.EntityNewPostResponse, error) {
	userUUID, err := uuid.Parse(accountId)
	if err != nil {
		return entities.EntityNewPostResponse{}, config.NewBadRequestError(errors.New("el ID de usuario no es un UUID válido"))
	}
	newPost := Post{
		UserID:  userUUID,
		Message: message,
	}

	result := sd.dbgorm.Create(&newPost)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return entities.EntityNewPostResponse{}, config.NewInternalServerError(result.Error)
	}

	return entities.EntityNewPostResponse{
		Message: "OK",
		Details: struct {
			PublicationId string "json:\"publication_id\""
		}{
			PublicationId: newPost.ID.String(),
		},
	}, nil
}

func (sd *ServicesDatabaseAdapter) LikePublication(PostId uuid.UUID, UserId uuid.UUID) (entities.EntityLikePostResponse, error) {

	newLike := Like{
		PostID: PostId,
		UserID: UserId,
	}

	result := sd.dbgorm.Create(&newLike)
	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				slog.Error(result.Error.Error())
				return entities.EntityLikePostResponse{},
					config.NewBadRequestError(errors.New("tu ya le diste like a esta publicacion"))
			}
		}
		return entities.EntityLikePostResponse{}, config.NewInternalServerError(result.Error)
	}
	return entities.EntityLikePostResponse{
		Message: "Ok",
		Details: struct {
			Like_id string "json:\"like_id\""
		}{
			Like_id: newLike.ID.String(),
		},
	}, nil
}

func (sd *ServicesDatabaseAdapter) GetPublications() (entities.EntityPublicationsResponse, error) {

	var posts []Post

	/* Obtenemos posts */
	result := sd.dbgorm.
		Order("created_at desc").
		Preload("Likes").        // Cargamos los likes para contarlos
		Preload("User").         // Cargamos el usuario
		Preload("User.Profile"). // Cargamos el perfil anidado
		Find(&posts)

	if result.Error != nil {
		slog.Error("Error consultando posts", "error", result.Error)
		return entities.EntityPublicationsResponse{}, config.NewInternalServerError(result.Error)
	}
	var response []map[string]interface{}
	for _, p := range posts {
		// Obtenemos los datos del usuario de forma segura
		var userEmail, userAlias string
		userEmail = "Usuario eliminado"
		if p.UserID != uuid.Nil {
			userEmail = p.User.Email
			userAlias = p.User.Profile.Alias
		}

		item := map[string]interface{}{
			"post_id":    p.ID,
			"message":    p.Message,
			"created_at": p.CreatedAt,
			"author": map[string]string{
				"email": userEmail,
				"alias": userAlias,
			},
			"likes_count": len(p.Likes),
		}
		response = append(response, item)
	}
	return entities.EntityPublicationsResponse{
		Message: "ddd",
		Details: struct {
			Data []map[string]interface{} "json:\"data\""
		}{
			Data: response,
		},
	}, nil
}
