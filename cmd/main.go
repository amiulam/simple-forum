package main

import (
	"log"

	"github.com/amiulam/simple-forum/internal/configs"
	"github.com/amiulam/simple-forum/internal/handlers/memberships"
	"github.com/amiulam/simple-forum/internal/handlers/posts"
	membershipRepo "github.com/amiulam/simple-forum/internal/repository/memberships"
	postRepo "github.com/amiulam/simple-forum/internal/repository/posts"
	membershipSvc "github.com/amiulam/simple-forum/internal/services/memberships"
	postSvc "github.com/amiulam/simple-forum/internal/services/posts"
	"github.com/amiulam/simple-forum/pkg/internalsql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisialisasi config", err)
	}

	cfg = configs.Get()
	// log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)

	if err != nil {
		log.Fatal("Gagal inisialisasi database", err)
	}

	// Middleware
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Repositories
	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	// Services
	membershipService := membershipSvc.NewService(membershipRepo, cfg)
	postService := postSvc.NewService(postRepo, cfg)

	// Handler
	membershipHandler := memberships.NewHandler(r, membershipService)
	postHandler := posts.NewHandler(r, postService)

	// Register Routes
	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
