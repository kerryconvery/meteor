import React from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';
import { NavLink } from 'react-router-dom';

export default class Media extends React.Component {
  state = {
    profile: '',
    media: [],
  }

  componentDidMount() {
    const params = new URLSearchParams(this.props.location.search);
    axios({
      method: 'GET',
      url: `/api/profiles/${params.get('profile')}/media`,
    }).then(res => this.setState({ profile: params.get('profile'), media: res.data }));
  }

  render() {
    return (
      <ul>
        {this.state.media.map(p => (
          <li>
            <NavLink exact to={`/media?profile=${this.state.profile}&uri=${p.uri}`}>{p.name}</NavLink>
          </li>))}
      </ul>
    );
  }
}

Media.propTypes = {
  location: PropTypes.shape({
    search: PropTypes.string,
  }).isRequired,
};
