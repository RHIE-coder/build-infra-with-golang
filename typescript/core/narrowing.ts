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
type Fish = { name: string, swim: () => void}
type Bird = { name: string, fly: () => void}

type Animal = Fish | Bird

function getInstance(name:string, adoptType: "fish" | "bird"): Animal {
    if(adoptType === "fish") {
        return {
            name,
            swim() {
                console.log("swimming")
            }
        }
    } else {
        return {
            name,
            fly() {
                console.log("flying")
            }
        }
    }
}

function BirdGuard(pet: Animal): pet is Bird {
    // return (pet as Bird).name !== undefined 
    return true
}


function typeGuardExample() {
    type PetType = "bird" | "fish"
    const name: string = process.argv[2]
    const type: PetType = process.argv[3] as PetType
    if (!(type === "bird" || type === "fish")) {
        throw new SyntaxError("the type must be 'bird' or 'fish'")
    }
    const pet: Animal = getInstance(name, type)

    if(BirdGuard(pet)) {
        pet.fly()
    } else {
        pet.swim()
    }
}

typeGuardExample()






