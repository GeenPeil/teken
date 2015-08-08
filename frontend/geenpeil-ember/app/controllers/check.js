import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  formItems : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  }),

  allValid : Ember.computed('formItems.@each.isValid', function() {
      var formItems = this.get('formItems');
      var allValid = true;
      formItems.forEach(function(formItem) {
        if(formItem.isValid !== true) {
          allValid = false;
        }
      });
      return allValid;
  }),

  actions : {

    send : function() {
      //TODO - check captcha

      var obj = this.get('applicationController').get('model.form.fields').map(function(formItem) {
        return {
          _id : formItem._id,
          value : formItem.value
        };
      });

      var data = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(obj));

      window.open(data,null);
    }
  },



});
