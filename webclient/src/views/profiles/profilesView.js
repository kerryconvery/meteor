import React from 'react';
import { getProfiles } from '../mediaServices';
import ProfileList from '../../components/profile/profileList';

class ProfilesView extends React.Component {
  state = {
    profiles: [],
  }

  componentDidMount() {
    getProfiles().then(profiles => this.setState({ profiles }));
  }

  render = () => <ProfileList items={this.state.profiles} />;
}

export default ProfilesView;
