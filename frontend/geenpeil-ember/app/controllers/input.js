import Ember from 'ember';

export default Ember.Controller.extend({

  /** DEPRECATED
  applicationController: Ember.inject.controller('application'),

  inputName : null,

  formItem : Ember.computed('inputName', function() {
    return this.get('applicationController').formItemForId(this.get('inputName'));
  }),

  inputNameChanged : Ember.observer('inputName', function() {
    //TODO - return previously entered data if available
    this.set('value',this.get('formItem.value') || undefined);
  }),

  valueChanged : Ember.observer('value', function() {
    //TODO - store the new value
    console.log('new value is',this.get('value'));
    this.set('formItem.value',this.get('value'));
  }),

  actions : {
    cancel : function() {
      var previous = this.get('applicationController').previousFormItem(this.get('formItem'));

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

      if(next) {
        this.transitionTo('input',next._id);
      }
      else {
        this.transitionToRoute('check');
      }
    }
  }
  **/

});
