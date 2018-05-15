# How to interact with the blockchain

## JSON-RPC

* web3_clientVersion
* web3_sha3
* net_version
* net_peerCount
* net_listening
* eth_protocolVersion
* eth_syncing
* eth_coinbase
* eth_mining
* eth_hashrate
* eth_gasPrice
* eth_accounts
* eth_blockNumber
* eth_getBalance
* ...

## Which languages are most used to interact with ethereum

1.  Javascript - DAPPs
2.  Golang - main ethereum client (abigen - demo?)
3.  Rust - second biggest ethereum client

## Deploying the contract - or JS or golang or remix browser

* Truffle (compile, deploy...)
* Web3
* ethjs
* Parity + Parity UI
* golang (abigen)
* [remix browser](https://remix.ethereum.org)

## DAPPs - decentralized apps

[dapp demo](./dapp)

## Components of a dapp

* Static websites (can be stored on IPFS)
* Using web3 to interact with a local node (or remote) via jsonrpc
* Wallet (metamask, parity, mew, ledger...)

## Our attendance dapp demo

[See here](./dapp)

## Implementation in golang

### Ropsten is too long to sync, let's take the opportunity to use parity dev and its UI

launching a dev parity instance

```bash
docker volume create --driver=local parity-dev-demo

docker run -dit --name parity-dev-demo --volume parity-dev-demo:/root/.local/share/io.parity.ethereum/ -p 10545:10545 -p 10546:10546 parity/parity:v1.10.3 --config=dev --jsonrpc-port=10545 --jsonrpc-cors=all --ws-port=10546 --port=40303 --unsafe-expose --reseal-min-period=0 --no-persistent-txqueue --base-path=/root/.local/share/io.parity.ethereum
```

launching the [UI](https://github.com/parity-js/shell/releases)

on linux

```bash
./parity-ui --ws-interface http://127.0.0.1 --ws-port 10546
```

on a mac

```bash
open -a /Applications/Parity\ UI.app/Contents/MacOS/Parity\ UI --args  --ws-interface http://127.0.0.1 --ws-port 10546
```

Deploy your contract within the UI

modify the address in the first line of the main.go
run!

```bash
go run go/*.go
```

## Feedback

![Feedback wanted](./QR_code_ZNK8RV9.png)

[feedback wanted](https://fr.surveymonkey.com/r/ZNK8RV9)