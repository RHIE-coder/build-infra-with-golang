/*  
node call eth_getBalance 0xE36AE64156db78dd4797864E9A2f3C1C40625BF3 latest
node call eth_getCode 0xE36AE64156db78dd4797864E9A2f3C1C40625BF3 latest


*/

(async()=>{
    const target = process.argv[2]
    const sourceDir = `./json-rpc-spec/${target}`
    await require(sourceDir)(process.argv.slice(3))
})()
