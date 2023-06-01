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
	r.LoadHTMLGlob("template/**/*")
	r.GET("/book/list", booklListHandle)
	r.GET("/book/new", newBookhandle)
	r.POST("/book/new", createBookHandle)
	r.GET("/book/delete", deleteHandle)
	r.GET("/book/update", newHandle)
	r.POST("/book/update", updateHandle)
	r.Run()
}

func booklListHandle(c *gin.Context) {
	bookList, err := quetyAlllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err.Error(),
			"code": 1,
		})
		return
	}
	c.HTML(http.StatusOK, "book/book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func newBookhandle(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

func createBookHandle(c *gin.Context) {
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	err = insertAlllBook(titleVal, price)
	if err != nil {
		c.String(http.StatusOK, "插入数据失败")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func insertAlllBook(titleVal string, price float64) {
	panic("unimplemented")
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
			"code": 1,
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
			"code": 1,
		})
		return
	}
	book, err := querySingalBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err.Error(),
			"code": 1,
		})
		return
	}
	c.HTML(http.StatusOK, "book/updatebook.html", book)
}

func updateHandle(c *gin.Context) {
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	idVal := c.PostForm("id")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	id, err := strconv.ParseInt(idVal, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	err = updateBook(titleVal, price, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err":  err.Error(),
			"code": 1,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}
