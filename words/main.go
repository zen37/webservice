package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type input struct {
	Text string `json:"text"`
}

type txt struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func main() {

	router := gin.Default()

	router.POST("/mostusedtenwords", postText)

	router.Run("localhost:8080")

}

func postText(c *gin.Context) {

	var newInput input

	var t txt

	if err := c.BindJSON(&newInput); err != nil {
		return
	}

	fmt.Println(newInput)

	m := count(newInput.Text)

	for k, v := range m {
		t.Word = k
		t.Count = v
	}

	fmt.Println("t is ", t)
	/*
		user := &User{Name: "Frank"}
	    b, err := json.Marshal(user) */

	b, err := json.Marshal(txt{Word: t.Word, Count: t.Count})
	//b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	c.JSON(http.StatusCreated, gin.H{
		"code":          http.StatusOK,
		"mostusedwords": m})

}

func count(input string) map[string]int {

	s := strings.Fields(input)
	fmt.Println(s)

	m := make(map[string]int, len(s))

	for _, v := range s {
		m[strings.ToLower(v)]++
	}

	fmt.Println(m)
	return m
}
