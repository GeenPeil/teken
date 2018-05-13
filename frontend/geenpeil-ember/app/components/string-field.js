import Component from '@ember/component';
import { computed } from '@ember/object';
import { observer } from '@ember/object';

export default Component.extend({

  classNames : ['string-field'],

  classNameBindings: ['showError:error','formItem.case-sensitive:case-sensitive'],

  inputType : computed('formItem.type', function() {
    var type = this.get('formItem.type');

    switch(type) {

      case('email') :
        return 'email';
      case('string') :
      case('date') :
      default :
        return 'text';
    }
  }),

  showLengthWarning : false,

  showError : computed('formItem.{isValid,value.length}', function() {
    var isValid = this.get('formItem.isValid')

    return isValid === undefined ? false : !isValid;
  }),

  didInsertElement: function() {
    this.recalculateValues();
  },

  valueChanged : observer('formItem.value', function() {
    this.recalculateValues();
  }),

  recalculateValues: function() {
    var value = this.get('formItem.value') || "",
        minLength = this.get('formItem.minLength') || 0,
        maxLength = this.get('formItem.length'),
        caseSensitive = this.get('formItem.case-sensitive'),
        regex = new RegExp(this.get('formItem.regex'), "i"),
        isValid = false;

    // check the value against the regex
    var match = !!value.match(regex);

    // if the regex does not match because there was no input
    if(!match && value.length === 0) {
      // validity is kept or set as undefined
      isValid = undefined;
    }
    else {
      // check that the length does not exceed the maximum
      isValid = match && value.length <= maxLength;
    }

    // check if the value equals the maximum
    var maxLengthReached = value.length === maxLength;
    this.set('showLengthWarning',maxLengthReached);

    // for fields that have 'display' (FIXME) we want to postpone error messages
    // until a certain number of characters are entered
    if(this.get('formItem.minLength')) {
      isValid = maxLengthReached || value.length >= minLength ? isValid : undefined;
      this.set('showLengthWarning',false);
    }

    // fetch value
    var tmp = value;

    // only upper case the value if it is not case sensitive
    if(!caseSensitive) {
      tmp = tmp.toUpperCase();
    }

    //before date check
    if(this.get('formItem.type') === 'date') {
      var beforeDate = this.get('formItem.beforeDate'),
          separator = this.get('formItem.separator');
      if(beforeDate) {

        //chop em up
        var inputDate = parseInt(tmp.split(separator).reverse().join(''));
        var testDate = parseInt(beforeDate.split(separator).reverse().join(''));

        if(inputDate > testDate) {
          isValid = false;
        }

      }
    }

    //set all values
    this.set('formItem.isValid',isValid);
    this.set('formItem.value',tmp);
  }

});
