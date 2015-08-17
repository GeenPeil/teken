import Ember from 'ember';

export default Ember.Component.extend({

  recaptchaElementId : Ember.computed('formItem',function() {
    return 'recaptcha-' + this.get('elementId');
  }),

  onDidInsertElement : Ember.on('didInsertElement', function() {
    window.grecaptcha.render(this.get('recaptchaElementId'), {
        'sitekey' : this.get('formItem.sitekey'),
        'callback' : function(v) {
          console.log('callback',v);
          this.set('formItem.value',v);
          this.set('formItem.isValid',true);
        }.bind(this),
        'expired-callback' : function(v) {
          console.log('expired callback',v);
          this.set('formItem.value', undefined);
          this.set('formItem.isValid', undefined);
        }.bind(this),
        'theme' : this.get('formItem.theme')
    });
  })

});
