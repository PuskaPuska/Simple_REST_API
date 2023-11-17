package handlers

import (
	"context"

	"github.com/PuskaPuska/Simple_REST_API/database"
	"github.com/PuskaPuska/Simple_REST_API/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibrarDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exist" bson:"books"`
}

func DeleteLibrary(c *fiber.Ctx) error {
	// Получаем ID из параметров запроса
	id := c.Params("id")

	// Преобразование ID из строки в ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Возвращаем ошибку, если ID невалиден
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	libraryCollection := database.GetCollection("libraries")

	// Удаление библиотеки
	_, err = libraryCollection.DeleteOne(c.Context(), bson.M{"_id": objID})
	if err != nil {
		// Обработка ошибок при удалении
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting library")
	}

	// Отправка подтверждения об успешном удалении
	return c.SendString("Library deleted successfully")
}

func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var libraries []models.Library

	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}
	return c.JSON(libraries)
}

func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(LibrarDTO)
	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	nLibrary.Empty = make([]string, 0)
	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}
