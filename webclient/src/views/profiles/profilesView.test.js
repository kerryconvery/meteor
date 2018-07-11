import React from 'react';
import { shallow } from 'enzyme';
import ProfilesView from './profilesView';
import * as Services from '../mediaServices';

jest.mock('../mediaServices');
global.prompt = jest.fn();

describe('ProfilesView', () => {
  let props;

  const mountComponent = () => {
    props = {
      history: [],
    };

    return shallow(<ProfilesView {...props} />, { disableLifecycleMethods: true });
  };

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should call loadProfiles after mounting', () => {
    const wrapper = mountComponent();
    const loadMediaSpy = jest.spyOn(wrapper.instance(), 'loadProfiles');

    wrapper.instance().componentDidMount();

    expect(loadMediaSpy).toHaveBeenCalled();
  });

  it('should set profile list onClick to onProfileClick', () => {
    const profile = {
      name: 'profile1',
    };

    const wrapper = mountComponent();

    const profileList = wrapper.find('ProfileList');
    const onClickSpy = jest.spyOn(wrapper.instance(), 'onProfileClick');
    profileList.prop('onClick')(profile);
    expect(onClickSpy).toBeCalledWith(profile, wrapper.instance().promptPassword);
  });

  it('should call prompt', () => {
    const wrapper = mountComponent();

    wrapper.instance().promptPassword();

    expect(global.prompt).toHaveBeenCalled();
  });

  describe('ProfileClick', () => {
    it('should add the item name to the history if there is no password', () => {
      const profile = {
        name: 'profile1',
        parentalPassword: '',
      };

      const wrapper = mountComponent();

      wrapper.instance().onProfileClick(profile, () => {});
      expect(props.history[0]).toEqual(`/media?profile=${profile.name}`);
    });

    it('should add the item name to the history if there the supplied password is correct', () => {
      const profile = {
        name: 'profile1',
        parentalPassword: 'q',
      };

      const wrapper = mountComponent();

      wrapper.instance().onProfileClick(profile, () => 'q');
      expect(props.history[0]).toEqual(`/media?profile=${profile.name}`);
    });

    it('should not add the item name to the history if there the supplied password is incorrect', () => {
      const profile = {
        name: 'profile1',
        parentalPassword: 'q',
      };

      const wrapper = mountComponent();

      wrapper.instance().onProfileClick(profile, () => 'a');
      expect(props.history.length).toEqual(0);
    });
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
