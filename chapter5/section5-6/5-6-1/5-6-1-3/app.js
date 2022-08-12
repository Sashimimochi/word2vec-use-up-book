const express = require('express');
const app = express();
const port = process.env.PORT || 3000

const bodyParser = require('body-parser');

app.set('view engine', 'ejs');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const w2v = require('word2vec');
const modelPath = 'w2v.model.txt';
var w2v_model;

w2v.loadModel(modelPath, (error, model) => {
    console.log('Successfully loaded word2vec model!')
    console.log(`SIZE: ${model.size}`);
    console.log(`WORDS: ${model.words}`);
    w2v_model = model;
});    


app.get('/', (req, res) => {
    const data = {items: []}
    res.render("./index.ejs", data)
})

app.post('/', (req, res) => {
    const word = req.body.word;
    const data = {items: w2v_model.mostSimilar(word, 20) ?? [{word: `「${word}」はボキャブラリーにありませんでした`, dist: 0.0}]};
    res.render("./index.ejs", data);
})

app.listen(port, () => {
  console.log(`listening on *:${port}`);
})
