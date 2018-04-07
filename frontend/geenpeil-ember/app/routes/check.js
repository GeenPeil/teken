import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

      back : function() {
        window.history.back();
      },

      preview : function() {
        this.transitionTo('preview');
      }

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
