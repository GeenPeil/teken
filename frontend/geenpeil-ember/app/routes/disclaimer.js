import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    accept : function() {
      this.transitionTo('form',1);
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
