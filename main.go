package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"BuffettFive/config"
	"BuffettFive/models"
)

func setupRouter() *gin.Engine {
	f, err := os.Create(config.Config.LogFile)
    if err != nil {
        log.Fatalf("Couldn't create logfile", err)
    }
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	r := gin.Default()
    r.Static("/static", "./templates/todo")
	r.LoadHTMLGlob("templates/**/*")

	r.POST("/todos/save", func(c *gin.Context) {
        content := c.PostForm("content")
        category := c.PostForm("category")

        if err := models.CreateTodo(content, category); err != nil {
            var todos []models.Todo
            models.Db.Find(&todos)

            c.HTML(http.StatusOK, "list.html", gin.H{
                "title": "BuffettFive",
                "todos": todos,
                "errorMessage": err.Error(),
            })
            return
        }
		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	r.POST("/todos/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")
		category := c.PostForm("category")

		todo, err := models.GetTodo(id)
        if err != nil {
            c.String(http.StatusBadRequest, "Not fount target id")
            return
        }

		todo.Content = content
		todo.Category = category
        if err := models.UpdateTodo(todo); err != nil  {
            c.HTML(http.StatusOK, "edit.html", gin.H{
                "title": "BuffetFive",
                "todo": todo,
                "errorMessage": err.Error(),
            })
            return
        }
		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	r.GET("/todos/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}

		todo, _ := models.GetTodo(id)
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": "BuffettFive",
			"todo":  todo,
		})
	})

	r.GET("/todos/destroy", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}

		if err := models.DeleteTodo(id); err != nil {
	        c.String(http.StatusBadRequest, "Delete error: %v", err)
        }
        c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	r.GET("/todos/list", func(c *gin.Context) {
		var todos []models.Todo
		models.Db.Find(&todos)

		c.HTML(http.StatusOK, "list.html", gin.H{
			"title": "BuffettFieve" ,
			"todos": todos,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8989")
}
