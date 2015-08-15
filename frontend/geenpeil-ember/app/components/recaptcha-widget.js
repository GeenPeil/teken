import Ember from 'ember';

export default Ember.Component.extend({

  recaptchaElementId : Ember.computed('formItem',function() {
    return 'recaptcha-' + this.get('elementId');
  }),

  onDidInsertElement : Ember.on('didInsertElement', function() {
    window.grecaptcha.render(this.get('recaptchaElementId'), {
        'sitekey' : this.get('formItem.sitekey'),
        'callback' : function(v) {
          this.set('formItem.value',v);
          this.set('formItem.isValid',true);
        }.bind(this),
        'theme' : this.get('formItem.theme')
    });
  })

});
