package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
	"path/filepath"
	"strings"

	"github.com/Env-Co-Ltd/framinGo"
)

func InitApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// アプリケーション名をディレクトリ名から取得
	appName := filepath.Base(path)
	// 特殊文字を除去し、小文字に変換
	appName = strings.ToLower(strings.Replace(appName, "-", "", -1))

	//init framingo
	fra := &framinGo.FraminGo{}
	err = fra.New(path)
	if err != nil {
		log.Fatal(err)
	}

	fra.AppName = appName

	myMiddleware := &middleware.Middleware{
		App: fra,
	}

	myHandlers := &handlers.Handlers{
		App: fra,
	}

	app := &application{
		App:        fra,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}
	app.App.Routes = app.routes()

	models := data.New(app.App.DB.Pool)
	app.Models = models
	myHandlers.Models = app.Models
	app.Middleware.Models = *app.Models

	return app
}
