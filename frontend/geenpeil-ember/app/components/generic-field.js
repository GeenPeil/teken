import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['generic-field'],

  isStringField : Ember.computed('formItem.type',function() {
    var type = this.get('formItem.type');

    switch(type) {

      case('string') :
      case('email') :
      case('date') :
        return true;
      default :
        return false;

    }
  }),

  isSignatureField : Ember.computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  })

});
