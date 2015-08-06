import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

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

  completedFormItems : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  })

});
