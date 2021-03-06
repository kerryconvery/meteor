import React from 'react';
import PropTypes from 'prop-types';
import { ListGroup, ListGroupItem } from 'reactstrap';

const getMediaSource = (profile, item) => {
  if (item.isDirectory) {
    return '/web/assets/folder.png';
  }
  return `/api/profiles/${profile}/media/thumbnail?uri=${item.uri}`;
};

const mediaItem = (profile, onClick) => item => (
  <ListGroupItem key={item.name} tag='button' action onClick={() => onClick(item)}>
    <img src={getMediaSource(profile, item)} alt='' />
    <span style={{ marginLeft: '5px' }}>{item.name}</span>
  </ListGroupItem>
);

const MediaList = props => (
  <ListGroup>{props.items.map(mediaItem(props.profile, props.onItemClicked))}</ListGroup>
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
