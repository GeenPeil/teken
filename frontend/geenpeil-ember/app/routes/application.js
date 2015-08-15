import Ember from 'ember';

export default Ember.Route.extend({

    model : function() {
      return Ember.$.ajax({
        type : 'GET',
        url: 'form_debug.json',
        error : function(e) {
          console.error('no form data:',e);
        }
      });
    },

    setupController : function(controller,model) {
      if(typeof model === 'string') {
        model = JSON.parse(model);
      }

      //TODO - check for any stored form values here

      controller.set('model',Ember.Object.create(model));
    }

});
