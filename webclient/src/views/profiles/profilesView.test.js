import React from 'react';
import { shallow } from 'enzyme';
import ProfilesView from './profilesView';
import * as Services from '../mediaServices';

jest.mock('../mediaServices');

describe('ProfilesView', () => {
  const props = {
    history: [],
  };

  const mountComponent = () => (
    shallow(<ProfilesView {...props} />, { disableLifecycleMethods: true })
  );

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should call loadProfiles after mounting', () => {
    const wrapper = mountComponent();
    const loadMediaSpy = jest.spyOn(wrapper.instance(), 'loadProfiles');

    wrapper.instance().componentDidMount();

    expect(loadMediaSpy).toHaveBeenCalled();
  });

  it('should add the item name to the history', () => {
    const profile = {
      name: 'profile1',
    };

    const wrapper = mountComponent();

    wrapper.instance().onProfileClick(profile);
    expect(props.history[0]).toEqual(`/media?profile=${profile.name}`);
  });

  it('should set profile list onClick to onProfileClick', () => {
    const wrapper = mountComponent();
    const profileList = wrapper.find('ProfileList');
    expect(profileList.prop('onClick')).toEqual(wrapper.instance().onProfileClick);
  });

  describe('loadProfiles', () => {
    it('should call getProfiles and put the returned profiles on the state', async () => {
      const wrapper = mountComponent();
      const profiles = [
        {
          name: 'profile1',
          mediaPath: 'profile1/media',
          parentalPassword: '123',
          mediaArgs: '/arg1 /arg2',
        },
        {
          name: 'profile2',
          mediaPath: 'profile2/media',
          mediaArgs: '/arg1',
        },
      ];

      Services.getProfiles.mockResolvedValue(profiles);
      await wrapper.instance().loadProfiles();

      expect(Services.getProfiles).toHaveBeenCalled();
      expect(wrapper.state().profiles).toEqual(profiles);
    });

    it('should put an empty profile set on the state when getProfiles returns an error', async () => {
      const wrapper = mountComponent();

      Services.getProfiles.mockRejectedValue({});

      await wrapper.instance().loadProfiles();

      expect(Services.getProfiles).toHaveBeenCalled();
      expect(wrapper.state().profiles).toEqual([]);
    });
  });
});
