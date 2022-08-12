const corpusFilePath = 'corpus.txt'
const w2v = require('word2vec');
w2v.word2vec(corpusFilePath, "vectors.bin", {size: 200, binary: 1}, () => {
  console.log("DONE");
});
