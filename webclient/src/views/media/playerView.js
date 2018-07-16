import React from 'react';
import { pause, stop, resume, listen } from '../mediaServices';
import MediaController from '../../components/mediaPlayer/mediaController';
import OfflineButton from '../../components/mediaPlayer/offlineButton';

class PlayerView extends React.Component {
  state = {
    offline: true,
    nowPlaying: '',
    playerState: 0,
    position: 0,
    duration: 0,
  }

  componentDidMount = () => {
    this.connectToPlayer();
  }

  onPlayerUpdated = payload => this.setState({
    offline: false,
    nowPlaying: payload.nowPlaying,
    playerState: payload.state,
    position: payload.position,
    duration: payload.duration,
  });

  onPlayerDisconnected = () => this.setState({ offline: true });
  connectToPlayer = () => listen(this.onPlayerUpdated, this.onPlayerDisconnected);

  render = () => {
    if (this.state.offline) {
      return <OfflineButton onClick={this.connectToPlayer} />;
    }

    if (this.state.playerState !== 0) {
      return (<MediaController
        nowPlaying={this.state.nowPlaying}
        playerState={this.state.playerState}
        position={this.state.position}
        duration={this.state.duration}
        onPause={pause}
        onStop={stop}
        onResume={resume}
        onPark={() => {}}
        onRestart={() => {}}
      />);
    }

    return <div />;
  }
}

export default PlayerView;

