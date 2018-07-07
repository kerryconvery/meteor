import React from 'react';
import PropTypes from 'prop-types';
import { Button, ButtonGroup } from 'reactstrap';

const profileItem = onClick => item => (
  <Button outline color='primary' onClick={() => onClick(item)}>{item.name}</Button>
);

const ProfileList = props => (
  <ButtonGroup className='d-flex justify-content-md-center' vertical>{props.items.map(profileItem(props.onClick))}</ButtonGroup>
);

ProfileList.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.isRequired,
  })).isRequired,
  onClick: PropTypes.func.isRequired,
};

export default ProfileList;
