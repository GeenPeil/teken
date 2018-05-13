import Controller from '@ember/controller';
import { inject } from '@ember/controller';
import { computed } from '@ember/object';
import $ from 'jquery';
import { fixValue } from '../utils/form-fixes';

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
      this.set('isSending', true)

      var message = {};

      // Set all form items
      this.get('applicationController').get('model.form.fields').forEach(function(item) {
        message[item._id] = fixValue(item);
      });

      // Set ref
      message['ref'] = this.get('applicationController').get('ref');

      // Send the data
      var submitUrl = this.get('applicationController').get('model.form.properties.submitUrl')
      $.ajax({
        type : 'POST',
        url: submitUrl,
        data : JSON.stringify(message),
        contentType : 'application/json',
        error : function() {
          swal('Probleem', 'Er is een probleem opgetreden bij het versturen.', 'warning');
          this.set('isSending', false);
        }.bind(this),
        success : function(r) {
          var response = JSON.parse(r);

          if(response && response.success === true) {
            this.set('form.sent',true);
            this.transitionToRoute('complete');
          }
          else {

            if(response.error === 'mail has been used') {
              swal('Probleem', 'Dit e-mailadres is al gebruikt.', 'warning');
            }
            else {
              swal('Probleem', 'Er is een probleem opgetreden bij het versturen.', 'warning');
            }
            this.set('isSending', false);

          }

        }.bind(this),
      });

      //DEBUG
      // var data = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(message));
      // window.open(data,null);
    },

    sendNotVerified: function() {
      swal('Let op', 'U moet de beveilingsvraag beantwoorden voordat u het formulier kan versturen.', 'warning');
    }

  },

  isSending: false,

  readyToSend: computed('isSending', function() {
    return !this.get('isSending');
  })

});
