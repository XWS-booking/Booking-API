"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.sendNotificationToUser = exports.onDisconnect = exports.onConnection = void 0;
const token_handler_1 = require("./token.handler");
const onConnection = (redis) => {
    return (socket) => {
        const token = socket.handshake.headers.authorization;
        //const token = socker.hand
        const tokenPayload = (0, token_handler_1.decodeToken)(token);
        console.log(tokenPayload);
        if (!tokenPayload) {
            socket.disconnect();
            return;
        }
        redis.set(`${tokenPayload.id}`, socket.id);
        socket.on('disconnect', (0, exports.onDisconnect)(redis, socket));
    };
};
exports.onConnection = onConnection;
const onDisconnect = (redis, socket) => {
    return (reason, description) => {
        console.log(reason, description);
        deleteSingleByUserId(redis, socket.id)
            .catch(console.log);
    };
};
exports.onDisconnect = onDisconnect;
const deleteSingleByUserId = (redis, socketId) => __awaiter(void 0, void 0, void 0, function* () {
    const keys = yield redis.keys("*");
    yield Promise.all(keys.map((key) => __awaiter(void 0, void 0, void 0, function* () {
        const value = yield redis.get(key);
        if (value === socketId) {
            yield redis.del(key);
            return;
        }
    })));
});
const sendNotificationToUser = (redis, io, notification) => __awaiter(void 0, void 0, void 0, function* () {
    const connectionId = yield redis.get(`${notification.userId}`);
    if (!connectionId)
        return;
    io.to(connectionId).emit('notification', notification);
});
exports.sendNotificationToUser = sendNotificationToUser;
