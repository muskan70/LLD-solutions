package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type book1 struct {
	BookID   int     `json:"id"`
	Name     string  `json:"name"`
	Writer   string  `json:"author"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"` //tells about total books initially
	Sold     int     `json:"sold"`
	InStock  int     `json:"in_stock"`
	Status   string  `json:"status"`
}

var db *sql.DB

func main() {
	router := gin.Default()

	router.GET("/allbooks", getallbooks)
	router.GET("/allbooks/:new_id", getbookbyid)
	router.DELETE("/deletebook/:new_id", deletebookbyid)
	router.POST("/addbook", addnewbook)
	router.PUT("/update", updatedetails)
	str := "root:Ayushdd@123@tcp(127.0.0.1:3306)/library"
	var err error
	db, err = sql.Open("mysql", str)
	if err != nil {
		fmt.Println("Not able to connect")
		return
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Not able to connect")
		return
	}

	fmt.Println("connected")
	router.Run("localhost:8080")
	defer db.Close()
}
func getallbooks(c *gin.Context) {
	var book2 []book1
	rows, err := db.Query("SELECT ID, book_name, author, price, quantity, sold, in_stock, status FROM book")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	for rows.Next() {
		var bb book1
		if err := rows.Scan(&bb.BookID, &bb.Name, &bb.Writer, &bb.Price, &bb.Quantity, &bb.Sold, &bb.InStock, &bb.Status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "false",
				"message": "bad request",
				"error":   err.Error(),
			})
			return
		}
		book2 = append(book2, bb)
	}

	c.JSON(http.StatusOK, book2)
	defer rows.Close()
}
func getbookbyid(c *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(c.Param("new_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	rows, err := db.Query("SELECT ID, book_name, author, price, quantity, sold, in_stock, status FROM book")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	for rows.Next() {
		var bb book1
		if err := rows.Scan(&bb.BookID, &bb.Name, &bb.Writer, &bb.Price, &bb.Quantity, &bb.Sold, &bb.InStock, &bb.Status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "false",
				"message": "bad request",
				"error":   err.Error(),
			})
			return
		}
		if bb.BookID == id {
			c.JSON(http.StatusOK, bb)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "false",
		"message": "book not present",
		"error":   "please enter correct book id",
	})

	fmt.Println(id)
	defer rows.Close()
}
func deletebookbyid(c *gin.Context) {
	var err error
	id := (c.Param("new_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	row, err := db.Prepare("DELETE FROM book WHERE id=?")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	_, err = row.Exec((id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "false",
			"message": "Not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "true",
		"message": "Succesfully Deleted",
	})
	defer row.Close()
}

func addnewbook(c *gin.Context) {
	var req book1
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	row, err := db.Prepare("INSERT INTO book(id, book_name, author, price, quantity, sold, in_stock, status) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	_, err2 := row.Exec(req.BookID, req.Name, req.Writer, req.Price, req.Quantity, 0, req.Quantity, "A")
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "false",
			"message": "Not found",
			"error":   err2.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "true",
		"message": "Succesfully Added",
	})

	defer row.Close()
}
func updatedetails(c *gin.Context) {
	var req book1
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})

		return
	}
	row, err := db.Prepare("update book set id = ?, book_name = ?, author = ?, price = ?, quantity=?, sold=?, in_stock=?, status=? where id = ?")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}
	row.Exec(req.BookID, req.Name, req.Writer, req.Price, req.Quantity, req.Sold, req.InStock, req.Status, req.BookID)
	c.JSON(http.StatusOK, gin.H{
		"status":  "true",
		"message": "Updated Succesfully",
	})
	defer row.Close()
}
