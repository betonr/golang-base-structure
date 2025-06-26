package main

import (
	"context"

	"github.com/betonr/golang-base-structure/internal/database/memory"
	"github.com/betonr/golang-base-structure/internal/service"
	"github.com/betonr/golang-base-structure/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	repo := memory.NewMemoryArticleRepository()
	svc := service.NewArticleService(repo)

	app.Get("/articles", func(c *fiber.Ctx) error {
		articles, err := svc.ListArticles(context.Background())
		if err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(articles)
	})

	app.Get("/articles/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		article, err := svc.GetArticleByID(context.Background(), id)
		if err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
		}
		return c.JSON(article)
	})

	app.Post("/articles", func(c *fiber.Ctx) error {
		var input service.CreateArticleInput
		if err := c.BodyParser(&input); err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		author := "Carlos Noronha" // TODO: Get author from context
		article, err := svc.CreateArticle(context.Background(), input.Title, input.Content, author)
		if err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(201).JSON(article)
	})

	app.Put("/articles/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var input service.UpdateArticleInput
		if err := c.BodyParser(&input); err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		article := svc.UpdateArticle(context.Background(), id, input.Title, input.Content)
		return c.JSON(article)
	})

	app.Delete("/articles/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := svc.DeleteArticle(context.Background(), id)
		if err != nil {
			logger.Error("Erro: %v", err)
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	})

	logger.Info("Servidor iniciado na porta %d", 8080)
	app.Listen(":8080")
}
