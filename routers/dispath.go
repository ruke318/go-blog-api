package routers

import "github.com/kataras/iris"
import "github.com/kataras/iris/core/router"
import "blog/controllers"
import "github.com/jinzhu/gorm"
import "github.com/iris-contrib/middleware/cors"
import "github.com/olivere/elastic"

func Dispath(db *gorm.DB, esClient *elastic.Client) (api *iris.Application) {
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
		controllers.SetOrm(db, esClient)
		app.PartyFunc("/user", func(users router.Party) {
			userCtr := controllers.UserController()
			users.Get("/", userCtr.GetAll)
			users.Get("/{id: uint}", userCtr.GetUserById)
			users.Post("/", userCtr.AddUser)
		})
		app.PartyFunc("/link", func(links router.Party) {
			linkCtr := controllers.LinkController()
			links.Get("/", linkCtr.GetAll)
		})
		app.PartyFunc("/posts", func(posts router.Party) {
			postsCtr := controllers.PostsController()
			posts.Get("/", postsCtr.GetList)
			posts.Post("/", postsCtr.AddPosts)
			posts.Get("/{id: uint}", postsCtr.GetDetail)
			posts.Get("/search", postsCtr.Search)
		})
		app.PartyFunc("/nav", func(nav router.Party) {
			navCtr := controllers.NavController()
			nav.Get("/", navCtr.GetNav)
			nav.Post("/", navCtr.AddNav)
		})
		app.PartyFunc("/replys", func(reply router.Party) {
			replysCtr := controllers.ReplysController()
			reply.Get("/", replysCtr.GetList)
		})
	}
	return
}
