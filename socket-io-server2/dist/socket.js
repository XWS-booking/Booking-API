"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initWebsockets = void 0;
const socket_handler_1 = require("./handlers/socket.handler");
const initWebsockets = (io, redis) => {
    io.on('connection', (0, socket_handler_1.onConnection)(redis));
};
exports.initWebsockets = initWebsockets;
