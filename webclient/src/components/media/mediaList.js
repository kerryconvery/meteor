import React from 'react';
import PropTypes from 'prop-types';
import { ListGroup, ListGroupItem } from 'reactstrap';

const mediaItem = (profile, onClick) => item => (
  <ListGroup>
    <ListGroupItem key={item.name} tag='button' action onClick={() => onClick(item)}>
      <img src={`/api/profiles/${profile}/media/thumbnail?uri=${item.uri}`} alt='' />
      {item.name}
    </ListGroupItem>
  </ListGroup>
);

const MediaList = props => (
  <ul>{props.items.map(mediaItem(props.profile, props.onItemClicked))}</ul>
);

MediaList.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.isRequired,
    isDirectory: PropTypes.bool,
    uri: PropTypes.string,
  })).isRequired,
  profile: PropTypes.string.isRequired,
  onItemClicked: PropTypes.func.isRequired,
};

export default MediaList;