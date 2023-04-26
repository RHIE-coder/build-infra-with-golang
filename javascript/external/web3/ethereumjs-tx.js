const Tx = require('ethereumjs-tx').Transaction;
require("dotenv").config()

// 트랜잭션 정보
// const txParams = {
//   nonce: '0x31',
//   gasPrice: '0x356214d',
//   gasLimit: '0x5208',
//   to: '0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293',
//   value: '0x' + (100000000000000).toString(16), // 16진수로 변환
// //   data: ''
// };

// const txParams = {
//     nonce: "0x0",
//     to: "0x5BCAeb0a67299973b0B44972D6F0Ba9924102741",
//     data: "0xa9059cbb0000000000000000000000005bcaeb0a67299973b0b44972d6f0ba9924102741000000000000000000000000000000000000000000000000002386f26fc10000",
//     "gasLimit": "0xd6e3",
//     "gasPrice": "0xd4932429"
// }

// from
// data
// gasLimit
// gasPrice
// chainId
// to

const txParams = {
    from: "0x2D81c2486F2C8a286B067cdEdda2E6815e61DDdA",
    data: "0xa9059cbb000000000000000000000000f44ec05e8d0065252e3a9d2b8334225d3ee71b4b000000000000000000000000000000000000000000000000000000000000000a111100001111cc64152eb36ec24a1cb6d52da658ba2605486d13",
    gasLimit: "0x15a7b2979400",
    gasPrice: "0x33450",
    chainId: 256,
    to: "0xa3036f584f4f1c3b352739ca60785e44947e5560",
    nonce:"0x0",
}
console.log(txParams)
// 개인키
const privateKey = Buffer.from(process.env.PRIVATE_KEY, 'hex'); // 개인키 입력

// 서명된 트랜잭션 생성
// const tx = new Tx(txParams, {chain:256})
const tx = new Tx(txParams, {chain:256}) //Error: Chain with ID 256 not supported
tx.sign(privateKey);
const serializedTx = tx.serialize().toString('hex')
console.log(tx.getChainId())

// 서명된 트랜잭션 출력
console.log('Signed Tx: ', '0x' + serializedTx);
console.log('----')
console.log(serializedTx)