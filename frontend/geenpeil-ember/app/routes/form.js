import Route from '@ember/routing/route';

export default Route.extend({

  setupController : function(controller,params) {
    controller.set('sectionNumber', parseInt(params.section) || 1);
  },

  activate : function() {
    window.scrollTo(0,0);
    window.onbeforeunload = function() { 
      var message = this.get('controller.form.text.alert_before_refresh');
      return message;
    }.bind(this);
  },

  deactivate() {
    window.onbeforeunload = undefined;
  }

});
