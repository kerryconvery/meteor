import React from 'react';
import { shallow } from 'enzyme';
import MediaView from './mediaView';
import * as Services from '../mediaServices';

jest.mock('../mediaServices');

describe('MediaView', () => {
  const mountComponent = () => {
    const props = {
      location: {
        search: 'profile=p1&uri=u1',
      },
      history: [],
    };

    return shallow(<MediaView location={props.location} history={props.history} />, { disableLifecycleMethods: true });
  };

  beforeEach(() => {
    Services.getMedia.mockResolvedValue([]);
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should call loadMedia after mounting', () => {
    const wrapper = mountComponent();
    const loadMediaSpy = jest.spyOn(wrapper.instance(), 'loadMedia');

    wrapper.instance().componentDidMount();

    expect(loadMediaSpy).toHaveBeenCalledWith({ profile: 'p1', uri: 'u1' });
  });

  it('should call loadMedia when receiving new props if the location search has changed', () => {
    const wrapper = mountComponent();
    const loadMediaSpy = jest.spyOn(wrapper.instance(), 'loadMedia');

    wrapper.instance().componentWillReceiveProps({ location: { search: 'profile=p2&uri=u2' } });

    expect(loadMediaSpy).toHaveBeenCalledWith({ profile: 'p2', uri: 'u2' });
  });

  it('should not call loadMedia when receiving new props if the location search has not changed', () => {
    const wrapper = mountComponent();
    const loadMediaSpy = jest.spyOn(wrapper.instance(), 'loadMedia');

    wrapper.instance().componentWillReceiveProps({ location: { search: 'profile=p1&uri=u1' } });

    expect(loadMediaSpy).not.toHaveBeenCalled();
  });

  describe('loadMedia', () => {
    it('should call getMedia and put the media on the state', async () => {
      const params = { profile: 'profile1', uri: '/folder1/folder2' };
      const media = [
        { name: 'media1', uri: '', isDirectory: false },
        { name: 'media2', uri: '', isDirectory: false },
      ];
      Services.getMedia.mockResolvedValue(media);
      const wrapper = mountComponent();

      await wrapper.instance().loadMedia(params);

      expect(Services.getMedia).toHaveBeenCalledWith(params.profile, params.uri);
      expect(wrapper.state().profile).toEqual(params.profile);
      expect(wrapper.state().media).toEqual(media);
    });
  });

  describe('onItemClicked', () => {
    it('should call navigate folder when item is a directory', () => {
      const item = {
        isDirectory: true,
        uri: 'testuri',
      };

      const wrapper = mountComponent();

      wrapper.setState({ profile: 'profile1' });

      const navigateFolderSpy = jest.spyOn(wrapper.instance(), 'navigateFolder');

      wrapper.instance().onItemClicked(item);

      expect(navigateFolderSpy).toHaveBeenCalledWith('profile1', 'testuri');
    });

    it('should call launch media when item is not a directory', () => {
      const item = {
        isDirectory: false,
        uri: 'testuri',
      };

      const wrapper = mountComponent();

      wrapper.setState({ profile: 'profile1' });

      wrapper.instance().onItemClicked(item);

      expect(Services.launchMedia).toHaveBeenCalledWith('profile1', 'testuri');
    });
  });
});
