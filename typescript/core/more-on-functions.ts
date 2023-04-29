/* function type expressions */
function greeter(fn: (a:string) => void) {
    fn("hello world");
}

greeter((b:string) => { console.log(b)})


/* call signature */
type MyFunction = {
    desc:string
    (num: number): boolean // callable
}

function myDoSth(fn: MyFunction): void {
    console.log(fn.desc + " : " + fn(10))
}

const myFunc = (num: number) => num > 3;
myFunc.desc = "check number bigger than 3"
myDoSth(myFunc)


/* constructor signature */
type MyObject = {
    msg: string
    len: number
    print:()=>void
}

interface MyConstructor {
    new (s: string): MyObject;
}

class MyClass {
    msg: string
    len: number

    constructor(s: string) {
        this.msg = s
        this.len = s.length
    }

    print() {
        console.log(this.msg + " has length of " + this.len)
    }
}

function myNew(clazz: MyConstructor, s: string): void {
    const inst: MyObject = new clazz(s)
    inst.print()
}

myNew(MyClass, "abcdef")

/* generic function */
function map<T,O>(list:T[], fn:(arg: T)=>O): O[] {
    return list.map(fn)
}
const parsed = map<number,number>([1,2,3,4,5], (n)=>n+10)
console.log(parsed)

// constraints
type Subject = {
    name: string
    type: string
}

function issue<T extends Subject>(sub:T): T {
    return {
        ...sub,
        issuer: "ts-coder",
    }
}

const cred = issue({ name: "rhie", type: "dev"})


/* Unknown */
function f1(a: any) {
  a.b(); // OK
}

function f2(a: unknown) {
//   a.b(); // ERROR: 'a' is of type 'unknown'.
}
// The unknown type represents any value. 
// This is similar to the any type, but is safer because it’s not legal to do anything with an unknown value
// This is useful when describing function types 
// because you can describe functions that accept any value without having any values in your function body.
function safeParse(s: string): unknown {
  return JSON.parse(s);
}
 
// Need to be careful with 'obj'!
type Me = { name:string, age:number }
// const obj:Me = safeParse(`{"name":"rhie-coder","age":30}`); // Type 'unknown' is not assignable
const obj = safeParse(`{"name":"rhie-coder","age":30}`); // Type 'unknown' is not assignable
console.log(obj)

const obj2:unknown = safeParse(`{"name":"rhie-coder","age":30}`);
const obj3:Me = safeParse(`{"name":"rhie-coder","age":30}`) as Me;

/* destructing */
type ABC = { a: number; b: number; c: number };
function sum({ a, b, c }: ABC) {
  console.log(a + b + c);
}
sum({ a: 10, b: 3, c: 9 });
