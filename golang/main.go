package main

import (
	"fmt"
)

type HttpHeaderSample struct {
	Header map[string]interface{}
	Body   string
}

func main() {
	publicKey := "-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAS8O+o7Cqd3mtpriJvOVK/cdo8se5VU5vZkCMUttd0WA=\n-----END PUBLIC KEY-----\n"

	sampleHttp := HttpHeaderSample{
		Header: map[string]interface{}{
			"X-Plug-Date":      int64(1661795163719),
			"X-Plug-Signature": "5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01",
		},
		Body: "{\"event\":\"ping\",\"payload\":{\"object\":{}}}",
	}

	isValid, err := Verify(
		publicKey,
		sampleHttp.Body,
		sampleHttp.Header["X-Plug-Date"].(int64),
		sampleHttp.Header["X-Plug-Signature"].(string),
	)

	if err != nil {
		panic(err)
	}

	if isValid {
		fmt.Println("Congrats ... signature is valid!!!")
	} else {
		fmt.Println("Ops ... Signature is not valid!")
	}
}
