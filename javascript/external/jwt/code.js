require('dotenv').config();
const jwt = require('jsonwebtoken');
const fs = require('fs');

(async()=>{
    const targetToken = process.env.JWT_TOKEN
    const publicKey = fs.readFileSync('./public-key-loy.pem', 'utf-8');

    const headerBase64 = targetToken.split(".")[0]
    const payloadBase64 = targetToken.split(".")[1]
    const signatureBase64 = targetToken.split(".")[2]

    const decodeHeader = Buffer.from(headerBase64, 'base64').toString('utf-8');
    const decodePayload = Buffer.from(payloadBase64, 'base64').toString('utf-8');
    const decodeSignature = Buffer.from(signatureBase64, 'base64').toString('utf-8');
    // console.log(decodeHeader)
    // console.log(decodePayload)
    // console.log(decodeSignature)

    const keyPair = require("./lib/createRSA")()
    // const myToken = jwt.sign(payload, keyPair.privateKey, {algorithm:'RS256'})
    // const decodeJWT = jwt.decode(myToken)

    console.log(publicKey)
    console.log("size: " + Buffer.byteLength(publicKey))
    console.log(keyPair.publicKey)
    console.log("size: " + Buffer.byteLength(keyPair.publicKey))
    
    // const verifiedPayload = await jwt.verify(targetToken, publicKey, { algorithms: ['RS256'] });
    // console.log('\nJWT verification succeeded.');
    // console.log('Verified JWT payload:');
    // console.log(verifiedPayload)
})()