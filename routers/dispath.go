package routers

import "github.com/kataras/iris"
import "github.com/kataras/iris/core/router"
import "blog/controllers"
import "github.com/jinzhu/gorm"
import "github.com/iris-contrib/middleware/cors"

func Dispath(db *gorm.DB) (api *iris.Application) {
	api = iris.New()
	api.Get("/", func(ctx iris.Context) {
		ctx.WriteString("index")
	})
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
	})
	app := api.Party("/", crs).AllowMethods(iris.MethodOptions) 
	{	
		app.PartyFunc("/user", func(users router.Party) {
			userCtr := controllers.UserController()
			userCtr.SetOrm(db)
			users.Get("/", userCtr.GetAll)
			users.Get("/{id: uint}", userCtr.GetUserById)
			users.Post("/", userCtr.AddUser)
		})
		app.PartyFunc("/link", func(links router.Party) {
			linkCtr := controllers.LinkController()
			linkCtr.SetOrm(db)
			links.Get("/", linkCtr.GetAll)
		})
	}
	return
}
