import React from 'react';
import { shallow } from 'enzyme';
import ProfileList from './profileList';

test('renders a list of profiles', () => {
  const media = [
    {
      name: 'profile1',
    },
    {
      name: 'profile2',
    },
  ];
  const wrapper = shallow(<ProfileList items={media} />);
  expect(wrapper).toMatchSnapshot();
});

