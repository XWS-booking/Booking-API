import { RedisClientType, createClient } from "redis"


export type RedisConnection = RedisClientType<Record<string, never>, Record<string, never>>

export default (): RedisConnection => createClient({
    url: process.env.REDIS_URL
})