const axios = require('axios')

// Follow : https://www.jsonrpc.org/specification
/*  
--> {"jsonrpc": "2.0", "method": "foobar, "params": "bar", "baz]
<-- {"jsonrpc": "2.0", "error": {"code": -32700, "message": "Parse error"}, "id": null}

--> {"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}
<-- {"jsonrpc": "2.0", "result": 19, "id": 3}
*/
module.exports = class JsonRpcRequester{

    #_baseURL
    #_axios

    constructor(baseURL){
        this.#_baseURL = baseURL;
        this.#_axios = axios.create({baseURL})
    }

    get baseURL(){
        return this.#_baseURL
    }

    async dial(req){
        const body = { jsonrpc: "2.0", ...req};
        return (await this.#_axios.post("", body)).data
    }

}