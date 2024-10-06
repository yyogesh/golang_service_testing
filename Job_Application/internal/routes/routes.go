package routes

import (
	"database/sql"
	"job_portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	{
		// AUTH ROUTES
		r.POST("/login", handlers.LoginHandler(db))
		r.POST("/register", handlers.RegisterHandler(db))

		// User routes // employer
		r.GET("/users/:id", handlers.GetUsersHandler(db))
		r.PUT("/users/:id", handlers.GetUsersHandler(db))
		r.POST("/users/:id/picture", handlers.GetUsersHandler(db))
		r.GET("/users/:id/jobs", handlers.GetUsersHandler(db))

		// JOB Routes
		r.POST("/jobs", handlers.GetUsersHandler(db))
		r.GET("/jobs", handlers.GetUsersHandler(db))
		r.GET("/jobs/:id", handlers.GetUsersHandler(db))
		r.PUT("/jobs/:id", handlers.GetUsersHandler(db))
		r.DELETE("/jobs/:id", handlers.GetUsersHandler(db))

		// Admin Routes
		r.GET("/admin/jobs", handlers.GetUsersHandler(db))
		r.GET("/admin/jobs/:id", handlers.GetUsersHandler(db))
		r.GET("/admin/users", handlers.GetUsersHandler(db))

	}
}
