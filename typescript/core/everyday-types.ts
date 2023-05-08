/* 
    기본 타입
    with
    Type Annotations
*****************/
/* boolean */
const a1: boolean = true;

/* number */
const a2: number = 100;

/* string */
const a3 : string = "hello world";

/* object */
const a4: object = { a: 10, b: 20 }
const a44: {a:string, b:number} = {a:"hello", b:33};
// const a444: {a:string, b:number} = {a:"hello", b:33, c:true}; // ERROR
// const a4444: {a:string, b:number} = {a:22, b:33}; // ERROR
const a44444: {a:string, b?:string} = {a:"good"} // optional properties

/* array */
// TypeScript does not assume that arrays are immutable
const a5: number[] = [1, 2, 3, 4]
const a55: (number|string)[] = [1, "2"] // union type
const a555: Array<number> = [5, 6, 7]
const a5555: Array<number|string> = [1, 2, "world"]

/* tuple */
const a6:[number, string, number] = [100, "typescript", 200]
// const a66:[number, string, number] = ["100", "typescript", 200] // ERROR

/* enum */
enum A7 { a, b, c, }
const a7: A7 = A7.c

/* any */
let a8: any = 100
a8 = "hello"
a8 = [1, 2, 3, 4]

/* void */
// const a9: void = null // ERROR
const a99: void = undefined
// const a999: void = {} // ERROR
// const a9999:void = "" // ERROR

/* never */
// Never: 함수의 리턴 타입으로 사용되며 항상 오류를 출력하거나 리턴 값을 절대로 내보내지 않음(무한루프)
function endless(): never {
    while(true) {
        console.log("this function won't end")
    }
}

/* function */
function add(n1:number, n2:number):number{
    return n1 + n2;
}

function getById(id:number|string) {
    // it called `Union Type`
}

// what if function expression?
let getDataFromFile = function(path:string): string[] {
    return ["file binary"]
}
    // OR
let getDataFromFile2:(path:string) => string[]

getDataFromFile2 = function(path) {
    return ["file binary"]
}
    // OR
let getDataFromFile3:(path:string) => string[] = function(path) {
    return ["file binary"]
}


/* type */
type Person = {
    name: string
    age: number
}

function toStr(person: Person): string { // type alias
    return `${person.name} :: ${person.age}`
}

/* interface */
interface Point {
    x: number;
    y: number;
}

function printCoord(pt: Point) {
    console.log(`(${pt.x}, ${pt.y})`)
}

// type vs interface ?
// type cannot be re-opened to add new properties vs an interface which is always extendable.

    // 1. Extending (add properties)

// - interface
interface IAnimal {
    name: string
}

interface IDog extends IAnimal {
    bark: string
}

const cooky: IDog = {
    name: "cooky",
    bark: "wal wal"
}

// - type
type TAnimal = {
    name: string
}

type TDog = TAnimal & { // intersection
    bark: string
}

const mirae: TDog = {
    name: "mirae",
    bark: "meong meong"
}

    // 2. Adding new fields to an existing

// - interface
interface IPerson {
    name: string
}

interface IPerson {
    age: number
}

const student: IPerson = {
    name: "rhie",
    age: 20,
}

// - type
/* 
type TPerson = {
    name: string
}

type TPerson = {
    age: number
}
*/ // ERROR

/********************************/
/* type assertion */
// 주의: 컴파일 시 제거되므로 런타임환경에서 type assertion 기능에 대한 효과를 기대 no no 

type UserData = {
    name:string
    age:number
}

// const jenny: UserData= {name:"jenny"} // ERROR
// const num = 100 as string // ERROR
// type assertion은 컴파일 시에만 체크가 된다
// 그러나, 아래는 의도한 바와 다르게 컴파일이 성공적으로 되므로, 런타임 때 에러가 일어날 수도 있다.
const jenny = {name:"jenny"} as UserData // 방법 1
const owen = <UserData>{name:"owen"} // 방법2
// 그러면 왜 type assertion이 필요할까?

class Machine {}
class Calculator extends Machine { calculate(){} }
class Scanner extends Machine { scan(){} }

function operate(machine:Machine): void {
    const machineName:string = machine.constructor.name

    if(machineName === "Calculator") {
        // machine.calculate() // ERROR
        (machine as Calculator).calculate()
        return
    } 

    if(machineName === "Scanner") {
        // machine.scan() // ERROR
        (machine as Scanner).scan()
        return
    }

    console.log("untyped machine")
    return
}

/* literal type */
let unliteral = "Hello World"
unliteral = "can change. it is string type."

const literal = "Hello World"
// literal = "cannot change. it is 'Hello World' type"

let foo1: "Hello World" = "Hello World"
let foo2: 1 | 2 | 3
foo2 = 2
// foo2 = 4 // ERROR

/* literal interface */
// TypeScript assumes that the properties of that object might change values later.
const foo3 = { counter: 0 }; // counter is number type
if (true) {
    foo3.counter = 1;
}

function bar(url:string, method: "GET" | "POST") {}

const foo4 = { url: "https://example.com", method: "GET" }; // url and method are string
// bar(foo4.url, foo4.method); // ERROR(method is "GET" type)
bar(foo4.url, foo4.method as "GET")

// url is "https://example.com" type
// method is "GET" type
// to convert the entire object to be type literals:
const foo5 = { url: "https://example.com", method: "GET" } as const 
bar(foo5.url, foo5.method)

/* Non-null Assertion Operator `!` */
function liveDangerously(x?: number | null) {
//   console.log(x.toFixed()); // ERROR: 'x' is possibly 'null' or 'undefined'
  console.log(x!.toFixed()) 
}
