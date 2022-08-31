<?php
require(__DIR__ . '/vendor/autoload.php');

// install lib
// composer require phpseclib/phpseclib:~3.0

use phpseclib3\Crypt\EC;
use phpseclib3\Crypt\PublicKeyLoader;

// public key receive in create webhook
$key = PublicKeyLoader::load("-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAS8O+o7Cqd3mtpriJvOVK/cdo8se5VU5vZkCMUttd0WA=\n-----END PUBLIC KEY-----\n", false);

// sample http request
$sampleHttp = [
  'headers' => [
    'X-Plug-Date' => '1661795163719',
    'X-Plug-Signature' => '5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01'
  ],
  'body' => '{"event":"ping","payload":{"object":{}}}'
];

// rule signature
$data = "{$sampleHttp['headers']['X-Plug-Date']}\n{$sampleHttp['body']}";

// convert signature hex to bin
$signature = hex2bin($sampleHttp['headers']['X-Plug-Signature']);

if ($key->verify($data, $signature)) {
  echo 'Congrats ... signature is valid!!!';
} else {
  echo 'Ops ... Signature is not valid!';
}