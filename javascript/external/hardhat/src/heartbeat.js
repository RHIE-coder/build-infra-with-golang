const axios = require('axios');


class Dialer {

    #baseUrl

    constructor(url) { 
        this.#baseUrl = url
    }

    async request(method, data, headers) {
        return await axios({
            method,
            url: this.#baseUrl,
            data: data ?? {},
            headers: headers ?? {
                [`Content-Type`]: 'application/json',
            }
        })
    }

}

class JsonRpcDialer extends Dialer {
    constructor(url) {
        super(url);
    }

    async call(method, params) {
        const requestBody = {
            jsonrpc:"2.0",
            method,
            params: params ?? [], 
        }
        console.log(requestBody)
        return (await this.request("post", {
            jsonrpc:"2.0",
            method,
            params: params ?? [],
            id:1, 
        })).data
    }
}


(async()=>{
    // const caller = new JsonRpcDialer("http://192.168.100.73:8999");
    // const caller = new JsonRpcDialer("http://localhost:8545");
    const caller = new JsonRpcDialer("http://localhost:9545");
    console.log("request!")
    console.log(
        await caller.call("eth_coinbase")
    )
    // console.log(
    //     BigInt(
    //             (await caller.call("eth_chainId")).result
    //         ).toString()
    // )
    console.log("end")
})()