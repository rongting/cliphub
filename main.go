package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
)

//go:embed templates/*
var htmlFS embed.FS

//go:embed static/*
var assetFS embed.FS

func main() {
	router := gin.Default()

	router.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(assetFS))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	router.SetHTMLTemplate(template.Must(template.New("").ParseFS(htmlFS, "templates/*")))

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/send", func(c *gin.Context) {
		message := c.PostForm("message")
		ok, errormsg := Set(MessageKey, message, MESSAGE)
		if ok {
			c.String(http.StatusOK, "OK")
		} else {
			c.String(http.StatusForbidden, errormsg)
		}
	})

	router.POST("/receive", func(c *gin.Context) {
		message := Get(MessageKey, MESSAGE)
		c.String(http.StatusOK, message)
	})

	router.POST("/upload", func(c *gin.Context) {
		err := os.RemoveAll(FilesDir)
		if err != nil {
			fmt.Println(err)
		}
		err = os.Mkdir(FilesDir, os.ModeDir)
		if err != nil {
			c.String(http.StatusInternalServerError, "Create directory failed")
			return
		}
		file, _ := c.FormFile("fileName")
		if err := c.SaveUploadedFile(file, FilesDir + "/" + file.Filename); err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "Save file failed")
			return
		}
		Set(FileKey, file.Filename, FILE)
		c.String(http.StatusOK, "OK")
	})

	router.GET("/download", func(c *gin.Context) {
		fileName := Get(FileKey, FILE)
		fmt.Println("filename: ", fileName)
		if fileName == "" {
			c.String(http.StatusForbidden, "No record found for token")
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(FilesDir + "/" + fileName)
	})

	router.Run(":9000")
}

