import dotenv from "dotenv";
import { createServer } from 'http';
import { Server } from 'socket.io';
dotenv.config()

import express from "express";
import path from "path";
import { onConnection, sendNotificationToUser } from "./handlers/socket.handler";
import redisClient, { RedisConnection } from "./redis";
import cors from 'cors'

//Constants
const staticDir = path.resolve(process.cwd(), 'static')
const PORT: string = process.env.PORT ?? '8080';
const ioOptions: any = {
    cors: {
        allowedHeaders: ["*"],
        origin: ["http://localhost:3001"],
        methods: ["GET", "POST", "PUT", "DELETE"]
    }
}


const app = express()
app.use(cors());
const server = createServer(app)
const io = new Server(server, ioOptions);

//Connect redis
const pubClient: RedisConnection = redisClient();
const subClient: RedisConnection = redisClient();
pubClient.connect()
subClient.connect()
    

subClient.subscribe('notification', async(payloadStr: string) => {
    const payload = JSON.parse(payloadStr)
    console.log(payload)
    await sendNotificationToUser(pubClient, io, payload)
})

//Init websockets
io.on('connection', onConnection(pubClient))

process.on('beforeExit', () => {
    pubClient.flushAll()
    subClient.unsubscribe('notification')
})

server.listen(PORT, () => {
    console.log('listening on PORT', PORT)
})








