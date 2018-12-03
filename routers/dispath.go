package routers

import "github.com/kataras/iris"
import "github.com/kataras/iris/core/router"
import "blog/controllers"
import "github.com/jinzhu/gorm"

func Dispath(db *gorm.DB) (app *iris.Application) {
	app = iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("index")
	})
	app.PartyFunc("/user", func(users router.Party) {
		userCtr := controllers.UserController()
		userCtr.SetOrm(db)
		users.Get("/", userCtr.GetAll)
	})
	app.PartyFunc("/link", func(links router.Party) {
		linkCtr := controllers.LinkController()
		linkCtr.SetOrm(db)
		links.Get("/", linkCtr.GetAll)
	})
	return
}
