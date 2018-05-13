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

Sign an image

## Compiling the contract

* Most common tool is [Truffle](http://truffleframework.com/)
* Solc
* Golang abigen

## Testing the contract

* Most common tool is [Truffle](http://truffleframework.com/)
* Less common ones: [perigord](https://github.com/polyswarm/perigord), [dapple](https://dapp.readthedocs.io/en/latest/)...
* Manually, thanks to a mock node (ethereumjs)
