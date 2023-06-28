const crypto = require('crypto');

module.exports = function() {
    // Generate a new RSA key pair with a modulus length of 2048 bits
    const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
    modulusLength: 4200,
    publicKeyEncoding: {
        type: 'pkcs1',
        format: 'pem',
    },
    privateKeyEncoding: {
        type: 'pkcs1',
        format: 'pem',
    },
    });

    return {
        publicKey,
        privateKey,
    }
}
