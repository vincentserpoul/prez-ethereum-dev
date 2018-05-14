import React, { Component } from "react";
import Eth from "ethjs";

import Instance from "./containers/EvtAttendance";

import "./App.css";

class App extends Component {
  constructor() {
    super();
    const eth = new Eth(window.web3.currentProvider);
    this.state = { ethProvider: eth, accounts: [] };
  }

  componentWillMount = () => {
    this.state.ethProvider
      .accounts()
      .then(accts => this.setState({ accounts: accts }));
  };

  render = () => {
    const accts = [];
    this.state.accounts.forEach(a => accts.push(<p key={a}>{a}</p>));
    return (
      <section>
        <div className="accounts">
          <h1>accounts available</h1>
          {accts}
        </div>
        <Instance
          ethProvider={this.state.ethProvider}
          account={this.state.accounts[0]}
        />
      </section>
    );
  };
}

export default App;
