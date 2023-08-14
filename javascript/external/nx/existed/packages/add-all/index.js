const _ = require('lodash');
const {isValid} = require('json-check')

const arr_sum = function(numArr) {
    _.sum(numArr);
}

module.exports.arr_sum = arr_sum;

module.exports.arr_sum_from_object = function(obj) {
    const data = {
        sum_target: [1, 2, 3, 4, 5],
        foo: 123,
        bar: "abc",
    }

    const schema = {
        type: "object",
        properties: {
            sum_target: {type: "array"},
        },
        required:["sum_target"],
        additionalProperties: true,
    }

    if(isValid(data, schema)) {
        return arr_sum(data.sum_target);
    }

    throw new Error("not valid object");
}