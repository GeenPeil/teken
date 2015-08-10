import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['string-field'],

  classNameBindings: ['showError:error'],

  showError : Ember.computed('formItem.isValid','formItem.value.length', function() {
    var isValid = this.get('formItem.isValid'),
        value = this.get('formItem.value');

    return !isValid && value.length > 0;
  }),

  formItemChanged : Ember.observer('formItem', function() {
    this.set('value',this.get('formItem.value') || "");
  }),

  valueChanged : Ember.observer('value', function() {
    var value = this.get('value'),
        maxLength = this.get('formItem.length'),
        regex = new RegExp(this.get('formItem.regex'), "i"),
        isValid = false;

    //console.log('matched ',this.get('formItem._id'),value.match(regex));

    var match = !!value.match(regex);

    if(!match && value.length === 0) {
      isValid = undefined;
    }
    else {
      isValid = match && value.length <= maxLength;
    }

    this.set('formItem.isValid',isValid);

    //
    var separated = this.separateValue(this.get('value'));

    //set both values
    this.set('value',separated);
    this.set('formItem.value',separated.toUpperCase());
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
