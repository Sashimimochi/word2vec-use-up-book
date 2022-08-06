package main

import (
	"os"
	"log"

	"github.com/ynqa/wego/pkg/model/modelutil/vector"
	"github.com/ynqa/wego/pkg/model/word2vec"
)

func main() {
	model, err := word2vec.New(
		word2vec.Dim(200),
		word2vec.Window(5),
		word2vec.Model(word2vec.Cbow),
		word2vec.Optimizer(word2vec.NegativeSampling),
		word2vec.NegativeSampleSize(5),
		word2vec.Verbose(),
	)
	if err != nil {
		log.Fatal(err)
	}

	input, _ := os.Open("corpus.txt")
	defer input.Close()

	if err = model.Train(input); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("word2vec.model.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	model.Save(f, vector.Agg)
}
