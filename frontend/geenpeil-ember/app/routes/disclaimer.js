import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

    accept : function() {
      this.transitionTo('form',1);
    },

    acceptWithoutConsent : function() {
      swal ('Let op', 'U moet akkoord gaan met de voorwaarden om het formulier in te vullen.', 'error');
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
