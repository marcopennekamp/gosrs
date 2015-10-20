package main

import (
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/julienschmidt/httprouter"

	"github.com/marcopennekamp/gosrs/database"
	"github.com/marcopennekamp/gosrs/handlers"
	"github.com/marcopennekamp/gosrs/core"
)

func createRouter(ctx *core.Context) *httprouter.Router {
	r := httprouter.New()

	// Front facing routes.
	r.GET("/", core.CreateHandle(ctx, handlers.Index))

	// Static file routing.
	r.ServeFiles("/static/*filepath", http.Dir(ctx.Conf.GetFilePath("static")))

	// Member routes.
	r.GET("/login", core.CreateHandle(ctx, handlers.ShowLogin))
	r.POST("/login", core.CreateHandle(ctx, handlers.Login))
	r.GET("/register", core.CreateHandle(ctx, handlers.ShowRegistration))
	r.POST("/register", core.CreateHandle(ctx, handlers.Register))

	return r
}

func main() {
	conf, err := core.ReadConfig("conf.yml")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.Open(conf)
	if err != nil {
		log.Fatalln(err)
	}

	tl, err := core.NewTemplateLoader(conf.GetFilePath("templates"))
	if err != nil {
		log.Fatalln(err)
	}

	ctx := &core.Context{Conf: conf, Db: db, Templates: tl}
	http.ListenAndServe(":8000", createRouter(ctx))
}