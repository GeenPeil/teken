import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    clicky : function() {
      console.log('CLICKY!');
    }

  }

    // model: function() {
    //     console.log('running home route model');
    //     return Ember.Object.create({
    //
    //     });
    // }
});
