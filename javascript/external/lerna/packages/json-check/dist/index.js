"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isValid = void 0;
var ajv_1 = require("ajv");
var ajv = new ajv_1.default();
function isValid(data, schema) {
    var valid = ajv.validate(schema, data);
    return valid;
}
exports.isValid = isValid;
