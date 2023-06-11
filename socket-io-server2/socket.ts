import { Server, Socket } from "socket.io";
import { Notification, onConnection } from "./handlers/socket.handler";
import { RedisConnection } from "./redis";

export interface ServerToClientEvents {
}
export interface ClientToServerEvents {
    notification: (message: Notification) => void
}
export interface InterServerEvents {}
export type SocketServer = Server<ServerToClientEvents, ClientToServerEvents, InterServerEvents, any>
export type Websocket = Socket<ServerToClientEvents, ClientToServerEvents, InterServerEvents, any>



export const initWebsockets = (
    io: SocketServer, 
    redis: RedisConnection
) => {
    io.on('connection', onConnection(redis))
}