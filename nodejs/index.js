const PlugSignature = require('./plugSignature');

const publicKey = '-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAS8O+o7Cqd3mtpriJvOVK/cdo8se5VU5vZkCMUttd0WA=\n-----END PUBLIC KEY-----\n';

// change some value below and see the result change
const sampleHttp = {
  headers: {
    'X-Plug-Date': 1661795163719,
    'X-Plug-Signature': '5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01'
  },
  body: '{"event":"ping","payload":{"object":{}}}'
};

const isSignatureValid = PlugSignature.verify(publicKey, sampleHttp.body, sampleHttp.headers['X-Plug-Date'], sampleHttp.headers['X-Plug-Signature']);

if (isSignatureValid) {
  console.log('Congrats ... signature is valid!!!');
} else {
  console.log('Ops ... Signature is not valid!');
}