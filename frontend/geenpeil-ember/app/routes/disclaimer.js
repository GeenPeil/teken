import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    accept : function() {
      //TODO - get the first input name programatically
      this.transitionTo('input','voornaam');
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
