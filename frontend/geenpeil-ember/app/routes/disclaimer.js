import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

    accept : function() {
      this.transitionTo('form',1);
    },

    acceptWithoutConsent : function() {
      var title = this.get('controller.model.form.text.alert_title_generic');
      var text = this.get('controller.model.form.text.disclaimer_must_accept');
      swal (title, text, 'error');
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
