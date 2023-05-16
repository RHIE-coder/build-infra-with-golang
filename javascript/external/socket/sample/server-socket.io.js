const express = require("express")
const app = express();
const createServer = require("http").createServer;
const { Server } = require("socket.io");
const path = require("path");

const httpServer = createServer(app);
const io = new Server(httpServer, {
    cors: {
        origin:"http://localhost:7777"
    }
})

app.use('/public', express
    .static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
    res.sendFile(__dirname + '/index.html');
});

app.get('/signal/:sock_id', (req, res) => {
    const sockID = req.params.sock_id;
    if(sockID === "all"){
        io.emit('msg', `이용해주셔서 감사합니다.`);
        return res.send(`all`);
    }else{
        io.to(sockID).emit(`msg`, `이용해주셔서 감사합니다, ${sockID}님!`)
        return res.send(`${sockID}`);
    }
    
})

io.on("connection", function (socket) {
    console.log(socket)
    const user = `${socket.id}`
    console.log(socket.id, 'connected');

    socket.on("msg", function (data) {
        console.log(`[${user}]:${data}`);
        io.emit('msg', `[${socket.id}]:${data}`); // This will emit the event to all connected sockets
    })

    socket.on('disconnect', (socket) => {
        io.emit('msg', `<${user}>님의 연결이 끊겼습니다.`);
        console.log('user disconnected');
    });

    socket.broadcast.emit("msg", `<${user}>님이 입장하셨습니다.`);
})


const port = 4444;
httpServer.listen(port, () => {
    console.log("Server Running...");
})

/* import express from 'express';
const app = express();
import { createServer } from "http";
import { Server  } from "socket.io";
import path from "path";
import { fileURLToPath } from 'url';
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
console.log(__filename);
console.log(__dirname);
console.log(import.meta.url);


const httpServer = createServer(app);
const io = new Server(httpServer)
const port = 4444;

app.use('/public', express
    .static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
    res.sendFile(__dirname + '/index.html');
});

io.on("connection", function (socket) {
    const user = socket.id
    console.log(socket.id, 'connected');

    socket.on("msg", function (data) {
        console.log(`[${user}]:${data}`);
        io.emit('msg', `[${socket.id}]:${data}`); // This will emit the event to all connected sockets
    })

    socket.on('disconnect', (socket) => {
        io.emit('msg', `<${user}>님의 연결이 끊겼습니다.`);
        console.log('user disconnected');
    });

    socket.broadcast.emit("msg", `<${user}>님이 입장하셨습니다.`);
})

httpServer.listen(port, () => {
    console.log("Server Running...");
}) */