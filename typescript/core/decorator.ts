import "reflect-metadata";


function Klazz(name:string) {
    console.log("klaazzzz")
    return function(constructor:Function) {
        Reflect.defineMetadata("bbbbb", 2000, constructor)
        console.log(Reflect.getOwnMetadataKeys(constructor))
    }

}

function ShowMe(flag:boolean) {
    console.log("showwwww me")
    return function(target:any, propKey:string, descriptor:PropertyDescriptor) {
        console.log("Show Me CB")
        console.log(Reflect.getOwnMetadataKeys(target))
    }
}

function configurable(value: string) {
  return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
    console.log("success config")
    console.log(target)
    target.AAAA = value; 
    console.log(target)
  };
}

const k1 = Symbol('one')
const k2 = Symbol('two')

function justCheck(){
    console.log("hfehfhefh")
    return Reflect.metadata(k1, 3333);
}
function justCheck2(name:string){
    console.log("421-4r12hfsankgj")
    return Reflect.metadata(k2, 4444);
}

function justCheck3(){
    return Reflect.metadata("c1c", "oppa");
}
function justCheck4(){
    return Reflect.metadata("d1d", "zzang");
}

@Klazz("hello")
@Reflect.metadata("aaaaa",1000)
class Validator{

    @justCheck()
    @justCheck3()
    attrA:string

    @justCheck2("2321")
    @justCheck4()
    nnn:number

    constructor(attrA:string, nnn:number){
        this.attrA = attrA;
        this.nnn = nnn;
    }

    @configurable("ha hah")
    get N() {
        return this.nnn + 1000000
    }

    set N(newa:number) {
        this.nnn = newa * 2;
    }

    @ShowMe(true)
    myMethod(){
        console.log("myMethod()")
    }
}


@Reflect.metadata("key", "base value")
class B {
    get prop(): number { return 0; }
}

class C extends B{ }

// "base value", metadata was not defined on C but was defined on it's prototype B
console.log(Reflect.getMetadata("key", C)); 

// undefined, metadata was not defined on C
console.log(Reflect.getOwnMetadata("key", C)); 


console.log('412412');

const formatMetadataKey = Symbol("format");
function format(formatString: string) {
  return Reflect.metadata(formatMetadataKey, formatString);
}
function getFormat(target: any, propertyKey: string) {
  return Reflect.getMetadata(formatMetadataKey, target, propertyKey);
}

class Greeter {
  @format("Hello, %s")
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    let formatString = getFormat(this, "greeting");
    return formatString.replace("%s", this.greeting);
  }
}
(async()=>{
    console.log('+_-')
    console.log()

    const valid = new Validator("hhhhh", 321321)
    console.log("--------------------")
    
    valid.myMethod() 

    console.log(Reflect.getMetadata("design:paramtypes", Validator))
    console.log(Reflect.getOwnMetadata("bbbbb", Validator))
    console.log(Reflect.getMetadata("bbbbb", valid.constructor))
    console.log(Reflect.getOwnMetadata("aaaaa", valid.constructor))
    console.log(Reflect.getMetadata("aaaaa", Validator))


    console.log(valid.N)
    valid.N = 200
    console.log(valid.N)
    console.log(valid)

    console.log(Reflect.getMetadata(k1, valid, "attrA"))
    console.log(Reflect.getMetadata(k2, valid, "nnn"))
    console.log(Reflect.getMetadata("c1c", valid, "attrA"))
    console.log(Reflect.getMetadata("d1d", valid, "nnn"))
})()