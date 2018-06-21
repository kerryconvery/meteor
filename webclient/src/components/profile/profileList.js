import React from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

const profileItem = item => (
  <li key={item.name}>
    <NavLink exact to={`/media?profile=${item.name}`}>{item.name}</NavLink>
  </li>
);

const ProfileList = props => (
  <ul>{props.items.map(profileItem)}</ul>
);

ProfileList.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.isRequired,
  })).isRequired,
};

export default ProfileList;
