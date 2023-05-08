/* data through objects (fundamental) */
function greet(person: {name:string, age: number}) {
    return "Hello " + person.name;
}

// OR

interface PersonObject { // type Person = { name:string; age:number}
    name: string;
    age: number;
}

function greet2(person: PersonObject) {
    return "Hello " + person.name;
}

/* optional // readonly */
interface PaintOptions {
    shape: Object;
    xPos?: number;
    yPos?: number;
    readonly prop: string; 
    readonly painter: {name: string; age: number} 
}

/* index signature */
interface StringArray {
    [index: number]: string;
}

const myArr: StringArray = ["hello", "only string here"]
const item = myArr[1]; // const item : string

interface NumberDictionary {
  [index: string]: number;
 
  length: number; // ok
  myProp: number;
//   name: string; // Property 'name' of type 'string' is not assignable to 'string' index type 'number'.
}

interface ReadonlyStringArray {
  readonly [index: number]: string;
}
 
let myArray: ReadonlyStringArray = ["this", "array", "is", "readonly"]
// myArray[2] = "Mallory"; // Index signature in type 'ReadonlyStringArray' only permits reading.

/* extending type */
/*  
interface BasicAddress {
  name?: string;
  street: string;
  city: string;
  country: string;
  postalCode: string;
}

interface AddressWithUnit {
  name?: string;
  unit: string;
  street: string;
  city: string;
  country: string;
  postalCode: string;
}
*/
interface BasicAddress {
  name?: string;
  street: string;
  city: string;
  country: string;
  postalCode: string;
}
 
interface AddressWithUnit extends BasicAddress {
  unit: string;
}

// AND

interface Colorful {
  color: string;
}
 
interface Circle {
  radius: number;
}
 
interface ColorfulCircle extends Colorful, Circle {}
 
const cc: ColorfulCircle = {
  color: "red",
  radius: 42,
};

/* intersection type */
interface MyColor {
  color: string;
}
interface MyCircle {
  radius: number;
}
 
type MyColorCircle = MyColor & MyCircle

// const mcc:MyColorCircle = {radius:30 } // ERROR
const mcc:MyColorCircle = { color: "blue", radius:30 }
