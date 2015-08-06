import Ember from 'ember';

export default Ember.Component.extend({

  isStringField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'string';
  }),

  isSignatureField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  })

});
