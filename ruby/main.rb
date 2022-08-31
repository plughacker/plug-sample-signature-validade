require 'ed25519'

# https://github.com/RubyCrypto/ed25519
# gem "ed25519"

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

# Convert hex to bin
def hex_to_bin(hex)
  hex = '0' << hex unless hex.length.even?
  hex.scan(/[A-Fa-f0-9]{2}/).inject('') { |encoded, byte| encoded << [byte].pack('H2') }
end

public_key_hex = '4bc3bea3b0aa7779ada6b889bce54afdc768f2c7b9554e6f66408c52db5dd160'

public_key_bytes = hex_to_bin(public_key_hex)
verify_key = Ed25519::VerifyKey.new(public_key_bytes)

http_sample = {
  'body' => '{"event":"ping","payload":{"object":{}}}',
  'headers' => {
    'X-Plug-Date' => '1661795163719',
    'X-Plug-Signature' => '5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01'
  }
}

sig_bytes = hex_to_bin(http_sample['headers']['X-Plug-Signature'])

message = http_sample['headers']['X-Plug-Date'] + "\n" + http_sample['body']

begin
  verify_key.verify(sig_bytes, message)
  puts 'Congrats ... signature is valid!!!'
rescue StandardError
  puts 'Ops ... Signature is not valid!'
end