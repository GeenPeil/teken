import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['generic-display'],

  isStringField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'string';
  }),

  isSignatureField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  }),

  onDidInsertElement : Ember.on('didInsertElement', function() {
    this.$().css({ top: this.get('formItem.y') });
    this.$().css({ left: this.get('formItem.x') });
  })

});
