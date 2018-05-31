import Component from '@ember/component';
import { computed } from '@ember/object';

export default Component.extend({

  recaptchaElementId : computed('formItem',function() {
    return 'recaptcha-' + this.get('elementId');
  }),

  didInsertElement : function() {
    if(!window.grecaptcha) {
      var script = document.createElement('script');
      window.onRecaptchaLoaded = function() {
        this.renderCaptcha();
      }.bind(this);
      var lang = this.get('formItem.lang');
      script.src = 'https://www.google.com/recaptcha/api.js?onload=onRecaptchaLoaded&render=explicit&hl='+lang;
      document.getElementsByTagName('head')[0].appendChild(script);
    }
    else {
      this.renderCaptcha();
    }
  },

  renderCaptcha : function() {   
    window.grecaptcha.render(this.get('recaptchaElementId'), {
      'sitekey' : this.get('formItem.sitekey'),
      'callback' : function(v) {
        this.set('formItem.value',v);
        this.set('formItem.isValid',true);
      }.bind(this),
      'expired-callback' : function() {
        this.set('formItem.value', undefined);
        this.set('formItem.isValid', undefined);
      }.bind(this),
      'theme' : this.get('formItem.theme')
    });
  }

});
