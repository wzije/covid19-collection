package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wzije/covid19-collection/handlers"
)

func Route() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", handlers.Home)
		v1.GET("/provinces", handlers.GetAllCasesInProvince)
		v1.GET("/provinces/latest", handlers.LatestCasesInProvince)
		v1.GET("/provinces/populate", handlers.CrawlInProvince)

		v1.GET("/temanggungs", handlers.GetAllCaseInTemanggung)
		v1.GET("/temanggungs/latest", handlers.GetLatestCasesInTemanggung)
		v1.GET("/temanggungs/populate", handlers.CrawlTmg)

		v1.GET("/populate", handlers.CrawlAll)
	}

	return r
}
