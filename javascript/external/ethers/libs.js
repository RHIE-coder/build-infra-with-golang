const { ethers } = require('ethers');

module.exports = {
    getPrivateKeyFromMnemonic(mnemonic) {
        const wallet = ethers.Wallet.fromPhrase(mnemonic);
        const privateKey = wallet.privateKey;
        return privateKey
    },
}