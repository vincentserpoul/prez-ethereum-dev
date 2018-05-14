import React, { Component } from "react";
import PropTypes from "prop-types";

import { getInstance } from "./contract";

class Instance extends Component {
  constructor(props) {
    super(props);
    this.state = { events: [] };
  }

  addEvent = evt => {
    const evts = this.state.events;
    evts.push(evt);
    this.setState({ events: evts });
  };

  componentWillMount = () => {
    const deployedAt = "0x7e5836836bb23ba8b06295c4ffbf168456a5976d";
    const att = getInstance(this.props.ethProvider, deployedAt);
    this.setState({ address: deployedAt });
    att.getEventCount().then(evtCount => {
      console.log(evtCount[0].toNumber());
      for (let i = 0; i < evtCount[0].toNumber(); i++) {
        att.getEvent(i).then(evt => {
          console.log(evt);
          this.addEvent(evt);
        });
      }
    });
  };

  render = () => {
    const evts = [];
    this.state.events.forEach(evt => {
      evts.push(
        <p key={evt[0].toNumber()}>
          id:{evt[0].toNumber()}
          <br />organizer: {evt[1]}
          <br />title: {evt[2]}
          <br />attendees: {evt[3].join(", ")}
        </p>
      );
    });
    return (
      <div>
        <h1>Events listed for contract @{this.state.address}</h1>
        {evts}
      </div>
    );
  };
}

Instance.propTypes = {
  ethProvider: PropTypes.object
};

export default Instance;
