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
    console.log('formItemChanged');
    this.set('value',this.get('formItem.value') || "");
  }),

  valueChanged : Ember.observer('value', function() {
    var value = this.get('value'),
        maxLength = this.get('formItem.length'),
        regex = new RegExp(this.get('formItem.regex'), "i"),
        isValid = false;

    console.log('matched ',this.get('formItem._id'),value.match(regex));

    var match = !!value.match(regex);

    if(!match && value.length === 0) {
      isValid = undefined;
    }
    else {
      isValid = match && value.length <= maxLength;
    }

    console.log(this.get('formItem._id'),isValid);

    this.set('formItem.isValid',isValid);

    //store the value
    this.set('formItem.value',this.get('value'));
  })

  //TODO - based on the pattern render one or more textfields
  //TODO - check the text input with the regex (probably not so simple)
  //TODO - do some good length checks
  //TODO - save the string data in the form model

});
