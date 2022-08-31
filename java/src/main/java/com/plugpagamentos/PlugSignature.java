package com.plugpagamentos;

import org.bouncycastle.crypto.params.Ed25519PublicKeyParameters;
import org.bouncycastle.crypto.signers.Ed25519Signer;
import org.bouncycastle.util.encoders.Hex;

public class PlugSignature {

    public static boolean verify(String publicKeyHex, String payload, String signatureTime, String signatureHex) {
        var publicKeyBytes = Hex.decode(publicKeyHex);

        var signatureBytes = Hex.decode(signatureHex);
        var signingDataBytes = (signatureTime+"\n"+payload).getBytes();

        var params = new Ed25519PublicKeyParameters(publicKeyBytes, 0);
        var verifier = new Ed25519Signer();

        verifier.init(false, params);
        verifier.update(signingDataBytes, 0, signingDataBytes.length);

        return verifier.verifySignature(signatureBytes);
    }

}
