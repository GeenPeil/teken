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
      var message = this.get('controller.form.text.alert_before_refresh');
      return message;
    }.bind(this);
  },

  deactivate() {
    window.onbeforeunload = undefined;
  }

});
