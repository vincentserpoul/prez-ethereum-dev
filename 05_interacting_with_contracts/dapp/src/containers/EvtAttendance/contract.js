export const getInstance = (ethProvider, deployedAddr) => {
  const abi = [
    {
      constant: true,
      inputs: [{ name: "_evtId", type: "uint256" }],
      name: "hasAttended",
      outputs: [{ name: "", type: "bool" }],
      payable: false,
      stateMutability: "view",
      type: "function"
    },
    {
      constant: false,
      inputs: [{ name: "_evtId", type: "uint256" }],
      name: "attend",
      outputs: [],
      payable: false,
      stateMutability: "nonpayable",
      type: "function"
    },
    {
      constant: true,
      inputs: [{ name: "_evtId", type: "uint256" }],
      name: "getEvent",
      outputs: [
        { name: "", type: "uint256" },
        { name: "", type: "address" },
        { name: "", type: "string" },
        { name: "", type: "address[]" }
      ],
      payable: false,
      stateMutability: "view",
      type: "function"
    },
    {
      constant: false,
      inputs: [{ name: "_title", type: "string" }],
      name: "createEvt",
      outputs: [{ name: "", type: "uint256" }],
      payable: false,
      stateMutability: "nonpayable",
      type: "function"
    },
    {
      constant: true,
      inputs: [],
      name: "getEventCount",
      outputs: [{ name: "", type: "uint256" }],
      payable: false,
      stateMutability: "view",
      type: "function"
    },
    {
      anonymous: false,
      inputs: [
        { indexed: false, name: "_organizer", type: "address" },
        { indexed: false, name: "_id", type: "uint256" },
        { indexed: false, name: "_title", type: "string" }
      ],
      name: "EvtCreated",
      type: "event"
    },
    {
      anonymous: false,
      inputs: [
        { indexed: false, name: "_who", type: "address" },
        { indexed: false, name: "_id", type: "uint256" }
      ],
      name: "EvtAttended",
      type: "event"
    }
  ];
  return ethProvider.contract(abi).at(deployedAddr);
};
