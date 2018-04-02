import Route from '@ember/routing/route';

export default Route.extend({

  actions : {

    done : function() {
      location.reload();
    }

  }

});
