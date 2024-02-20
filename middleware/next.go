package middleware 

import (
	"log"

	"github.com/gofiber/fiber/v3"
)
//When Next is called, it executes the next method in the stack that matches the current route.
func Next() {
    app := fiber.New()

    app.Static("/", "./public")
    // => http://localhost:3000/js/script.js
    // => http://localhost:3000/css/style.css

    app.Static("/prefix", "./public")
    // => http://localhost:3000/prefix/js/script.js
    // => http://localhost:3000/prefix/css/style.css

    app.Static("*", "./public/index.html")
    // => http://localhost:3000/any/path/shows/index/html

    log.Fatal(app.Listen(":3000"))
}

