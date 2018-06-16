import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.css';
import ProfilesView from './views/profiles/profilesView';
import MediaView from './views/media/mediaView';

ReactDOM.render(
  <BrowserRouter>
    <Switch>
      <Route exact path='/' component={ProfilesView} />
      <Route exact path='/media' component={MediaView} />
    </Switch>
  </BrowserRouter>,
  document.getElementById('app')
);
