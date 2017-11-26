package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var (
	db       *gorm.DB
	dbUser   = flag.String("dbUser", "", "db user")
	dbPasswd = flag.String("dbPasswd", "", "db password")
	dbURI    = flag.String("dbURI", "127.0.0.1:3306", "db uri")
	dbName   = flag.String("dbName", "", "db name")
)

func initDB() {
	var err error
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		*dbUser, *dbPasswd, *dbURI, *dbName,
	)
	db, err = gorm.Open("mysql", uri)
	if err != nil {
		logrus.Panicf("connect to db(%s) got error: %s", *dbURI, err)
	}
}

func redirectHandler(c *gin.Context) {
	short := c.Param("short")

	id := Decode(short)
	var link Links
	db.Where("id=?", id).First(&link)

	if link.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusFound, link.Original)
}

func indexHandler(c *gin.Context) {
	original := c.Query("original")

	if original == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var link Links

	db.Where("original=?", original).First(&link)

	if link.ID == 0 {
		link.Original = original
		db.Create(&link)
	}

	c.String(http.StatusOK, "%s/%s", c.Request.Host, Encode(link.ID))
}

func main() {
	// parse flag
	flag.Parse()
	// init database
	initDB()
	defer db.Close()

	// migration
	db.AutoMigrate(&Links{})

	r := gin.Default()
	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	r.GET("/", indexHandler)
	r.GET("/:short", redirectHandler)

	r.Run()
}
