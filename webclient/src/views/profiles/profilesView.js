import React from 'react';
import axios from 'axios';
import { NavLink } from 'react-router-dom';

export default class Profiles extends React.Component {
  state = {
    profiles: [],
  }

  componentDidMount() {
    axios({
      method: 'GET',
      url: '/api/profiles',
    }).then(res => this.setState({ profiles: res.data }));
  }

  render() {
    return (
      <ul>
        {this.state.profiles.map(p => <li><NavLink exact to={`/media?profile=${p.name}`}>{p.name}</NavLink></li>)}
      </ul>
    );
  }
}
