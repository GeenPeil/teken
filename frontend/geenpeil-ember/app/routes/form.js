import Route from '@ember/routing/route';

export default Route.extend({

  setupController : function(controller,params) {
    controller.set('sectionNumber', parseInt(params.section) || 1);
  },

  activate : function() {
    window.scrollTo(0,0);
    var message = this.get('controller.model.form.text.alert_before_refresh');
    window.onbeforeunload = function() { 
      return message;
    };
  },

  deactivate() {
    window.onbeforeunload = undefined;
  }

});
