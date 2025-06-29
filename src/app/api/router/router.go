package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/misalima/nano-link-backend/src/app/api/container"
	"github.com/misalima/nano-link-backend/src/app/api/handlers"
)

func NewRouter(c *container.Container) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodDelete,
		},
	}))

	urlHandler := handlers.NewURLHandler(c.URLService())
	tagHandler := handlers.NewTagHandler(c.TagService(), c.URLTagService())
	visitHandler := handlers.NewVisitHandler(c.URLService())

	setUpRoutes(e, urlHandler, tagHandler, visitHandler)

	return e
}

func setUpRoutes(e *echo.Echo, urlHandler *handlers.URLHandler, tagHandler *handlers.TagHandler, visitHandler *handlers.VisitHandler) {
	e.GET("/:short_id", urlHandler.RedirectToOriginalURL)

	api := e.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	urls := api.Group("/urls")
	urls.POST("", urlHandler.CreateShortURL)
	urls.GET("/:short_id", urlHandler.GetURLDetails)
	urls.DELETE("/:short_id", urlHandler.DeleteURL)

	urls.POST("/:short_id/tags", tagHandler.AddTagToURL)
	urls.DELETE("/:short_id/tags/:tag_id", tagHandler.RemoveTagFromURL)
	urls.GET("/:short_id/tags", tagHandler.GetTagsForURL)

	urls.GET("/:short_id/visits/count", visitHandler.GetVisitCount)
	urls.GET("/:short_id/visits", visitHandler.GetVisitHistory)
}
