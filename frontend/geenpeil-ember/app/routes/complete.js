import Ember from 'ember';

export default Ember.Route.extend({

  actions : {

    done : function() {
      location.reload();
    }

  }

});
