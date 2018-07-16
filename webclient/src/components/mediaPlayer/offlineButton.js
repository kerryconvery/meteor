import React from 'react';
import PropTypes from 'prop-types';
import { Button } from 'reactstrap';
import Backboard from './backboard';

const OfflineButton = props => (
  <Backboard>
    <div className='d-flex justify-content-sm-center'>
      <Button color='warning' onClick={props.onClick} >Offline</Button>
    </div>
  </Backboard>
);

OfflineButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};

export default OfflineButton;
