import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  model : null,

  formImage : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.properties.filename');
  }),

  completedFormItems : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  })
});
