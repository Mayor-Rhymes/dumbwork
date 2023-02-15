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


   
   port := os.Getenv("PORT")
     

   app.Listen(`0.0.0.0:` + port)
    
}