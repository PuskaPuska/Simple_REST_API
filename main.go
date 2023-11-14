package main

import (
	"context"

	"github.com/PuskaPuska/Simple_REST_API/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//init App
	err := initApp()
	if err != nil {
		panic(err)
	}

	//defer close database
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err :=collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		return c.JSON(nDoc)
	})

	app.Listen(":3000")
}

func initApp() error {
	//setup env
	err := loadENV()
	if err != nil {
		return err
	}

	//setup database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
