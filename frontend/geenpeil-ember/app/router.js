import EmberRouter from '@ember/routing/router';
import config from './config/environment';

const Router = EmberRouter.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('home', { path: '/' });

  this.route('share', { path: '/delen' });

  this.route('disclaimer', { path: '/voorwaarden' });

  this.route('form', { path: '/formulier/:section' });

  this.route('check', { path: '/controleren' });

  this.route('preview', { path: '/voorbeeld' });

  this.route('complete', { path: '/klaar' });

  this.route('down', { path: '/kapot' });
});

export default Router;
