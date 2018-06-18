import axios from 'axios';

export const launchMedia = (profile, uri) => axios.post(`/api/media?profile=${profile}&uri=${uri}`);
export const getMedia = (profile, uri) => axios.get(`/api/profiles/${profile}/media`, { params: { uri } }).then(res => res.data);
export const getProfiles = () => axios.get('/api/profiles').then(res => res.data);
export const pause = () => axios.post('api/media/pause');
export const stop = () => axios.delete('api/media');
export const resume = () => axios.post('api/media/resume');
