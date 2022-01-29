package main

import (
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

var punctuation = [...]string{",", ".", "?", "!", ";"}

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

	c.JSON(http.StatusCreated, gin.H{
		"code":          http.StatusOK,
		"mostusedwords": m})

}

func count(input string) map[string]int {

	s := strings.Fields(input)
	fmt.Println(s)

	m := make(map[string]int, len(s))

	for _, v := range s {

		for _, p := range punctuation {
			v = strings.Replace(v, p, "", -1)
		}

		m[strings.ToLower(v)]++
	}

	fmt.Println(m)
	return m
}
