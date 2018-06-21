import React from 'react';
import { getProfiles } from '../mediaServices';
import ProfileList from '../../components/profile/profileList';

class ProfilesView extends React.Component {
  state = {
    profiles: [],
  }

  componentDidMount() {
    this.loadProfiles();
  }

  loadProfiles = async () => {
    const profiles = await getProfiles().catch(() => []);
    this.setState({ profiles });
  }

  render = () => <ProfileList items={this.state.profiles} />;
}

export default ProfilesView;
