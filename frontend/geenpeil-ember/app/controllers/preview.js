import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  model : null,

  actions : {

    close : function() {
      window.history.back();
    }

  },

  properties : Ember.computed('model.form', function() {
    return this.get('model.form.properties');
  }),

  containerWidth : Ember.computed('model.form.properties.width', function() {
    return 'width: '+this.get('properties.width')+'px;';
  }),

  formImage : Ember.computed('',function() {
    return 'ext/'+this.get('applicationController').get('model.form.properties.filename');
  }),

  completedFormItems : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  })
});
