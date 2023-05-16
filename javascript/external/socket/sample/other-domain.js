const express = require("express")
const app = express();
const createServer = require("http").createServer;
const path = require("path");

const httpServer = createServer(app);


app.use('/public', express
    .static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
    res.sendFile(__dirname + '/other.html');
});

const port = 7777;
httpServer.listen(port, () => {
    console.log("Server Running...");
})