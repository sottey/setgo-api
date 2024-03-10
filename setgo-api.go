package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"setgo-api/structs"
	"setgo-api/utilities"

	_ "github.com/go-sql-driver/mysql"
)

// .env variables
var setAppUrl string
var cacheFile string
var subCatFile string
var helpTemplate string
var favIcon string
var url string

var subCats []structs.Subcat

func main() {
	setAppUrl = utilities.GetConfigValue("SETAPP_URL")
	cacheFile = utilities.GetConfigValue("CACHE_FILE")
	subCatFile = utilities.GetConfigValue("SUBCAT_FILE")
	helpTemplate = utilities.GetConfigValue("HELP_TEMPLATE")
	url = utilities.GetConfigValue("SERVER_URL")
	favIcon = utilities.GetConfigValue("FAVICON")

	subCats, _ = utilities.LoadSubCats(subCatFile)

	router := gin.Default()
	router.LoadHTMLFiles(helpTemplate)

	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "helpTemplate", map[string]interface{}{}) })
	router.StaticFile("/favicon.ico", favIcon)
	router.GET("/apps", GetApps)
	router.GET("/apps/:id", GetApps)
	router.GET("/cats", GetAllCategories)
	router.GET("/cats/:cat", GetByCategory)
	router.GET("/subcats", GetAllSubCategories)
	router.GET("/search/:query", GetBySearchTerm)

	router.Run(url)
}

func GetApps(c *gin.Context) {
	apps, _ := utilities.GetData(setAppUrl, cacheFile)
	id := c.Param("id")

	if id == "" {
		c.IndentedJSON(http.StatusOK, apps)
		return
	} else {
		for _, a := range apps {
			reqID, _ := strconv.Atoi(id)
			if reqID == a.Id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "app id " + id + " was not found"})
		return
	}
}

func GetAllCategories(c *gin.Context) {
	var catList []string
	apps, _ := utilities.GetData(setAppUrl, cacheFile)

	for _, a := range apps {
		catList = append(catList, a.Categories...)
	}

	catList = utilities.RemoveDuplicate[string](catList)

	c.IndentedJSON(http.StatusOK, catList)
}

func GetAllSubCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, subCats)
}

func GetByCategory(c *gin.Context) {
	var appList []structs.App
	apps, _ := utilities.GetData(setAppUrl, cacheFile)
	cat := c.Param("cat")

	if cat == "" {
		c.IndentedJSON(http.StatusOK, apps)
		return
	} else {
		for _, a := range apps {
			for _, currCat := range a.Categories {
				if strings.EqualFold(cat, currCat) {
					appList = append(appList, a)
				}
			}
		}

		if len(appList) < 1 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "category " + cat + " was not found"})
		} else {
			c.IndentedJSON(http.StatusOK, appList)
			return
		}
	}
}

func GetBySearchTerm(c *gin.Context) {
	var appList []structs.App
	apps, _ := utilities.GetData(setAppUrl, cacheFile)
	query := strings.ToLower(c.Param("query"))

	if query == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no search term provided"})
		return
	} else {
		for _, a := range apps {
			appName := strings.ToLower(a.Name)
			appDesc := strings.ToLower(a.Description)

			if strings.Contains(appName, query) || strings.Contains(appDesc, query) {
				appList = append(appList, a)
			}
		}

		if len(appList) < 1 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no results found for query '" + c.Param("query") + "'"})
		} else {
			c.IndentedJSON(http.StatusOK, appList)
			return
		}
	}
}
