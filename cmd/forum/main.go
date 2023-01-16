package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"db-performanc-eproject/internal/pkg"
)

func main() {
	// Config
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/api/configs/debug.toml", "path to config file")

	flag.Parse()

	config := pkg.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	// Logger
	logger, closeResource := pkg.NewLogger(&config.Logger)
	defer func(closer func() error, log *logrus.Logger) {
		err = closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	// Middleware
	mw := pkg.NewHTTPMiddleware(logger)

	// Router
	router := mux.NewRouter()

	// Connections
	// postgres := sqltools.NewPostgresRepository(&config.DatabaseParams)

	// Forum repository
	// forumStorage := repoForum.NewForumDatabase(postgres)

	// Forum service
	// forumService := serviceForum.NewForumService(authStorage)

	// Set middleware
	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	logrus.Info(config.ServerHTTPMain.ServiceName + " starting server at " + config.ServerHTTPMain.BindAddr + " on protocol " + config.ServerHTTPMain.Protocol)

	// Server
	server := pkg.NewServerHTTP(logger)

	err = server.Launch(config, router)
	if err != nil {
		logrus.Fatal(err)
	}
}
