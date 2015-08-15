import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    share : function() {
      this.transitionTo('share');
    },

    form : function() {
      this.transitionTo('disclaimer');
    }

  }

});
