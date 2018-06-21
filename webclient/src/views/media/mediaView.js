import React from 'react';
import PropTypes from 'prop-types';
import queryString from 'query-string';
import { getMedia, launchMedia } from '../mediaServices';
import MediaList from '../../components/media/mediaList';

class MediaView extends React.Component {
  state = {
    profile: '',
    media: [],
  }

  componentDidMount() {
    this.loadMedia(queryString.parse(this.props.location.search));
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location.search !== this.props.location.search) {
      this.loadMedia(queryString.parse(nextProps.location.search));
    }
  }

  onItemClicked = item => (
    item.isDirectory ?
      this.navigateFolder(this.state.profile, item.uri) :
      launchMedia(this.state.profile, item.uri)
  );

  navigateFolder = (profile, uri) => {
    this.props.history.push(`/media?${queryString.stringify({ profile, uri })}`);
  };

  loadMedia = async (params) => {
    const media = await getMedia(params.profile, params.uri).catch(() => []);
    this.setState({ profile: params.profile, media });
  }

  render = () => (
    <MediaList
      items={this.state.media}
      profile={this.state.profile}
      onItemClicked={this.onItemClicked}
    />
  );
}

MediaView.propTypes = {
  location: PropTypes.shape({
    search: PropTypes.string,
  }).isRequired,
  history: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
};

export default MediaView;
