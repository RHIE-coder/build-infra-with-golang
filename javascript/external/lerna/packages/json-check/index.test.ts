import {isValid} from 'json-check';

test("1 is 1", () => {
  expect(1).toBe(1);
});

test("module test", ()=>{
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

    const result = isValid(data, schema);

    expect(result).toBe(true)
});