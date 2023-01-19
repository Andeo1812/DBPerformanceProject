package main

import (
	serviceForum "db-performance-project/internal/forum/service"
	servicePost "db-performance-project/internal/post/service"
	serviceSerivce "db-performance-project/internal/service/service"
	serviceThread "db-performance-project/internal/thread/service"
	serviceUser "db-performance-project/internal/user/service"
	serviceVote "db-performance-project/internal/vote/service"

	repoForum "db-performance-project/internal/forum/repository"
	repoPost "db-performance-project/internal/post/repository"
	repoService "db-performance-project/internal/service/repository"
	repoThread "db-performance-project/internal/thread/repository"
	repoUser "db-performance-project/internal/user/repository"
	repoVote "db-performance-project/internal/vote/repository"

	deliveryForum "db-performance-project/internal/forum/delivery/handlers"
	deliveryPost "db-performance-project/internal/post/delivery/handlers"
	deliveryService "db-performance-project/internal/service/delivery/handlers"
	deliveryThread "db-performance-project/internal/thread/delivery/handlers"
	deliveryUser "db-performance-project/internal/user/delivery/handlers"
	deliveryVote "db-performance-project/internal/vote/delivery/handlers"
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
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
	postgres := sqltools.NewPostgresRepository(&config.DatabaseParams)

	// Repository
	forumStorage := repoForum.NewForumPostgres(postgres)
	userStorage := repoUser.NewUserPostgres(postgres)
	postStorage := repoPost.NewPostPostgres(postgres)
	threadStorage := repoThread.NewThreadPostgres(postgres)
	voteStorage := repoVote.NewVotePostgres(postgres)
	serviceStorage := repoService.NewServicePostgres(postgres)

	// Service
	forumService := serviceForum.NewForumService(forumStorage, userStorage)
	userService := serviceUser.NewUserService(userStorage)
	postService := servicePost.NewPostService(postStorage)
	threadService := serviceThread.NewThreadService(threadStorage, forumStorage, userStorage, postStorage)
	voteService := serviceVote.NewVoteService(voteStorage, threadStorage)
	serivceService := serviceSerivce.NewService(serviceStorage)

	// Delivery Forum
	createForumHandler := deliveryForum.NewForumCreateHandler(forumService)
	createForumHandler.Configure(router, mw)

	getThreadsHandler := deliveryForum.NewForumGetThreadsHandler(forumService)
	getThreadsHandler.Configure(router, mw)

	getUsersHandler := deliveryForum.NewForumGetUsersHandler(forumService)
	getUsersHandler.Configure(router, mw)

	getDetailsForumHandler := deliveryForum.NewForumGetDetailsHandler(forumService)
	getDetailsForumHandler.Configure(router, mw)

	// Delivery Post
	updatePostHandler := deliveryPost.NewPostUpdateHandler(postService)
	updatePostHandler.Configure(router, mw)

	getDetailsPostHandler := deliveryPost.NewPostGetDetailsHandler(postService)
	getDetailsPostHandler.Configure(router, mw)

	// Delivery Service
	getStatusHandler := deliveryService.NewServiceGetStatusHandler(serivceService)
	getStatusHandler.Configure(router, mw)

	clearHandler := deliveryService.NewServiceClearHandler(serivceService)
	clearHandler.Configure(router, mw)

	// Delivery User
	getProfileHandler := deliveryUser.NewUserGetProfileHandler(userService)
	getProfileHandler.Configure(router, mw)

	updateProfileHandler := deliveryUser.NewUserUpdateProfileHandler(userService)
	updateProfileHandler.Configure(router, mw)

	createUserHandler := deliveryUser.NewUserCreateHandler(userService)
	createUserHandler.Configure(router, mw)

	// Delivery Thread
	createThreadPostsHandler := deliveryThread.NewThreadCreatePostsHandler(threadService)
	createThreadPostsHandler.Configure(router, mw)

	getPostsThreadHandler := deliveryThread.NewThreadGetPostsHandler(threadService)
	getPostsThreadHandler.Configure(router, mw)

	getDetailsThreadHandler := deliveryThread.NewThreadGetDetailsHandler(threadService)
	getDetailsThreadHandler.Configure(router, mw)

	createThreadHandler := deliveryThread.NewForumCreateThreadHandler(threadService)
	createThreadHandler.Configure(router, mw)

	updateThreadHandler := deliveryThread.NewThreadUpdateDetailsHandler(threadService)
	updateThreadHandler.Configure(router, mw)

	// Delivery Vote
	voteHandler := deliveryVote.NewThreadProfileHandler(voteService)
	voteHandler.Configure(router, mw)

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
