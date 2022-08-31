package com.plugpagamentos;

import java.util.HashMap;

class Main {
    public static void main(String[] args) {
        var publicKey = "4bc3bea3b0aa7779ada6b889bce54afdc768f2c7b9554e6f66408c52db5dd160";

        var sampleHttp = new HashMap<String, String>();
        sampleHttp.put("body", "{\"event\":\"ping\",\"payload\":{\"object\":{}}}");
        sampleHttp.put("X-Plug-Date", "1661795163719");
        sampleHttp.put("X-Plug-Signature", "5b20c43cfd55f0c1884196786d58392e1828232e05d66e9153032ae9d374bff785d0efc4ed850ea5856b71c950bd2a1914f376381386355fce2c51163ed26e01");

        var isValid = PlugSignature.verify(
                publicKey,
                sampleHttp.get("body"),
                sampleHttp.get("X-Plug-Date"),
                sampleHttp.get("X-Plug-Signature"));

        if (isValid) {
            System.out.println("Congrats ... signature is valid!!!");
        } else {
            System.out.println("Ops ... Signature is not valid!");
        }
    }
}