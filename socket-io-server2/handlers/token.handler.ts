import jwt from "jsonwebtoken"

const secret = process.env.JWT_SECRET ?? ""

export type JwtPayload = {
    id: number
}


export const decodeToken = (token: string | undefined) => {
    try {
        if(!token) return null
        const payload = jwt.verify(token, secret)
        console.log(payload)
        return payload
    } catch(e) {
        console.log(e)
        return null
    }
}