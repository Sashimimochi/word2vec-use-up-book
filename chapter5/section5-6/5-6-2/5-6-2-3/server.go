package main

import (
	"net/http"

	"github.com/labstack/echo"
	"app/5-6-2-3/word2vec"
)

func main() {
	e := echo.New()
	e.GET("/mostsimilar", show)

	e.Logger.Fatal(e.Start(":5000"))
}

func show(c echo.Context) error {
	word := c.QueryParam("word")
	w2v, err := word2vec.NewWord2Vec()
	if err != nil {
		return err
	}

	res, err := w2v.MostSimilar(word, 10)
	if err != nil {
		return err
	}

	r := &word2vec.Similars{*res}

	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}
