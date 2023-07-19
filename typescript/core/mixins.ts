class Jumpable {
    jump() { }
}

class Duckable {
    duck() { }
}

class Sprite {
    x = 0;
    y = 0;
}

function applyMixins(derivedConstructor: any, extendConstructors: any[]) {
    extendConstructors.forEach((baseCtor) => {
        Object.getOwnPropertyNames(baseCtor.prototype).forEach((name) => {
            Object.defineProperty(derivedConstructor.prototype, name, Object.getOwnPropertyDescriptor(baseCtor.prototype, name) || Object.create(null))
        })
    })
}


(async () => {
    console.log(Object.getOwnPropertyNames(Sprite.prototype))
    console.log(Object.getOwnPropertyNames(Sprite.constructor))
    applyMixins(Sprite, [Duckable, Jumpable]);
    console.log(Object.getOwnPropertyNames(Sprite.prototype))
    const start = new Date().getTime();
    for (let i = 0; i < 100000000; i++) {
        Object.keys(Reflect.construct(Sprite, []))
        // Object.getOwnPropertyNames(new Sprite())
    }
    const end = new Date().getTime();
    console.log(end - start);

    const object1 = {
        property1: 42
    };

    console.log(Object.getOwnPropertyDescriptors(object1))
    Object.seal(object1);
    object1.property1 = 33;
    console.log(object1.property1);
    // Expected output: 33

    // delete object1.property1; // Cannot delete when sealed
    console.log(object1.property1);
    // Expected output: 33

    console.log(Object.getOwnPropertyDescriptors(object1))
})()