using PlugPagamentos;

var sampleHttp = new {
  headers = new {
    XPlugDate = 1661795163719,
    XPlugSignature = "5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01"
  },
  body = "{\"event\":\"ping\",\"payload\":{\"object\":{}}}"
};


var publicKeyBytes = "-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAS8O+o7Cqd3mtpriJvOVK/cdo8se5VU5vZkCMUttd0WA=\n-----END PUBLIC KEY-----\n";

var isValid = PlugSignature.verify(sampleHttp.body, sampleHttp.headers.XPlugDate, sampleHttp.headers.XPlugSignature, publicKeyBytes);

if (isValid) {
  Console.WriteLine("Congrats ... signature is valid!!!");
} else {
  Console.WriteLine("Ops ... Signature is not valid!");
}
