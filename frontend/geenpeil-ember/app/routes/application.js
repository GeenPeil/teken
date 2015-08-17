import Ember from 'ember';

export default Ember.Route.extend({

    healthy : true,

    beforeModel : function(transition) {
      console.log('Application beforeModel',transition);
      if(transition.targetName !== 'down') {
        return Ember.$.ajax({
          type : 'GET',
          url: 'https://teken.geenpeil.nl/pechtold/health-check',
          success : function(d) {
            console.log('health check ok',d);
          }.bind(this),
          error : function(e) {
            console.log('health check not ok',e);
            this.set('healthy',false);
            this.transitionTo('down');
          }.bind(this)
        });
      }
    },

    model : function() {
      return Ember.$.ajax({
        type : 'GET',
        url: 'form.json',
        error : function(e) {
          console.error('no form data:',e);
        }
      });
    },

    afterModel : function(model,transition) {
      if(!this.get('healthy') && transition.targetName === 'down') {
        //let through
      }
      else if(transition.targetName !== 'home') {
        this.transitionTo('home');
      }
    },

    setupController : function(controller,model) {
      if(typeof model === 'string') {
        model = JSON.parse(model);
      }

      controller.set('model',Ember.Object.create(model));
    }

});
