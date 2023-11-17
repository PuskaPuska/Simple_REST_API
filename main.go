package main

import (
	"os"

	"github.com/PuskaPuska/Simple_REST_API/database"
	"github.com/PuskaPuska/Simple_REST_API/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//init App
	err := initApp()
	if err != nil {
		panic(err)
	}

	//defer close database
	defer database.CloseMongoDB()

	app := generateApp()

	/* app.Post("/", func(c *fiber.Ctx) error {
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err :=collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		return c.JSON(nDoc)
	}) */

	//get the port from the env
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()

	//create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//create the library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Post("/book", handlers.CreateBook)
	libGroup.Delete("/:id", handlers.DeleteLibrary)
	libGroup.Delete("/book/:id", handlers.DeleteBook)
	return app
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
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
