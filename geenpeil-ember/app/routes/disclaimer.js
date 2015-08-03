import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    accept : function() {
      //todo get the first input name programatically
      this.transitionTo('input','voornaam');
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
