import Ember from 'ember';
import config from './config/environment';

var Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.resource('home', { path: '/' });

  this.resource('share', { path: '/delen' });

  this.resource('disclaimer', { path: '/voorwaarden' });

  this.resource('form', { path: '/formulier/:section' });

  this.resource('check', { path: '/controleren' });

  this.resource('preview', { path: '/voorbeeld' });

  this.resource('complete', { path: '/klaar' });

  this.resource('down', { path: '/kapot' });
});


export default Router;
