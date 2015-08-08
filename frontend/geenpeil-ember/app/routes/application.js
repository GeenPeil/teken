import Ember from 'ember';

export default Ember.Route.extend({

    model : function() {
      console.log('loading application route model');
      return Ember.$.ajax({
        type : 'GET',
        url: 'form.json',
        error : function() {
          console.error('no form data');
        }
      });
    },

    setupController : function(controller,model) {
      console.log('setupController',controller,model);

      if(typeof model === 'string') {
        model = JSON.parse(model);
      }

      //TODO - check for any stored form values here

      controller.set('model',Ember.Object.create(model));
    }

});
