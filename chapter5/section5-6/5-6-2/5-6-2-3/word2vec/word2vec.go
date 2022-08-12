package word2vec

import (
	"os"

	"github.com/ynqa/wego/pkg/embedding"
	"github.com/ynqa/wego/pkg/search"
)

type W2V struct {
	Model *search.Searcher
}

type Similars struct {
	Result search.Neighbors
}

func NewWord2Vec() (*W2V, error) {
	input, err := os.Open("model/word2vec.model.txt")
	if err != nil {
		return nil, err
	}
	defer input.Close()

	embeds, err := embedding.Load(input)
	if err != nil {
		return nil, err
	}

	seacher, err := search.New(embeds...)
	if err != nil {
		return nil, err
	}

	w2v := new(W2V)
	w2v.Model = seacher
	return w2v, nil
}

func (w2v W2V) MostSimilar(word string, n int) (*search.Neighbors, error) {
	neighbors, err := w2v.Model.SearchInternal(word, n)
	if err != nil {
		return nil, err
	}

	return &neighbors, nil
}
