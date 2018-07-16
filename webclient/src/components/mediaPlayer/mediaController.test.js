import React from 'react';
import { shallow } from 'enzyme';
import MediaController from './mediaController';

describe('MediaController', () => {
  it('should match snapshot when playing', () => {
    const wrapper = shallow(<MediaController
      nowPlaying='movie.avi'
      playerState={2}
      position={100}
      duration={1000}
      onRestart={() => {}}
      onPause={() => {}}
      onStop={() => {}}
      onResume={() => {}}
      onPark={() => {}}
    />);

    expect(wrapper).toMatchSnapshot();
  });

  it('should match snapshot when paused', () => {
    const wrapper = shallow(<MediaController
      nowPlaying='movie.avi'
      playerState={1}
      position={100}
      duration={1000}
      onRestart={() => {}}
      onPause={() => {}}
      onStop={() => {}}
      onResume={() => {}}
      onPark={() => {}}
    />);

    expect(wrapper).toMatchSnapshot();
  });
});
