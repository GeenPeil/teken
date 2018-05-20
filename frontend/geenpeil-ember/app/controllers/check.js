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

      // Get alert texts
      var title = this.get('model.form.text.alert_title_problem');

      $.ajax({
        type : 'POST',
        url: submitUrl,
        data : JSON.stringify(message),
        contentType : 'application/json',
        error : function() {
          this.sendFailed();
        }.bind(this),
        success : function(r) {
          var response = JSON.parse(r);

          if(response && response.success === true) {
            this.set('form.sent',true);
            this.transitionToRoute('complete');
          }
          else {
            if(response.error === 'mail has been used') {
              this.sendEmailUsed();
            }
            else {
              this.sendFailed();
            }
          }

        }.bind(this),
        complete: function() {
          this.set('isSending', false);
        }
      });

      //DEBUG
      // var data = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(message));
      // window.open(data,null);
    },

    sendNotVerified: function() {
      var title = this.get('model.form.text.alert_title_generic');
      var text = this.get('model.form.text.check_must_verify');
      swal(title, text, 'warning');
    },

    sendFailed: function() {
      var title = this.get('model.form.text.alert_title_problem');
      var text = this.get('model.form.text.check_send_failed');
      swal(title, text, 'warning');
    },

    sendEmailUsed: function() {
      var title = this.get('model.form.text.alert_title_problem');
      var text = this.get('model.form.text.check_invalid_email');
      swal(title, text, 'warning');
    }

  },

  isSending: false,

  readyToSend: computed('isSending', function() {
    return !this.get('isSending');
  })

});
