import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

    accept : function() {
      this.transitionTo('form',1);
    },

    decline : function() {
      this.transitionTo('home');
    }

  }

});
