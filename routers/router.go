// package routers

// import (
//     "rental_api/controllers"
//     "github.com/beego/beego/v2/server/web"
// )

// func init() {
//     // Property routes
//     ns := web.NewNamespace("/api/v1",
//         web.NSNamespace("/properties",
//             web.NSRouter("/", &controllers.PropertyController{}, "get:ListProperties"),
//             web.NSRouter("/:id", &controllers.PropertyController{}, "get:GetProperty;put:UpdateProperty;delete:DeleteProperty"),
//             web.NSRouter("/", &controllers.PropertyController{}, "post:CreateProperty"),
//             web.NSRouter("/:id/details", &controllers.PropertyController{}, "get:GetPropertyDetails"),
//         ),
//         web.NSNamespace("/locations",
//             web.NSRouter("/", &controllers.LocationController{}, "get:ListLocations;post:CreateLocation"),
//             web.NSRouter("/:id", &controllers.LocationController{}, "get:GetLocation;put:UpdateLocation;delete:DeleteLocation"),
//         ),
//     )
//     web.AddNamespace(ns)
// }











package routers

import (
    "property-listing/controllers"
    "github.com/beego/beego/v2/server/web"
)

func init() {
    // CRUD routes for properties
    web.Router("/properties", &controllers.BookingController{}, "post:Create")
    web.Router("/properties/:id", &controllers.BookingController{}, "get:Get")
    web.Router("/properties/:id", &controllers.BookingController{}, "put:Put")
    web.Router("/properties/:id", &controllers.BookingController{}, "delete:Delete")
    web.Router("/properties", &controllers.BookingController{}, "get:List")
}
// // @APIVersion 1.0.0
// // @Title beego Test API
// // @Description beego has a very cool tools to autogenerate documents for your API
// // @Contact astaxie@gmail.com
// // @TermsOfServiceUrl http://beego.me/
// // @License Apache 2.0
// // @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// package routers

// import (
// 	"rental_api/controllers"

// 	beego "github.com/beego/beego/v2/server/web"
// )

// func init() {
// 	ns := beego.NewNamespace("/v1",
// 		beego.NSNamespace("/object",
// 			beego.NSInclude(
// 				&controllers.ObjectController{},
// 			),
// 		),
// 		beego.NSNamespace("/user",
// 			beego.NSInclude(
// 				&controllers.UserController{},
// 			),
// 		),
// 	)
// 	beego.AddNamespace(ns)
// }
