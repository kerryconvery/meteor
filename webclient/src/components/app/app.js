import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Profiles from '../profiles/profiles';
import Media from '../media/media';

export default () => (
  <Switch>
    <Route exact path='/' component={Profiles} />
    <Route exact path='/media' component={Media} />
  </Switch>
);

