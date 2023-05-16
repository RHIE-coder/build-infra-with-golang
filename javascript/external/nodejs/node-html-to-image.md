```js
const fs = require('fs');

// const file = fs.readFileSync("./SCDream5.otf")
// const base64ImgString = 'data:font/otf;base64,' + Buffer.from(file).toString('base64');
// fs.writeFileSync('mytxt.txt', base64ImgString);

const bufHTML = fs.createReadStream("./index.html");
let htmlString = bufHTML.toString('utf-8');
const name = "안녕이"
const date = "'22.12.11"
const price = "1000"
htmlString = htmlString.replace('{{{ NAME }}}', name).replace('{{{ DATE }}}', date).replace('{{{ PRICE }}}', price)

const HTMLToIMG = require('node-html-to-image');

(async()=> {
    const [one, two] =  await Promise.all([
        HTMLToIMG({
            html: htmlString,
        }),
        fs.promises.readFile('./SCDream5.otf'),
    ])
    console.log(one.length, two.length);
})()
```
 - 다른 방식

```js
const fs = require('fs');

// const file = fs.readFileSync("./SCDream5.otf")

// const base64ImgString = 'data:font/otf;base64,' + Buffer.from(file).toString('base64');

// fs.writeFileSync('mytxt.txt', base64ImgString);
const bufHTML = fs.createReadStream("./index.html");
let htmlString = bufHTML.toString('utf-8');
const name = "안녕이"
const date = "'22.12.11"
const price = "1000 NEWS"
htmlString = htmlString.replace('{{{ NAME }}}', name).replace('{{{ DATE }}}', date).replace('{{{ PRICE }}}', price)

const HTMLToIMG = require('node-html-to-image');

(async()=> {
    const [one, two] =  await Promise.all([HTMLToIMG({
        html: htmlString,
    }),
    // console.log(result.toString('base64'));
     fs.promises.readFile('./SCDream5.otf'),
    ])
    console.log(one.length, two.length);
    // console.log(font);
    // const img = fs.createWriteStream("example.jpg", result);
    // bufHTML.pipe(img);
})()

```