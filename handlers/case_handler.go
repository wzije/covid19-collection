package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wzije/covid19-collection/services"
)

func CrawlAll(c *gin.Context) {
	services.CrawlAll()

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Generate data is successfully",
	})
}

func CrawlInProvince(c *gin.Context) {
	services.CrawlProvince()

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Generate data is successfully",
	})
}

func LatestCasesInProvince(c *gin.Context) {

	cases, err := services.GetLatestCasesProvince()

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "fetched",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "fetched",
		"message": "get latest cases done",
		"data":    cases,
	})
}

func GetAllCasesInProvince(c *gin.Context) {

	cases, err := services.GetAllCasesProvince()

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "fetched",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "fetched",
		"message": "get cases done",
		"data":    cases,
	})
}

//func CaseInfo(c *gin.Context) {
//
//	cases, err := services.GetCaseInfo()
//
//	if err != nil {
//		c.JSON(400, gin.H{
//			"code":    400,
//			"status":  "fetched",
//			"message": err.Error(),
//			"data":    nil,
//		})
//	}
//
//	c.JSON(200, gin.H{
//		"code":    200,
//		"status":  "fetched",
//		"message": "get case info done",
//		"data":    cases,
//	})
//}
//
//func CreateCaseInfo(c *gin.Context) {
//
//	_, err := services.StoreCaseInfo()
//
//	if err != nil {
//		c.JSON(400, gin.H{
//			"code":    400,
//			"status":  "Create case info failed.",
//			"message": err.Error(),
//			"data":    nil,
//		})
//	}
//
//	c.JSON(200, gin.H{
//		"code":    200,
//		"status":  "fetched",
//		"message": "create case info done",
//	})
//}

//---temanggung

func CrawlTmg(c *gin.Context) {
	services.CrawlTemanggung()

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Generate data temanggung is successfully",
	})
}

func GetLatestCasesInTemanggung(c *gin.Context) {
	cases, err := services.GetLatestCaseInTmg()

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "fetched",
		"message": "get latest cases in Temanggung done",
		"data":    cases,
	})
}

func GetAllCaseInTemanggung(c *gin.Context) {
	cases, err := services.GetAllCaseInTmg()

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "fetched",
		"message": "get cases in All temanggung done",
		"data":    cases,
	})
}
