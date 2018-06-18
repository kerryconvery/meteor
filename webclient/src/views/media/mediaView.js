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
    this.loadMedia(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location.search !== this.props.location.search) {
      this.loadMedia(nextProps);
    }
  }

  onItemClicked = item => (
    item.isDirectory ? this.navigateFolder(this.state.profile, item.uri) : launchMedia(this.state.profile, item.uri)
  );

  navigateFolder = (profile, uri) => this.props.history.push(`/media?${queryString.stringify({ profile, uri })}`);

  loadMedia = (props) => {
    const params = new URLSearchParams(props.location.search);
    const profile = params.get('profile');
    const uri = params.get('uri');

    getMedia(profile, uri).then(media => this.setState({ profile, media }));
  }

  render = () => <MediaList items={this.state.media} profile={this.state.profile} onItemClicked={this.onItemClicked} />
}

MediaView.propTypes = {
  location: PropTypes.shape({
    search: PropTypes.string,
  }).isRequired,
  history: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
};

export default MediaView;
