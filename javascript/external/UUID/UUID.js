const uuid = require('uuid');
const {
    v1: uuidV1,
    v3: uuidV3,
    v4: uuidV4,
    v5: uuidV5,
} = uuid;

const {v1, v5} = uuid;

const v1options = {
    node: [0x02, 0x23, 0x45, 0x67, 0x89, 0xab],
    clockseq: 0x3949,

};

const getV3 = () => v1(v1options);

function getUUID(uuid){
    return v5((new Date().getTime() + Math.random()).toString(), getV3())
}


(async()=>{
    const valueOfUUIDV4 = uuidV4();
    console.log(valueOfUUIDV4)
    console.log(uuid.validate(valueOfUUIDV4))
    console.log(uuid.version(valueOfUUIDV4))
    const getValue = getUUID()
    console.log(getValue);
    console.log(uuid.version(getValue))

})()