import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

      back : function() {
        window.history.back();
      },

      preview : function() {
        this.transitionTo('preview');
      }

  }

});
