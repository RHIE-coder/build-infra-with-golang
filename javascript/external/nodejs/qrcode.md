 - 간단한 사용법

```js
var QRCode = require('qrcode')
const app = require('express')();

app.get("/qr", (req, res) => {
    QRCode.toDataURL('qr code usage', function (err, url) {
        const onlyBase64 = url.replace("data:image/png;base64,",'');
        res.end(onlyBase64, 'base64');
    })
});


app.listen(3666, ()=> console.log("MY SERVER IS ON AIR"));
```