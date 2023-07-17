import * as math from "./modules/math";
import { User } from "./modules/data";
import call from "./modules/data";
export function doOutput() {
    console.log(math.add(10, 20));
    console.log(math.mul(10, 20));
    var user = new User("rhie");
    console.log(user.name);
    call();
}
