import Component from '@ember/component';
import { computed } from '@ember/object';

export default Component.extend({

  classNames : ['generic-field'],

  isStringField : computed('formItem.type',function() {
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

  isSignatureField : computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  })

});
