import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import ProfilesView from '../profiles/profilesView';
import MediaView from '../media/mediaView';
import PlayerView from '../media/playerView';

const Routes = () => (
  <BrowserRouter>
    <Switch>
      <Route exact path='/' component={ProfilesView} />
      <Route exact path='/media' component={MediaView} />
    </Switch>
  </BrowserRouter>
);

export default () => (
  <div>
    <Routes />
    <div className='fixed-bottom' >
      <PlayerView />
    </div>
  </div>
);
