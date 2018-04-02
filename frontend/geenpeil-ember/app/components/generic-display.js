import Component from '@ember/component';
import { computed } from '@ember/object';
import { htmlSafe } from '@ember/string';

export default Component.extend({

  attributeBindings: ['style'],

  classNames : ['generic-display'],

  isStringField : computed('formItem.type',function() {
    var type = this.get('formItem.type');

    switch(type) {

      case('string') :
      case('date') :
        return true;
      default :
        return false;

    }
  }),

  isSignatureField : computed('formItem.type',function() {
    return this.get('formItem.type') === 'signature';
  }),

  style : computed('formItem', function() {
    let x = this.get('formItem.x');
    let y = this.get('formItem.y');
    return htmlSafe(`top: ${y}px; left: ${x}px;`)
  })

});
