import Route from '@ember/routing/route';

export default Route.extend({

  //TODO handle social buttons

  actions : {

    back : function() {
      this.transitionTo('home');
    },
    share : function(network) {
      alert('TODO: '+network);
    }

  }

});
