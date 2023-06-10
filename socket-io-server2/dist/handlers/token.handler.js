"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.decodeToken = void 0;
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const secret = (_a = process.env.JWT_SECRET) !== null && _a !== void 0 ? _a : "";
const decodeToken = (token) => {
    try {
        if (!token)
            return null;
        const payload = jsonwebtoken_1.default.verify(token, secret);
        console.log(payload);
        return payload;
    }
    catch (e) {
        console.log(e);
        return null;
    }
};
exports.decodeToken = decodeToken;
