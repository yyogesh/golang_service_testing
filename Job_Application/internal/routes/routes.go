package routes

import (
	"database/sql"
	"job_portal/internal/auth"
	"job_portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	{
		// AUTH ROUTES
		r.POST("/login", handlers.LoginHandler(db))
		r.POST("/register", handlers.RegisterHandler(db))
		r.GET("/jobs", handlers.GetAllJobsHandler(db))

		// User routes // employer
		authenticated := r.Group("/")
		authenticated.Use(auth.AuthMiddleware())
		authenticated.GET("/users/:id", handlers.GetUsersHandler(db))
		authenticated.PUT("/users/:id", handlers.UpdateUserProfileHandler(db))
		authenticated.POST("/users/:id/picture", handlers.UpdateUserProfilePcitureHandler(db))
		authenticated.GET("/users/:id/jobs", handlers.GetUsersHandler(db))

		// JOB Routes
		// r.GET("/jobs", handlers.GetUsersHandler(db))
		authenticated.GET("/jobsByUser", handlers.GetUsersHandler(db))
		authenticated.POST("/jobs", handlers.CreateJobHandler(db))
		authenticated.PUT("/jobs/:id", handlers.GetUsersHandler(db))
		authenticated.DELETE("/jobs/:id", handlers.GetUsersHandler(db))

		// Admin Routes
		authenticated.GET("/admin/jobs", handlers.GetUsersHandler(db))
		authenticated.GET("/admin/jobs/:id", handlers.GetUsersHandler(db))
		authenticated.GET("/admin/users", handlers.GetUsersHandler(db))

	}
}
