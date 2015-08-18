import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  form : Ember.computed('', function() {
    return this.get('applicationController').get('model.form');
  }),

  actions : {

    share : function() {
      this.transitionToRoute('share');
    },

    form : function() {
      this.transitionToRoute('disclaimer');
    }

  }

});
