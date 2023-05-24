/*  
node call eth_getBalance 0x2894706deba1df71735053e8f55f65d34348c051 latest
node call eth_getCode 0x39c44CD8432e45B1bA8EDe1Ca6f5020ea04438E0 latest
node call web3_clientVersion
node call web3_sha3 0x68656c6c6f20776f726c64
node call net_version
node call net_listening
node call net_peerCount
node call eth_protocolVersion
node call eth_syncing
node call eth_coinbase
node call eth_mining
node call eth_hashrate
node call eth_gasPrice
node call eth_accounts
node call eth_blockNumber
node call eth_getStorageAt 0xE36AE64156db78dd4797864E9A2f3C1C40625BF3 0x0 latest
node call eth_getTransactionCount 0x468f9E09806256209388d9c0fBd911C4D49F9fbe latest
node call eth_getBlockTransactionCountByHash 0x9539c415683c031d3fe82318e9c43201d8cb202474e73bab8292dc8e26f66599
node call eth_getBlockTransactionCountByNumber 0xb7ef39
node call eth_getUncleCountByBlockHash 0xa2e87148321048b7b098252abfd6a22adeddadd1608a226b72ae330457b8f5a8
node call eth_getUncleCountByBlockNumber 0xe8
node call eth_getLogs '{"fromBlock":"0x335024","toBlock":"0x34d385","address":"0x468f9E09806256209388d9c0fBd911C4D49F9fbe"}'
node call eth_getLogs '{"fromBlock":3362852,"toBlock":3462021,"address":"0x468f9E09806256209388d9c0fBd911C4D49F9fbe"}'

 -- Unable To Infura --
X node call eth_sign 0x2894706debA1DF71735053E8f55f65D34348c051 0x68656c6c6f20776f726c64
X node call eth_sendTransaction '{"from": "0x2894706debA1DF71735053E8f55f65D34348c051", "to": "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293", "gas":"0x76c0", "gasPrice": "0x3b9aca08", "value":"100000000"}'
X node call eth_signTransaction '{"from": "0x2894706debA1DF71735053E8f55f65D34348c051", "to": "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293", "gas":"0x76c0", "gasPrice": "0x3b9aca08", "value":"100000000"}'

eth_sign
eth_signTransaction
eth_sendTransaction
eth_sendRawTransaction
eth_call

eth_estimateGas

eth_getBlockByHash
eth_getBlockByNumber
eth_getTransactionByHash
eth_getTransactionByBlockHashAndIndex
eth_getTransactionByBlockNumberAndIndex
eth_getTransactionReceipt
eth_getUncleByBlockHashAndIndex
eth_getUncleByBlockNumberAndIndex
eth_getCompilers

eth_compileSolidity
eth_compileLLL
eth_compileSerpent

eth_newFilter
eth_newBlockFilter
eth_newPendingTransactionFilter
eth_uninstallFilter
eth_getFilterChanges
eth_getFilterLogs

eth_getWork
eth_submitWork
eth_submitHashrate
*/

// node call toHex 3304322       --> 0x326b83
// node call toHex "hello world" --> 0x68656c6c6f20776f726c64

const web3 = new (require('web3'))();

(async()=>{
    const target = process.argv[2]
    console.log(target)
    if(target == "toHex") {
        console.log(web3.utils.toHex(process.argv[3]))
        return
    }

    const sourceDir = `./json-rpc-spec/${target}`
    
    await require(sourceDir)(process.argv.slice(3))
})()

/*  
eth_sign, eth_signTransaction, eth_sendTransaction, eth_sendRawTransaction 같은 Ethereum JSON-RPC 메서드는 Infura에서 제공하는 모든 노드에서 사용 가능합니다. 그러나 이러한 메서드를 사용하기 위해서는 이러한 메서드를 호출하는 Ethereum 계정의 프라이빗 키가 필요합니다.

Infura는 사용자의 개인 프라이빗 키를 보관하지 않습니다. 

이를 위해 사용자는 로컬에서 실행되는 Ethereum 클라이언트 노드를 설치하고, 해당 노드에 Ethereum 계정의 프라이빗 키를 등록하여 이러한 메서드를 사용할 수 있습니다. 예를 들어, Geth 또는 Parity와 같은 Ethereum 클라이언트 노드를 설치하고, 이를 사용하여 해당 메서드를 호출할 수 있습니다.
*/