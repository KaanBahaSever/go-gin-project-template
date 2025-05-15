package routes

import (
	"log"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"draw/internal/config"
	"draw/internal/controllers"
	"draw/internal/database"
	"draw/internal/middleware"
	"draw/internal/repository"
	"draw/internal/services"
)

func SetupRoutes() *gin.Engine {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize database connection
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize router
	r := gin.Default()

	// Initialize session store
	store := cookie.NewStore([]byte(cfg.Session.Secret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   true, // Set to true if you want cookies to be sent only over HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("mysession", store))

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize controllers
	homeController := controllers.NewHomeController()
	authController := controllers.NewAuthController(authService)

	// Static files
	r.Static("/static", "./static")

	// Public routes
	r.GET("/login", authController.LoginPage)
	r.POST("/login", authController.Login)
	r.GET("/register", authController.RegisterPage)
	r.POST("/register", authController.Register)
	r.GET("/logout", authController.Logout)

	// Global middleware for adding user to context
	r.Use(middleware.AddUserToContext())

	// Protected routes group
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/", homeController.Index)

	return r
}
