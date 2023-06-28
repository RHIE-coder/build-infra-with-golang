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
        })).data
    }
}


(async()=>{
    const caller = new JsonRpcDialer("http://192.168.100.73:8999");
    console.log("request!")
    console.log(
        await caller.call("net_listening")
    )
    console.log("end")
})()