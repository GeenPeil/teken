import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['generic-field'],

  isStringField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'string';
  }),

  isSignatureField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  })

});
