package http

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func InitSwagger(router *gin.Engine, logger *slog.Logger, dir string) {
	router.GET("/tvplus-middleware/swagger", func(c *gin.Context) {
		bytes, err := os.ReadFile(filepath.Join(dir, "oas.json"))
		if err != nil {
			logger.Error("error reading swagger.json")
			c.String(http.StatusInternalServerError, "Failed to read swagger.json: %v", err)
			return
		}

		c.Data(http.StatusOK, "application/json", bytes)
	})

	router.GET("/swagger", func(c *gin.Context) {
		html := `
<!DOCTYPE html>
<html>
  <head>
    <title>Swagger UI</title>
    <link href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" rel="stylesheet" />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
    <script>
      SwaggerUIBundle({
        url: '/tvplus-middleware/swagger',
        dom_id: '#swagger-ui'
      });
    </script>
  </body>
</html>
`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

}
