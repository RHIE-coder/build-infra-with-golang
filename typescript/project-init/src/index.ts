import * as math from "./modules/math"
import { User } from "./modules/model"
import call from "./modules/model"

export function doOutput():void{
    console.log(math.add(10, 20))
    console.log(math.mul(10, 20))
    const user = new User("rhie")
    console.log(user.name)
    call()
}