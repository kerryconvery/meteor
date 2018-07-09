import React from 'react';
import { shallow } from 'enzyme';
import ProfileList from './profileList';

describe('ProfileList', () => {
  it('should call on click', () => {
    const media = [{ name: 'profile1' }];
    const mockClick = jest.fn();

    const wrapper = shallow(<ProfileList items={media} onClick={mockClick} />);

    wrapper.find('Button').simulate('click');

    expect(mockClick).toHaveBeenCalledWith(media[0]);
  });
});

test('renders a list of profiles', () => {
  const media = [
    {
      name: 'profile1',
    },
    {
      name: 'profile2',
    },
  ];
  const wrapper = shallow(<ProfileList items={media} onClick={() => {}} />);
  expect(wrapper).toMatchSnapshot();
});

