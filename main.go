package main

import (
	"entitlement/entitlementsvc"
	"net/http"

	"github.com/gin-gonic/gin"
)

var entitlements = []entitlementsvc.Entitlement{
	{"E01", []entitlementsvc.Feature{
		{"F1", "Feature 1", "Rw"},
		{"F1", "Feature 2", "R"},
	}},
}

func main() {
	router := gin.Default()
	router.GET("/entitlements", getEntitlements)
	router.Run("localhost:8081")
}

func getEntitlements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entitlements)
}
