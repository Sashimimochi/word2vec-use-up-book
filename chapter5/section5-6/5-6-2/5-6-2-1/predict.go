package main

import (
	"log"
	"os"

	"github.com/ynqa/wego/pkg/embedding"
	"github.com/ynqa/wego/pkg/search"
)

func main() {
	input, err := os.Open("word2vec.model.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	embeds, err := embedding.Load(input)
	if err != nil {
		log.Fatal(err)
	}

	seacher, err := search.New(embeds...)
	if err != nil {
		log.Fatal(err)
	}

	word := "ニュース"
	n := 10
	neighbors, err := seacher.SearchInternal(word, n)
	if err != nil {
		log.Fatal(err)
	}

	neighbors.Describe()
}
