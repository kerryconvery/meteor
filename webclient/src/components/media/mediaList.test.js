import React from 'react';
import { shallow } from 'enzyme';
import MediaList from './mediaList';


describe('MediaList', () => {
  it('should call onItemClicked', () => {
    const media = [
      {
        name: 'media1',
        isDirector: false,
        uri: '',
      },
    ];
    const mockOnClick = jest.fn();
    const wrapper = shallow(<MediaList items={media} profile='profile1' onItemClicked={mockOnClick} />);

    wrapper.find('ListGroupItem').simulate('click');

    expect(mockOnClick).toHaveBeenCalled();
  });
});

test('renders a list of media', () => {
  const media = [
    {
      name: 'media1',
      isDirectory: false,
      uri: '',
    },
    {
      name: 'media2',
      isDirectory: true,
      uri: '',
    },
  ];
  const wrapper = shallow(<MediaList items={media} profile='profile1' onItemClicked={() => {}} />);
  expect(wrapper).toMatchSnapshot();
});

