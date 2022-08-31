const crypto = require('crypto');

module.exports = {
  /**
   * Example of a function that validates the payload signature
   * @param {string} publicKey       Key returned on webhook creation
   * @param {string} payload         Event sent by the plug. This information goes in the body http
   * @param {number} signatureTime   Time in unix timestamp that the event was subscribed to
   * @param {string} signature       Hash sha512
   * @returns {bool}
   */
  verify: function(publicKey, payload, signatureTime, signature) {
    const payloadUtf8 = Buffer.from(payload, 'utf-8').toString();
    const signatureBuffer = Buffer.from(signature, 'hex');
    const data = Buffer.from(`${signatureTime}\n${payloadUtf8}`);
    return crypto.verify(null, data, publicKey, signatureBuffer);
  }
}