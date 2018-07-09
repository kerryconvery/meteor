import axios from 'axios';
import { getMedia } from './mediaServices';

jest.mock('axios');

describe('MediaServices', () => {
  describe('Get Media', () => {
    it('should return back the payload on success', async () => {
      const response = {
        data: [],
      };

      axios.get.mockResolvedValue(response);

      const payload = await getMedia('profile1', '');

      expect(payload).toEqual(response.data);
    });

    it('should return back an empty list on error', async () => {
      const response = {
        status: 500,
      };

      axios.get.mockRejectedValue(response);

      const payload = await getMedia('profile1', '');

      expect(payload).toEqual([]);
    });
  });
});
