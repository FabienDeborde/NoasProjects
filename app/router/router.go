package router

func Init(app *App) {
	// setupRoutes(app)
	v1 := app.Group("/api/v1") // /api/v1
	projectRoutes(v1)
} 