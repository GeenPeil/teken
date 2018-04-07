import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

    accept : function() {
      this.transitionTo('form',1);
    },

    acceptWithoutConsent : function() {
      alert('U moet akkoord gaan met de voorwaarden om het formulier in te vullen.');
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
