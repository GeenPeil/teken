import Route from '@ember/routing/route';
import Object from '@ember/object';
import $ from 'jquery';

export default Route.extend({

    healthy : true,

    beforeModel : function(transition) {
      // if(transition.targetName !== 'down') {
      //   return $.ajax({
      //     type : 'GET',
      //     url: 'https://teken.geenpeil.nl/cupido/health-check',
      //     success : function(d) {
      //       console.log('health check success:',d);
      //     }.bind(this),
      //     error : function() {
      //       // this.set('healthy',false);
      //       // this.transitionTo('down');
      //     }.bind(this)
      //   });
      // }
    },

    model : function() {
      return $.ajax({
        type : 'GET',
        url: 'form.json',
        error : function() {
          // TODO - unhandled
        }
      });
    },

    afterModel : function(model,transition) {
      switch(transition.targetName) {
        case('down') :
          if(this.get('healthy')) {
            this.transitionTo('home')
          }
        case('disclaimer') :
          return;
        default :
          this.transitionTo('home')
      }
    },

    setupController : function(controller,model) {
      if(typeof model === 'string') {
        model = JSON.parse(model);
      }

      controller.set('model',Object.create(model));
    }

});
