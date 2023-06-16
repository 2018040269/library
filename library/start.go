package library

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Start() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("library/**/*")
	r.GET("/book/list", booklListHandle)
	r.GET("/book/new", newBookhandle)
	r.POST("/book/new", createBookHandle)
	r.GET("/book/delete", deleteHandle)
	r.GET("/book/update", newHandle)
	r.POST("/book/update", updateHandle)
	r.Run(":3309")
}

func booklListHandle(c *gin.Context) {
	bookList, err := queryAlllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err.Error(),
			"code": 1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": bookList,
	})
	// c.HTML(http.StatusOK, "book/book_list.html", gin.H{
	// 	"code": 0,
	// 	"data": bookList,
	// })
}

func newBookhandle(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

func createBookHandle(c *gin.Context) {
	titleVal := c.PostForm("title")
	numberVal := c.PostForm("number")
	fmt.Printf("number:%x\n", []byte(numberVal))
	number, err := strconv.ParseFloat(numberVal, 64)
	if err != nil {
		fmt.Println("转换失败")
		fmt.Println(err)

		return
	}
	err = insertAlllBook(titleVal, number)
	if err != nil {
		fmt.Println("转换失败2")

		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func deleteHandle(c *gin.Context) {
	idVal := c.Query("id")
	id, err := strconv.ParseInt(idVal, 10, 64)
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"err":  err.Error(),
			"code": 1,
		})
	}
	err = deleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":  err.Error(),
			"code": 2,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func newHandle(c *gin.Context) {
	idVal := c.Query("id")
	id, err := strconv.ParseInt(idVal, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":  err.Error(),
			"code": 3,
		})
		return
	}
	book, err := querySingalBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err.Error(),
			"code": 4,
		})
		return
	}
	c.HTML(http.StatusOK, "book/updatebook.html", book)
}

func updateHandle(c *gin.Context) {
	titleVal := c.PostForm("title")
	numberVal := c.PostForm("number")
	idVal := c.PostForm("id")

	number, err := strconv.ParseFloat(numberVal, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	id, err := strconv.ParseInt(idVal, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}

	err = updateBook(titleVal, number, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":  err.Error(),
			"code": 5,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/book/list")
}
