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

  onProfileClick = (profile, getPassword) => {
    if (profile.parentalPassword !== '' && profile.parentalPassword !== getPassword()) {
      return;
    }

    this.props.history.push(`/media?${queryString.stringify({ profile: profile.name })}`);
  }

  promptPassword = () => prompt('Please enter the password to continue');

  loadProfiles = async () => {
    const profiles = await getProfiles().catch(() => []);
    this.setState({ profiles });
  }

  render = () => (
    <div style={{ margin: '10px' }}>
      <ProfileList
        items={this.state.profiles}
        onClick={profile => this.onProfileClick(profile, this.promptPassword)}
      />
    </div>
  )
}

ProfilesView.propTypes = {
  history: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default ProfilesView;
