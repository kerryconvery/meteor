import React from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';
import queryString from 'query-string';
import MediaList from '../../components/media/mediaList';

export default class Media extends React.Component {
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
    item.isDirectory ?
      this.navigateFolder(this.state.profile, item.uri) :
      this.launchMedia(this.state.profile, item.uri));

  navigateFolder = (profile, uri) => this.props.history.push(`/media?${queryString.stringify({ profile, uri })}`);
  launchMedia = (profile, uri) => axios({ method: 'POST', url: `/api/media?profile=${profile}&uri=${uri}` })

  loadMedia = (props) => {
    const params = new URLSearchParams(props.location.search);
    const profile = params.get('profile');
    const uri = params.get('uri');

    axios({ method: 'GET', url: `/api/profiles/${profile}/media${uri ? `?uri=${uri}` : ''}` })
      .then(res => this.setState({ profile: params.get('profile'), media: res.data }));
  }

  render() {
    return (
      <MediaList items={this.state.media} profile={this.state.profile} onItemClicked={this.onItemClicked} />
    );
  }
}

Media.propTypes = {
  location: PropTypes.shape({
    search: PropTypes.string,
  }).isRequired,
  history: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
};
