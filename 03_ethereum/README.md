# Ethereum

## In brief

programmable contracts, State of the contract vars on the blockchain
Turing complete
compile code to EVM (assembly/lisp)

example:

```evmcode
PUSH1 0 CALLDATALOAD SLOAD NOT PUSH1 9 JUMPI STOP JUMPDEST PUSH1 32 CALLDATALOAD PUSH1 0 CALLDATALOAD SSTORE
```

## Anatomy of an ethereum transaction

[Ethereum blockchain explorer](https://etherscan.io/)

from - to - value - gas limit - gas price - input data (for contracts)
SIGNED

gas: unit of measuring the computational work (example Watts)
gas limit: maximum you re willing to spend
gas price: at what price

continue [here](https://github.com/vincentserpoul/prez-ethereum-dev/blob/master/04_solidity/README.md)