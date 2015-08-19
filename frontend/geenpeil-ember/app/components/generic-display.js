import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['generic-display'],

  isStringField : Ember.computed('formItem.type',function() {
    var type = this.get('formItem.type');

    switch(type) {

      case('string') :
      case('date') :
        return true;
      default :
        return false;

    }
  }),

  isSignatureField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  }),

  onDidInsertElement : Ember.on('didInsertElement', function() {
    this.$().css({ top: this.get('formItem.y') });
    this.$().css({ left: this.get('formItem.x') });
  })

});
