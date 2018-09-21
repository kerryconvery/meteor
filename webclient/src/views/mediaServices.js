import axios from 'axios';

export const launchMedia = (profile, uri) => axios.post(`/api/media?profile=${profile}&uri=${uri}`);
export const getMedia = (profile, uri) => axios.get(`/api/profiles/${profile}/media`, { params: { uri } })
  .then(res => res.data)
  .catch(() => []);
export const getProfiles = () => axios.get('/api/profiles').then(res => res.data);
export const pause = () => axios.post('api/media/pause');
export const stop = () => axios.delete('api/media');
export const resume = () => axios.post('api/media/resume');
export const restart = () => axios.post('api/media/restart');
export const listen = (onMessage, onClose) => {
  const ws = new WebSocket(`ws://${window.location.host}/ws`);
  ws.addEventListener('message', event => onMessage(JSON.parse(event.data)));
  ws.addEventListener('close', () => onClose());
  return ws;
};
