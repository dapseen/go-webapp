// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}

func getArticle(c *gin.Context) {
	//check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		//check if the articke exists
		if article, err := getArticleByID(articleID); err == nil {
			//call the HMTL method of the context to render a template
			c.HTML(
				http.StatusOK,

				"article.html",

				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		//if an invalid article ID is specified in the URL
		c.AbortWithStatus(http.StatusNotFound)
	}
}
