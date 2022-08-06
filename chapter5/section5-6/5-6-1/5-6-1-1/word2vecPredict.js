const w2v = require("word2vec");

const modelPath = "w2v.model.txt";
const word = "フルーツ";

w2v.loadModel(modelPath, (_, model) => {
    console.log(`SIZE: ${model.size}`);
    console.log(`WORDS: ${model.words}`);
    console.log("most similar words:");
    console.log(model.mostSimilar(word, 20));
});

const pair = ["スポーツ", "サッカー"];
const number_neighbors = 10;
w2v.loadModel(modelPath, (_, model) => {
    console.log(`analogy of ${word}, ex.${pair}`);
    console.log(model.analogy(word, pair, number_neighbors));
});
