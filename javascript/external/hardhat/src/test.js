class AAA {

    #num
    constructor(num) {
        this.#num = num
    }

    get NUM(){
        return this.#num
    }
}

const a = new AAA(111);
console.log(a)

console.log(Object.entries(a))