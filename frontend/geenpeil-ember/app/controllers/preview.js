import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  model : null,

  formImage : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.properties.filename');
  }),

  completedFormItems : Ember.computed('',function() {
    var model = this.get('model'),
        formItems = this.get('applicationController').get('model.form.fields');

    //match each form item with one of the model values
    return formItems.map(function(formItem) {
      formItem.value = model[formItem._id];
      return formItem;
    });
  })
});
