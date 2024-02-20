package main

import (
	//"log"

	//"github.com/gofiber/fiber/v3"
	m /*middleware*/ "github.com/welldn/restapi/middleware"
	"github.com/welldn/restapi/router"
)



func main() {
 //   app := fiber.New()

    // Perhaps name it with just 'r' for simplicity?
    router.Router()

    // Next route url
    m.Next()

    // Static files
    m.Static()


    //log.Fatal(app.Listen(":3000")) //-// think i can pass the app just in this but let me just test things first then i make it 'better'
}
