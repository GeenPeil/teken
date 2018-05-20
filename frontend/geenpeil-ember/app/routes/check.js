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
    var message = this.get('controller.model.form.text.alert_before_refresh');
    window.onbeforeunload = function() { 
      return message;
    };
  },

  deactivate() {
    window.onbeforeunload = undefined;
  }

});
