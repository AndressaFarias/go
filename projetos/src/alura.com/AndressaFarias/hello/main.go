package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello/:name", hello)

	app.Get("/ping",pong)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func pong(c *fiber.Ctx) error {
	err := c.SendString("pong")
	if err != nil {
		return err
	}

	return nil
}

func hello(c *fiber.Ctx) error {
	name := c.Params("name")

	message, err := getResponse(name)
	if err != nil {
		return err
	}

	err = c.SendString(message)
	if err != nil {
		return err
	}

	return nil
}

func getResponse(name string) (string, error)  {
	return fmt.Sprintf("Hellor World %s",name),nil
}
