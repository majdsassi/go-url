package controllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/majdsassi/go-url/actions"
	"github.com/majdsassi/go-url/initialziers"
	"github.com/majdsassi/go-url/models"
)

func init() {
	initialziers.LoadEnvVars()
	initialziers.Connect2DB()
}
func CreateNewURL(c *gin.Context) {
	url := c.PostForm("url")
	if url =="" {
		c.HTML(http.StatusOK,"error.tmpl",gin.H{"errorMessage": "invalid url" ,})
		return
	}
	re := regexp.MustCompile(`https?:\/\/(www\.)?[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)+([\/\w\.-]*)*\/?`)
	if re.MatchString(url) == false {
		c.HTML(http.StatusOK,"error.tmpl",gin.H{"errorMessage": "invalid url" ,})
		return
	}
	id := actions.RandStringBytes(5)
	record := models.Url{ID: id, RidrectTo: url, Visit: 0}
	result := initialziers.DB.Create(&record)
	if result.Error != nil {
		c.HTML(http.StatusOK,"error.tmpl",gin.H{"errorMessage": result.Error.Error(),})
		return

	}
	c.HTML(http.StatusOK,"result.tmpl",gin.H{"url": c.Request.Host+"/"+id ,})

}
func GetUrl(c *gin.Context) {
	recordID := c.Param("id")
	if len(recordID)!= 5 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid ID " + recordID,
		})
		return
	}
	var url models.Url
	result := initialziers.DB.Where("id = ?", recordID).First(&url)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	url.Visit += 1 
	initialziers.DB.Save(&url)
	c.Redirect(301, url.RidrectTo)
}
