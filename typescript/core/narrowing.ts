/* Union 연산으로 타입 판단이 애매할 때 Narrowing이 필요함 */
function needNarrowing(data: string | number): (number|void) {
    // return data.length // ERROR: string일 때만 유효함

    // do narrowing using `typeof`
    if(typeof data === "string") {
        return data.length // (parameter)data:string
    }

    if(typeof data === "number") {
        return data // (parameter)data:number
    }

    return
}
// 이 밖에도 in, instanceof 등을 활용하여 type narrowing

/* type predicates */
type Fish = { getName: () => void}
type Bird = { getName: () => void}

type Animal = Fish | Bird

function BirdGuard(pet :Animal): pet is Bird {
    return (pet as Bird).getName !== undefined 
}

let foo: Animal= {
    getName() {
        console.log("flying")
    }
}

foo // Animal
if(BirdGuard(foo)) {
    foo.getName() // expected Bird but it is still Animal
}else {
    foo.getName() // expected Fish but it occurs "does not exist" error
}