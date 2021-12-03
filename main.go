package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func f() {
	fmt.Println("{\"log\": \"123456\"}")
	panic("Hi, this is a panic")
}

func main() {
	fmt.Println("vim-go")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		f()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		time.Sleep(time.Second * 5)
		panic("潘妮可")
	}()

	s := &http.Server{
		Addr:           ":9090",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
