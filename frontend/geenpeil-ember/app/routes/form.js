import Route from '@ember/routing/route';

export default Route.extend({

  setupController : function(controller,params) {
    controller.set('sectionNumber', parseInt(params.section) || 1);
  },

  activate : function() {
    window.onbeforeunload = function() { 
      return "Door te refreshen verliest u ingevulde gegevens op het formulier en moet u opnieuw beginnen."; 
    };
  },

  deactivate() {
    window.onbeforeunload = undefined;
  }

});
