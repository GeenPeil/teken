import Route from '@ember/routing/route';

export default Route.extend({

  setupController : function(controller,params) {
    controller.set('sectionNumber', parseInt(params.section) || 1);
  }

});
