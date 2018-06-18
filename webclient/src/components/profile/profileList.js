import React from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

const profileItem = () => item => (
  <li>
    <NavLink exact to={`/media?profile=${item.name}`}>{item.name}</NavLink>
  </li>
);

const ProfileList = props => (
  <ul>{props.items.map(profileItem(props.onItemClicked))}</ul>
);

ProfileList.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.isRequired,
  })).isRequired,
  onItemClicked: PropTypes.func.isRequired,
};

export default ProfileList;
