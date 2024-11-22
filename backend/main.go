package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}

type Question struct {
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
}

// reads questions from the json file
func getQuestions() ([]Question, error) {
	content, err := os.ReadFile("./questions.json")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var questions []Question
	err = json.Unmarshal(content, &questions)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
		return nil, err
	}

	return questions, nil
}

func main() {
	// starting fiber server
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/questions", func(c *fiber.Ctx) error {
		questions, err := getQuestions()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error getting questions",
			})
		}
		return c.JSON(questions)
	})

	log.Fatal(app.Listen(":3000"))
}
