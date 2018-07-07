import React from 'react';
import PropTypes from 'prop-types';
import { Progress, Button, ButtonGroup } from 'reactstrap';

export const playerState = {
  paused: 1,
  playing: 2,
};

const MediaPlayer = props => (
  <div style={{ backgroundColor: 'lightgray' }}>
    <div style={{ marginLeft: '10px', marginRight: '10px' }}>
      <span>Now Playing: {props.nowPlaying}</span>
      <Progress
        value={props.position}
        max={props.duration}
      />
      <ButtonGroup className='d-flex justify-content-sm-center'>
        <Button color='info' onClick={props.onRestart}>Restart</Button>
        {
          props.playerState === playerState.playing ?
            <Button color='info' onClick={props.onPause}>Pause</Button> :
            <Button color='info' onClick={props.onResume}>Resume</Button>
        }
        <Button color='info' onClick={props.onStop}>Stop</Button>
        <Button color='info' onClick={props.onPark}>Park</Button>
      </ButtonGroup>
    </div>
  </div>
);

MediaPlayer.propTypes = {
  nowPlaying: PropTypes.string.isRequired,
  playerState: PropTypes.number.isRequired,
  position: PropTypes.number.isRequired,
  duration: PropTypes.number.isRequired,
  onRestart: PropTypes.func.isRequired,
  onPause: PropTypes.func.isRequired,
  onStop: PropTypes.func.isRequired,
  onPark: PropTypes.func.isRequired,
  onResume: PropTypes.func.isRequired,
};

export default MediaPlayer;
