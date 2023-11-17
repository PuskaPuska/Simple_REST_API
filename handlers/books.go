package handlers

import (
	"context"

	"github.com/PuskaPuska/Simple_REST_API/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type newBookDTO struct {
	Title     string `json:"Title" bson:"title"`
	Author    string `json:"Author" bson:"author"`
	ISBN      string `json:"ISBN" bson:"isbn"`
	LibraryId string `json:"LibraryId" bson:"libraryId"`
}

func DeleteBook(c *fiber.Ctx) error {
	// Получаем id из параметров
	id := c.Params("id")

	// Преобразуем строку id в ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	// Получаем коллекцию из базы данных
	libraryCollection := database.GetCollection("book")

	// Удаляем запись
	_, err = libraryCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting book")
	}

	// Отправляем подтверждение об успешном удалении
	return c.SendString("Book deleted successfully")
}

func CreateBook(c *fiber.Ctx) error {
	createData := new(newBookDTO)

	if err := c.BodyParser(createData); err != nil {
		return err
	}

	// get the collection reference
	coll := database.GetCollection("book")

	result, err := coll.InsertOne(c.Context(), createData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create book",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}
