import React from 'react';
import { pause, stop, resume, listen } from '../mediaServices';
import MediaPlayer, { playerState } from '../../components/media/mediaPlayer';

class PlayerView extends React.Component {
  state = {}

  componentDidMount = () => {
    listen(this.onPlayerUpdated);
  }

  onPlayerUpdated = (payload) => {
    console.log(payload);
  }

  render = () => <MediaPlayer playerState={playerState.playing} onPause={pause} onStop={stop} onResume={resume} />
}

export default PlayerView;

