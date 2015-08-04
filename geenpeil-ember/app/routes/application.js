import Ember from 'ember';

export default Ember.Route.extend({

    model : function() {
      console.log('loading application route model');
      return Ember.$.ajax({
        type : 'GET',
        url: 'form.json',
        success : function(data) {
          console.log('form data:',data);
          return Ember.Object.create(data);
        },
        error : function() {
          console.error('no form data');
        }
      });
    },

    setupController : function(controller,model) {
      console.log('setupController',controller,model);
      controller.set('model',model);
    }

});
