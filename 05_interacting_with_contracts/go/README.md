# Interacting with contract with golang

## Generating the contract bindings

```bash
abigen --sol=../../04_solidity/contract/EvtAttendance.sol --pkg=main > attendance.go
```

## Interacting

Let's simply get the events count

```bash
go run *.go
```
