import React from 'react';
import { pause, stop, resume, listen } from '../mediaServices';
import MediaPlayer from '../../components/media/mediaPlayer';

class PlayerView extends React.Component {
  state = {
    nowPlaying: '',
    playerState: 0,
    position: 0,
    duration: 0,
  }

  componentDidMount = () => {
    listen(this.onPlayerUpdated);
  }

  onPlayerUpdated = payload => this.setState({
    nowPlaying: payload.nowPlaying,
    playerState: payload.state,
    position: payload.position,
    duration: payload.duration,
  });

  MediaView = () => (
    <MediaPlayer
      nowPlaying={this.state.nowPlaying}
      playerState={this.state.playerState}
      position={this.state.position}
      duration={this.state.duration}
      onPause={pause}
      onStop={stop}
      onResume={resume}
    />
  )

  render = () => (this.state.playerState !== 0 ? <this.MediaView /> : <div />)
}

export default PlayerView;

