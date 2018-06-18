import React from 'react';
import { pause, stop, resume } from '../mediaServices';
import MediaPlayer, { playerState } from '../../components/media/mediaPlayer';

export default () => <MediaPlayer playerState={playerState.playing} onPause={pause} onStop={stop} onResume={resume} />;
