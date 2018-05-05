import Route from '@ember/routing/route';

export default Route.extend({

  activate: function() {
    window.scrollTo(0,0);
  },

  actions : {

    done : function() {
      location.reload();
    }

  }

});
