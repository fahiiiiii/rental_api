// routers/router.go
package routers

import (
    "github.com/beego/beego/v2/server/web"
    "your-project/controllers"
)

func init() {
    // Create a new namespace for API versioning
    ns := web.NewNamespace("/v1",
        // Standard CRUD Routes for Rental Properties
        web.NSRouter("/properties", &controllers.BookingController{}, "post:Create"),
        web.NSRouter("/properties/:id", &controllers.BookingController{}, "get:Get"),
        web.NSRouter("/properties/:id", &controllers.BookingController{}, "put:Put"),
        web.NSRouter("/properties/:id", &controllers.BookingController{}, "delete:Delete"),
        web.NSRouter("/properties", &controllers.BookingController{}, "get:List"),

        // Specific Rental Property Routes
        web.NSRouter("/rental-properties", &controllers.BookingController{}, "post:CreateRentalProperty"),
        web.NSRouter("/rental-properties/:id", &controllers.BookingController{}, "get:GetRentalProperty"),
        web.NSRouter("/rental-properties", &controllers.BookingController{}, "get:ListRentalProperties"),

        // Additional Utility Routes
        web.NSRouter("/cities", &controllers.BookingController{}, "get:FetchCities"),
        web.NSRouter("/properties/city/:cityName", &controllers.BookingController{}, "get:FetchPropertiesForCity"),
    )

    // Add the namespace to the web application
    web.AddNamespace(ns)
}