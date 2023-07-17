const sdk = require('./hub-luniverse-sdk');
const axios = require('axios');
(async()=>{
    require("dotenv").config(__dirname);
    
    const {
        Luniverse,
        WalletHandler,
    } = sdk;

    console.log(WalletHandler.generateMnemonic())
    const mnemonic = process.env.LUNIVERSE_MNEMONIC;
    const wallet0 = WalletHandler.createWalletFromMnemonic(mnemonic);
    const wallet1 = WalletHandler.createWalletFromMnemonic(mnemonic, 1);
    const wallet2 = WalletHandler.createWalletFromMnemonic(mnemonic, 2);

    console.log(wallet0.mnemonic.phrase)
    console.log(wallet1.mnemonic.phrase)
    console.log(wallet2.mnemonic.phrase)

    const sender = wallet0.address
    console.log(wallet0.privateKey)
    const receiver = "0xe7517164cBd1943eD5dffe1fbAC14E467Db41a75";

    const api = new Luniverse(axios.create({
        baseURL: 'https://devapi.publish-hub.io/',
        headers: {
            Authorization: `Bearer ${process.env.JWT_TOKEN_DEV}`,
            [`x-service-name`]:'newsthomas',
        }
    }))

    let resp;
    
    resp = await api.tokenBalance(sender)
    const balance = resp.data
    console.log(resp)
    console.log(balance)


    resp = await api.accountNonce(sender)
    const nonce = resp.data.nonce
    console.log(resp);
    console.log(nonce);


    resp = await api.tokenRawTxn({
        amount: "0.0001",
        from:sender,
        to: receiver,
    })
    const raw = resp.data;
    console.log(resp);
    console.log(raw);

    
    const signedTxHex = (await WalletHandler.sign(wallet0, {
        from: raw.from,
        to: raw.to,
        gas: raw.gas_limit,
        gasPrice: raw.gas_price,
        nonce: nonce,
        data: raw.data,
    }));
    const signedTx = signedTxHex.slice(2);
    console.log(signedTx);




    // resp = await api.sendTxn(signedTx);
    // const result = resp.data;
    // console.log(resp);
    // console.log(result);
})()
