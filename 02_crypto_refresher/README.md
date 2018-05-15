# 02 Crypto refresher

## IMPORTANT reminders

_Rule #1_ - Never create your own crypto

_Rule #2_ - Never implement your own crypto library

## Hash

[Recommendation](http://latacora.singles/2018/04/03/cryptographic-right-answers.html#hashing-algorithm): SHA-2 or SHA-3

Ethereum uses SHA-3 (Keccak256)

```bash
go run 02_crypto_refresher/hash/main.go
```

## Asymmetric signatures

[Recommendation](http://latacora.singles/2018/04/03/cryptographic-right-answers.html#asymmetric-signatures): Nacl or Ed25519

Ethereum and bitcoin use secp256k1 (elliptic curve - ECDSA) - [potentially unsafe](https://safecurves.cr.yp.to/)

```bash
go run 02_crypto_refresher/signature/main.go
```

continue [here](./03_ethereum/README.md)