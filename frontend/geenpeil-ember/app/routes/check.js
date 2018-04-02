import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

      back : function() {
        window.history.back();
      },

      preview : function() {
        this.transitionTo('preview');
      }

  }

});
