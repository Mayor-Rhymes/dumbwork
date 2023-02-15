package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

   app := fiber.New()



   app.Get("/", func (c *fiber.Ctx) error {


	    return c.SendString("Hello World")
   })


   var port string
   port = os.Getenv("PORT")

   app.Listen(port)
    
}