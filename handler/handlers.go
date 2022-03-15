package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ivanauliaa/shortener-go/database"
	"github.com/ivanauliaa/shortener-go/model"
	"github.com/ivanauliaa/shortener-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Redirect(c *gin.Context) {
	collection := database.Connect()
	urlCode := c.Param("code")

	result := &model.URLDoc{}
	err := collection.FindOne(utils.GLOBAL_CONTEXT, bson.M{"urlCode": urlCode}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No URL with code: %s", urlCode)})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	log.Print(result)

	longURL := result.LongURL
	c.Redirect(http.StatusPermanentRedirect, longURL)
}

func Shorten(c *gin.Context) {
	collection := database.Connect()
	requestBody := &model.RequestBody{}
	if err := c.BindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := url.ParseRequestURI(requestBody.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlCode, err := shortid.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := &model.URLDoc{}
	err = collection.FindOne(utils.GLOBAL_CONTEXT, bson.M{"urlCode": urlCode}).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if result.URLCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Code in use: %s", urlCode)})
		return
	}

	docID := primitive.NewObjectID()
	shortURL := utils.BASE_URL + urlCode
	createdAt := time.Now()
	expires := createdAt.AddDate(0, 0, 5)

	newURLDoc := &model.URLDoc{
		ID:        docID,
		URLCode:   urlCode,
		LongURL:   requestBody.LongURL,
		ShortURL:  shortURL,
		CreatedAt: createdAt,
		Expires:   expires,
	}

	_, err = collection.InsertOne(utils.GLOBAL_CONTEXT, newURLDoc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Print("New URL document created")

	response := gin.H{
		"message": "Success create new short URL",
		"data": gin.H{
			"new_url": shortURL,
			"expires": expires.Format("2006-01-02 15:04:05"),
			"id":      docID,
		},
	}
	c.JSON(http.StatusOK, response)
}
