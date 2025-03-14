package service

import (
	"encoding/json"
	"github.com/AkulinIvan/ToDo-crud/internal/dto"
	"github.com/AkulinIvan/ToDo-crud/internal/repo"
	"github.com/AkulinIvan/ToDo-crud/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Слой бизнес-логики. Тут должна быть основная логика сервиса

// Service - интерфейс для бизнес-логики
type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
}

type service struct {
	repo repo.Repository
	log  *zap.SugaredLogger
}

// NewService - конструктор сервиса
func NewService(repo repo.Repository, logger *zap.SugaredLogger) Service {
	return &service{
		repo: repo,
		log:  logger,
	}
}

// CreateTask - обработчик запроса на создание задачи
func (s *service) CreateTask(ctx *fiber.Ctx) error {
	var req TaskRequest

	// Десериализация JSON-запроса
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		s.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	// Валидация входных данных
	if vErr := validator.Validate(ctx.Context(), req); vErr != nil {
		return dto.BadResponseError(ctx, dto.FieldIncorrect, vErr.Error())
	}

	// Вставка задачи в БД через репозиторий
	task := repo.Task{
		Title:       req.Title,
		Description: req.Description,
	}
	taskID, err := s.repo.CreateTask(ctx.Context(), task)
	if err != nil {
		s.log.Error("Failed to insert task", zap.Error(err))
		return dto.InternalServerError(ctx)
	}

	response := dto.StatusOK(ctx, map[string]int{"task_id": taskID})

	return response
}

func (s *service) GetTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Failed id of task", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	
	task, err := s.repo.GetTask(ctx.Context(), uint32(id))
	if err != nil {
		s.log.Error("Have not task in DB", zap.Error(err))
		return dto.InternalServerError(ctx)
	}

	response := dto.StatusOK(ctx, task)

	return response
}

func (s *service) UpdateTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Have not task in DB", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	var req TaskRequest

	// Десериализация JSON-запроса
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		s.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	// Валидация входных данных
	if vErr := validator.Validate(ctx.Context(), req); vErr != nil {
		return dto.BadResponseError(ctx, dto.FieldIncorrect, vErr.Error())
	}

	status := req.Status

	err = s.repo.UpdateTask(ctx.Context(), uint32(id), status)
	if err != nil {
		s.log.Error("Failed to update task", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to update task")
	}

	response := dto.StatusOK(ctx, "Success updating")

	return response
}

func (s *service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Have not id in BD", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	err = s.repo.DeleteTask(ctx.Context(), uint32(id))
	if err != nil {
		s.log.Error("Failed to delete of task", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to delete of task")
	}

	response := dto.StatusOK(ctx, "Delete was successful")

	return response
}

func (s *service) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := s.repo.GetTasks(ctx.Context())
	if err != nil {
		s.log.Error("Failed to select all tasks from table", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to select all tasks from table")
	}

	response := dto.StatusOK(ctx, tasks)

	return response
}