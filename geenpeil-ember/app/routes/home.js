import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    share : function() {
      console.log('CLICKY!');
      this.transitionTo('share');
    },

    form : function() {
      console.log('action form');
      this.transitionTo('disclaimer');
    }

  }

    // model: function() {
    //     console.log('running home route model');
    //     return Ember.Object.create({
    //
    //     });
    // }
});
