import Ember from 'ember';

export default Ember.Route.extend({

  //TODO get form model
  //TODO render image
  //TODO string data in component
  //TODO signature
  //TODO handle back button

  model : function() {
    //TODO - get the latest timestamped data
    return {
        "voornaam" : "Jan",
        "tussenvoegsel" : "de",
        "achternaam" : "Vries"
    }
  },

  setupController : function(controller, model) {
    controller.set('model',model);
  }

});
