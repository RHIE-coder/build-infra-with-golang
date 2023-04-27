class Immutable {
    final(propertyName, propertyValue) {
        Object.defineProperty(this, propertyName, {
            value: propertyValue,
            writable: false,
        })
    }
}

module.exports = class Wallet extends Immutable{
    constructor(){
        super()
        this.final("OWEN", "0X2894706DEBa1df71735053e8F55F65d34348C051")
        this.final("ALICE", "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293")
        this.final("RHIENY", "0x2D81c2486F2C8a286B067cdEdda2E6815e61DDdA")
        this.final("HUBER", "0xe1625a0d89B0fB0BfC3835E91B1FA8475409aE8E")
        this.final("ANDY", "0xf44ec05E8d0065252e3a9D2b8334225d3Ee71B4B")
    }
}

module.exports = class API extends Immutable {
    constructor(){
        super()
        this.final("API_SERVER","http://localhost:5000")
    }
}

module.exports = class Chains extends Immutable {
    constructor() {
        super()
        this.final("ETH", "ethereum")
        this.final("GEO", "goerli")
        this.final("SEP", "sepolia")
        this.final("LUN", "luniverse")
    }
}