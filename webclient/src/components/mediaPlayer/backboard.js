import React from 'react';
import PropTypes from 'prop-types';

const Backboard = props => (
  <div style={{ backgroundColor: 'lightgray' }}>
    {props.children}
  </div>
);

Backboard.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Backboard;
