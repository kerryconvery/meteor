import React from 'react';
import PropTypes from 'prop-types';

export const playerState = {
  playing: 1,
  paused: 2,
};

const MediaPlayer = props => (
  <div>
    <button onClick={props.onRestart}>Restart</button>
    {
      props.playerState === playerState.playing ?
        <button onClick={props.onPause}>Pause</button> :
        <button onClick={props.onResume}>Resume</button>
    }
    <button onClick={props.onStop}>Stop</button>
    <button onClick={props.onPark}>Park</button>
  </div>
);

MediaPlayer.propTypes = {
  playerState: PropTypes.number.isRequired,
  onRestart: PropTypes.func.isRequired,
  onPause: PropTypes.func.isRequired,
  onStop: PropTypes.func.isRequired,
  onPark: PropTypes.func.isRequired,
  onResume: PropTypes.func.isRequired,
};

export default MediaPlayer;
