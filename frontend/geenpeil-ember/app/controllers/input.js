import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  inputName : null,

  formItem : Ember.computed('inputName', function() {
    return this.get('applicationController').formItemForId(this.get('inputName'));
  }),

  actions : {
    cancel : function() {
      var previous = this.get('applicationController').previousFormItem(this.get('formItem'));
      console.log('previous',previous);

      if(previous) {
        this.transitionToRoute('input',previous._id);
      }
      else {
        //TODO - ask user for permission to dump data
        //TODO - dump data
        this.transitionToRoute('home');
      }
    },

    continue : function() {
      var next = this.get('applicationController').nextFormItem(this.get('formItem'));
      console.log('next',next);

      if(next) {
        this.transitionTo('input',next._id);
      }
      else {
        this.transitionToRoute('check');
      }
    }
  },

  isStringInput : Ember.computed('formItem', function() {
    return this.get('formItem.type') === 'string';
  }),

  isSignatureInput : Ember.computed('formItem', function() {
    return this.get('formItem.type') === 'signature';
  }),

});
