import { DisconnectReason } from "socket.io";
import { RedisConnection } from "../redis";
import { SocketServer, Websocket } from "../socket";
import { JwtPayload, decodeToken } from "./token.handler";

export type Notification = {
    title: string,
    body: string,
    userId: number
}

export const onConnection = (redis: RedisConnection) => {
    return (socket: Websocket) => {
        const token = socket.handshake.headers.authorization
        //const token = socker.hand
        const tokenPayload = decodeToken(token) as JwtPayload
        console.log(tokenPayload)
        if(!tokenPayload) {
            socket.disconnect()
            return
        }
        redis.set(`${tokenPayload.id}`, socket.id)
        
        socket.on('disconnect', onDisconnect(redis, socket))
    }
}


export const onDisconnect = (redis: RedisConnection, socket: Websocket) => {
    return (reason: DisconnectReason, description?: any) => {
        console.log(reason, description)
        deleteSingleByUserId(redis, socket.id)
        .catch(console.log)
    }
}


const deleteSingleByUserId = async (redis: RedisConnection, socketId: string) => {
    const keys = await redis.keys("*")
    await Promise.all(
        keys.map(async (key: any) => {
            const value = await redis.get(key)
            if(value === socketId) {
                await redis.del(key)
                return
            }
        })
    )
}

export const sendNotificationToUser = async (redis:RedisConnection, io: SocketServer, notification: Notification) => {
    const connectionId = await redis.get(`${notification.userId}`)
    if(!connectionId) return
    io.to(connectionId).emit('notification', notification)
}
