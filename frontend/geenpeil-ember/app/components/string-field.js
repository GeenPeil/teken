import Ember from 'ember';

export default Ember.Component.extend({

  tagName : 'input',

  type : 'text',

  classNames : ['string-field']

  //TODO - based on the pattern render one or more textfields
  //TODO - check the text input with the regex (probably not so simple)
  //TODO - do some good length checks
  //TODO - save the string data in the form model

});
