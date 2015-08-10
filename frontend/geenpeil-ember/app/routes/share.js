import Ember from 'ember';

export default Ember.Route.extend({

  //TODO handle social buttons

  actions : {

    back : function() {
      this.transitionTo('home');
    },
    share : function(network) {
      alert('TODO: '+network);
    }

  }

});
