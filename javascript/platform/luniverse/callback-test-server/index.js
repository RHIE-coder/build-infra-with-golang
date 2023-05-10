const protocol = "http"
const port = 3333

const express = require('express');
const app = express();

app.use(express.urlencoded({extended: true}));
app.use(express.json());

const morgan = require('morgan');
app.use(
    morgan(':remote-addr - :remote-user [:date] ":method :url HTTP/:http-version" :status :res[content-length] - :response-time ms')
);

const count = {
    total: 0,
    mileage: 0,
    point: 0,
    nftcreate: 0,
    nftdeprecate: 0,
    nftmint: 0,
    nftburn: 0,
};

const router = express.Router()

router.post("/mileage", (req, res) => {
    count.total += 1
    count.mileage += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.mileage);
    res.send({msg: "success"})
})

router.post("/point", (req, res) => {
    count.total += 1
    count.point += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.point);
    res.send({msg: "success"})
})

router.post("/nft/create", (req, res) => {
    count.total += 1
    count.nftcreate += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.nftcreate);
    res.send({msg: "success"})
})

router.post("/nft/deprecate", (req, res) => {
    count.total += 1
    count.nftdeprecate += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.nftdeprecate);
    res.send({msg: "success"})
})

router.post("/nft/mint", (req, res) => {
    count.total += 1
    count.nftmint += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.nftmint);
    res.send({msg: "success"})
})

router.post("/nft/burn", (req, res) => {
    count.total += 1
    count.nftburn += 1
    console.log(req.headers)
    console.log(req.body);
    console.log(count.total);
    console.log(count.nftburn);
    res.send({msg: "success"})
})

app.use("/callback", router)

app.get("/health", (req, res) => {
    res.send({msg: "success"})
})

require(protocol)
    .createServer(app)
    .listen(port, ()=> {
        console.log(`Server Listening...`);
    })
