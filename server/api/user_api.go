package api

import "github.com/gin-gonic/gin"

func createUser(c *gin.Context) {

}


// URIのIDを取得する。
func getUserID(c *gin.Context) string {
	return c.Param("id")
}