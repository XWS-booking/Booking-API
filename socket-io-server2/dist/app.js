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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
const dotenv_1 = __importDefault(require("dotenv"));
const http_1 = require("http");
const socket_io_1 = require("socket.io");
dotenv_1.default.config();
const express_1 = __importDefault(require("express"));
const path_1 = __importDefault(require("path"));
const socket_handler_1 = require("./handlers/socket.handler");
const redis_1 = __importDefault(require("./redis"));
const cors_1 = __importDefault(require("cors"));
//Constants
const staticDir = path_1.default.resolve(process.cwd(), 'static');
const PORT = (_a = process.env.PORT) !== null && _a !== void 0 ? _a : '8080';
const ioOptions = {
    cors: {
        allowedHeaders: ["*"],
        origin: ["http://localhost:3000"],
        methods: ["GET", "POST", "PUT", "DELETE"]
    }
};
const app = (0, express_1.default)();
app.use((0, cors_1.default)());
const server = (0, http_1.createServer)(app);
const io = new socket_io_1.Server(server, ioOptions);
//Connect redis
const pubClient = (0, redis_1.default)();
const subClient = (0, redis_1.default)();
pubClient.connect();
subClient.connect();
subClient.subscribe('notification', (payloadStr) => __awaiter(void 0, void 0, void 0, function* () {
    const payload = JSON.parse(payloadStr);
    console.log(payload);
    yield (0, socket_handler_1.sendNotificationToUser)(pubClient, io, payload);
}));
//Init websockets
io.on('connection', (0, socket_handler_1.onConnection)(pubClient));
process.on('beforeExit', () => {
    pubClient.flushAll();
    subClient.unsubscribe('notification');
});
server.listen(PORT, () => {
    console.log('listening on PORT', PORT);
});
