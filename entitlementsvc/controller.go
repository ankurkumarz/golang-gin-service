package entitlementsvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/entitlements", getEntitlements)
	return router
}

func getEntitlements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entitlements)
}
