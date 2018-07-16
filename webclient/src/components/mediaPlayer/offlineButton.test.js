import React from 'react';
import { shallow } from 'enzyme';
import OfflineButton from './offlineButton';

describe('OfflineButton', () => {
  it('should match snapshot', () => {
    const wrapper = shallow(<OfflineButton onClick={() => {}} />);

    expect(wrapper).toMatchSnapshot();
  });
});
