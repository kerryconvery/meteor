
import React from 'react';
import { shallow } from 'enzyme';
import * as Services from '../mediaServices';
import PlayerView from './playerView';

jest.mock('../mediaServices');

describe('PlayerView', () => {
  const mountComponent = () => shallow(<PlayerView />);

  it('should have an initial state', () => {
    const wrapper = mountComponent();

    expect(wrapper.state()).toEqual({
      offline: true,
      nowPlaying: '',
      playerState: 0,
      position: 0,
      duration: 0,
    });
  });

  it('should start listening for media player events', () => {
    const wrapper = mountComponent();
    const connectSpy = jest.spyOn(wrapper.instance(), 'connectToPlayer');

    wrapper.instance().componentDidMount();
    expect(connectSpy).toHaveBeenCalled();
  });

  it('should render offline button when not online', () => {
    const wrapper = mountComponent();

    expect(wrapper.find('OfflineButton').exists()).toEqual(true);
    expect(wrapper).toMatchSnapshot();
  });

  it('should not render player when not online', () => {
    const wrapper = mountComponent();

    expect(wrapper.find('MediaPlayer').exists()).toEqual(false);
  });

  it('should not render offline button when online', () => {
    const wrapper = mountComponent();

    wrapper.setState({ offline: false });
    wrapper.update();

    expect(wrapper.find('OfflineButton').exists()).toEqual(false);
  });

  it('should render player when online', () => {
    const wrapper = mountComponent();

    wrapper.setState({ offline: false, playerState: 2 });
    wrapper.update();

    const player = wrapper.find('MediaController');

    expect(player.exists()).toEqual(true);
    expect(wrapper).toMatchSnapshot();
  });

  it('should not render offline button or player when online but not playing', () => {
    const wrapper = mountComponent();

    wrapper.setState({ offline: false, playerState: 0 });
    wrapper.update();

    const player = wrapper.find('MediaController');

    expect(player.exists()).toEqual(false);
    expect(wrapper.find('OfflineButton').exists()).toEqual(false);
  });

  describe('connectToPlayer', () => {
    it('should call media services listen', () => {
      const wrapper = mountComponent();

      wrapper.instance().connectToPlayer();
      expect(Services.listen).toHaveBeenCalledWith(wrapper.instance().onPlayerUpdated, wrapper.instance().onPlayerDisconnected);
    });
  });

  describe('onPlayerUpdated', () => {
    it('should update the state with player state', () => {
      const payload = {
        nowPlaying: 'movie.avi',
        state: 2,
        position: 100,
        duration: 1000,
      };

      const wrapper = mountComponent();

      wrapper.instance().onPlayerUpdated(payload);
      wrapper.update();

      expect(wrapper.state()).toEqual({
        offline: false,
        nowPlaying: payload.nowPlaying,
        playerState: payload.state,
        position: payload.position,
        duration: payload.duration,
      });
    });
  });

  describe('onPlayerDisconnected', () => {
    it('should set the offline state to true', () => {
      const wrapper = mountComponent();

      wrapper.setState({ offline: false });
      wrapper.instance().onPlayerDisconnected();

      expect(wrapper.state().offline).toEqual(true);
    });
  });
});
