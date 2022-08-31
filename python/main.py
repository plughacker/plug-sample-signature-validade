from cryptography.hazmat.primitives.asymmetric.ed25519 import Ed25519PublicKey

#
# Format plug webhook public key:
# 
# -----BEGIN PUBLIC KEY----
# MCowBQYDK2VwAyEAS8O+o7Cqd3mtpriJvOVK/cdo8se5VU5vZkCMUttd0WA=
# -----END PUBLIC KEY-----
#
# Use this online tools and convert to hex publicKey format
# https://go.dev/play/p/YGI55tRH565
#

public_key_byte = bytes.fromhex('4bc3bea3b0aa7779ada6b889bce54afdc768f2c7b9554e6f66408c52db5dd160')
public_key = Ed25519PublicKey.from_public_bytes(public_key_byte)

httpSample = {
  'body': '{"event":"ping","payload":{"object":{}}}',
  'headers': {
    'X-Plug-Date': '1661795163719',
    'X-Plug-Signature': '5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01'
  }
}

signaturePayload = str.encode(httpSample['headers']['X-Plug-Date'] + "\n" + httpSample['body'])
signature = bytes.fromhex(httpSample['headers']['X-Plug-Signature'])

try:
  public_key.verify(signature, signaturePayload)
  print('Congrats ... signature is valid!!!')
except:
  print('Ops ... Signature is not valid!')
