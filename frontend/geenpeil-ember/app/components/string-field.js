import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['string-field'],

  classNameBindings: ['showError:error','formItem.case-sensitive:case-sensitive'],

  inputType : Ember.computed('formItem.type', function() {
    var type = this.get('formItem.type');

    switch(type) {

      case('email') :
        return 'email';
        break;
      case('string') :
      case('date') :
      default :
        return 'text';
    }
  }),

  showLengthWarning : false,

  showError : Ember.computed('formItem.isValid','formItem.value.length', function() {
    var isValid = this.get('formItem.isValid'),
        value = this.get('formItem.value');

    return isValid === undefined ? false : !isValid;
  }),

  formItemChanged : Ember.observer('formItem', function() {
    this.set('value',this.get('formItem.value') || "");
  }),

  valueChanged : Ember.observer('value', function() {
    var value = this.get('value'),
        maxLength = this.get('formItem.length'),
        caseSensitive = this.get('formItem.case-sensitive'),
        regex = new RegExp(this.get('formItem.regex'), "i"),
        isValid = false;

    //console.log('matched ',this.get('formItem._id'),value.match(regex));

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

    // for fields that have 'display' (FIXME),
    // check that input length matches to be valid
    if(this.get('formItem.display')) {
      isValid = maxLengthReached ? isValid : undefined;
      this.set('showLengthWarning',false);
    }


    // fetch value
    var tmp = this.get('value');

    // auto replace separators
    // tmp = this.separateValue(this.get('value'));

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

    console.log('isValid',isValid);

    //set all values
    this.set('formItem.isValid',isValid);
    this.set('value',tmp);
    this.set('formItem.value',tmp);
  }),

  separateValue : function(s) {
    var display = this.get('formItem.display');

    if(display) {
      var separator = this.get('formItem.separator');
      var parts = display.split(' ');
      var index = 0;
      for(var i=0;i<parts.length;i++) {
        index += parts[i].length;

        if(index < s.length) {
          s = s.substr(0, index) + separator + s.substr(index+separator.length);
        }

        index += 1;
      }
    }
    return s;
  }

});
