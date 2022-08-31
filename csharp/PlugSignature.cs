namespace PlugPagamentos;

using System.Text;
using NSec.Cryptography;

public class PlugSignature
{
    private static Ed25519 algorithm = SignatureAlgorithm.Ed25519;

    public static Boolean verify(String payload, long signatureDate, String signatureHex, string publicKeyRaw)
    {
        var signature = Convert.FromHexString(signatureHex);
        var payloadToVerify = Encoding.UTF8.GetBytes($"{signatureDate}\n{payload}");
        var publicKeyBytes = Encoding.UTF8.GetBytes(publicKeyRaw);
        
        var publicKey = PublicKey.Import(algorithm, publicKeyBytes, KeyBlobFormat.PkixPublicKeyText);
        return algorithm.Verify(publicKey, payloadToVerify, signature);
    }

}