package main

import (
	"flag"
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

	// コマンドライン引数からアプリケーション名を取得
	var appNameFlag string
	flag.StringVar(&appNameFlag, "name", "", "アプリケーション名")
	flag.Parse()

	// アプリケーション名の優先順位:
	// 1. コマンドライン引数
	// 2. 環境変数
	// 3. config/app.name
	// 4. ディレクトリ名
	appName := appNameFlag
	if appName == "" {
		appName = os.Getenv("APP_NAME")
	}
	if appName == "" {
		if nameBytes, err := os.ReadFile(filepath.Join(path, "config", "app.name")); err == nil {
			appName = strings.TrimSpace(string(nameBytes))
		}
	}
	if appName == "" {
		appName = filepath.Base(path)
	}

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
