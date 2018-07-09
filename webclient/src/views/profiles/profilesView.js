import React from 'react';
import queryString from 'query-string';
import PropTypes from 'prop-types';
import { getProfiles } from '../mediaServices';
import ProfileList from '../../components/profile/profileList';

class ProfilesView extends React.Component {
  state = {
    profiles: [],
  }

  componentDidMount() {
    this.loadProfiles();
  }

  onProfileClick = (profile) => {
    this.props.history.push(`/media?${queryString.stringify({ profile: profile.name })}`);
  }

  loadProfiles = async () => {
    const profiles = await getProfiles().catch(() => []);
    this.setState({ profiles });
  }

  render = () => (
    <div style={{ margin: '10px' }}>
      <ProfileList items={this.state.profiles} onClick={this.onProfileClick} />
    </div>
  )
}

ProfilesView.propTypes = {
  history: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
};

export default ProfilesView;
