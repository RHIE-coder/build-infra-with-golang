// const addMod = require('add-all');
const addMod = require('./index');

test("1 is 1", () => {
  expect(1).toBe(1);
});

test("module function check",()=>{
    console.log(addMod)
    const arr_sum_result = addMod.arr_sum([1, 2, 3, 4, 5]);
    console.log(arr_sum_result)
    expect(arr_sum_result).toBe(15);

    const data = {
        sum_target: [1, 2, 3, 4, 5],
        foo: 123,
        bar: "abc",
    }
    const arr_sum_from_object_result = addMod.arr_sum_from_object(data)
    expect(arr_sum_from_object_result).toBe(15);
})