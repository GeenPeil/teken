import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['string-field'],

  formItemChanged : Ember.observer('formItem', function() {
    console.log('formItemChanged');
    this.set('value',this.get('formItem.value'));
  }),

  valueChanged : Ember.observer('value', function() {

    //check length
    this.set('formItem.value',this.get('value'));
  })

  //TODO - based on the pattern render one or more textfields
  //TODO - check the text input with the regex (probably not so simple)
  //TODO - do some good length checks
  //TODO - save the string data in the form model

});
