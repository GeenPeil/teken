import Controller from '@ember/controller';
import { inject } from '@ember/controller';
import { computed } from '@ember/object';
import $ from 'jquery';

export default Controller.extend({

  applicationController: inject('application'),

  form : computed('',function() {
    return this.get('applicationController').get('model.form');
  }),

  formItems : computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  }),

  allValid : computed('formItems.@each.isValid', function() {
      var formItems = this.get('formItems');
      var allValid = true;
      formItems.forEach(function(formItem) {
        if(formItem.type !== 'recaptcha' && formItem.isValid !== true) {
          allValid = false;
        }
      });
      return allValid;
  }),

  recaptchaItem : computed('formItems', function() {
    return this.get('formItems').findBy('_id','captchaResponse');
  }),

  actions : {

    send : function() {
      var message = {};

      this.get('applicationController').get('model.form.fields').forEach(function(item) {
        message[item._id] = item.value;
      });

      // Strip the signature so only base64 is left
      message['handtekening'] = message['handtekening'].replace(/^data:image\/(png|jpg);base64,/, "");

      // Send the data
      $.ajax({
        type : 'POST',
        url: 'https://teken.hartvoordemocratie.nl/cupido/submit',
        data : JSON.stringify(message),
        contentType : 'application/json',
        error : function() {
          alert('Er is een probleem opgetreden bij het versturen.');
        },
        success : function(r) {
          var response = JSON.parse(r);

          if(response && response.success === true) {
            this.set('form.sent',true);
            this.transitionToRoute('complete');
          }
          else {

            if(response.error === 'mail has been used') {
              alert('Dit e-mailadres is al gebruikt.');
            }
            else {
              alert('Er is een probleem opgetreden bij het versturen.');
            }

          }

        }.bind(this),
      });

      //DEBUG
      // var data = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(message));
      // window.open(data,null);
    }

  }

});
