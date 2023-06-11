"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const redis_1 = require("redis");
exports.default = () => (0, redis_1.createClient)({
    url: process.env.REDIS_URL
});
