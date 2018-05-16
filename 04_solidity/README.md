# Solidity

## In brief

* Other languages compile to the EVM (serpent - python like, LLL - assembly like, solidity - js like)
* Soon webassembly could be compiled to EVM (C/C++/rust and maybe GC languages later!)

## What is a contract

* Types / Control structures
* Local storage
* Function of different types (public, internal) (pure, view)

## WARNING

Solidity is improving but is still very difficult to master.
A contract is written _forever_ on the chain, not js, more realtime programming.
There are tools and companies to audit your contracts, but it's still very tricky.

## Let's try to something today

Proof of attendance (what would be good proof?)

Let's make it simple though!

## Compiling the contract

* Most common tool is [Truffle](http://truffleframework.com/)
* Solc - available from [go-ethereum tools](https://github.com/ethereum/go-ethereum)

Let's use the manual way, solc

get the ABI - application binary interface

```bash
solc contract/EvtAttendance.sol --abi
```

```json
[
  {
    "constant": true,
    "inputs": [{ "name": "_evtId", "type": "uint256" }],
    "name": "hasAttended",
    "outputs": [{ "name": "", "type": "bool" }],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [{ "name": "_evtId", "type": "uint256" }],
    "name": "attend",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [{ "name": "_evtId", "type": "uint256" }],
    "name": "getEvent",
    "outputs": [
      { "name": "", "type": "uint256" },
      { "name": "", "type": "address" },
      { "name": "", "type": "string" },
      { "name": "", "type": "address[]" }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [{ "name": "_title", "type": "string" }],
    "name": "createEvt",
    "outputs": [{ "name": "", "type": "uint256" }],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "getEventCount",
    "outputs": [{ "name": "", "type": "uint256" }],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "anonymous": false,
    "inputs": [
      { "indexed": false, "name": "_organizer", "type": "address" },
      { "indexed": false, "name": "_id", "type": "uint256" },
      { "indexed": false, "name": "_title", "type": "string" }
    ],
    "name": "EvtCreated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      { "indexed": false, "name": "_who", "type": "address" },
      { "indexed": false, "name": "_id", "type": "uint256" }
    ],
    "name": "EvtAttended",
    "type": "event"
  }
]
```

get the bytecode if needed (deploy from node for ex)

```bash
solc contract/EvtAttendance.sol --bin
```

## Testing the contract

* Most common tool is [Truffle](http://truffleframework.com/)
* Less common ones: [perigord](https://github.com/polyswarm/perigord), [dapple](https://dapp.readthedocs.io/en/latest/)...
* Manually, thanks to a mock node (ethereumjs)

continue [here](https://github.com/vincentserpoul/prez-ethereum-dev/blob/master/05_interacting_with_contracts/README.md)