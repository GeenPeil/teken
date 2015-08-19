import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  form : Ember.computed('',function() {
    return this.get('applicationController').get('model.form');
  }),

  formItems : Ember.computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  }),

  allValid : Ember.computed('formItems.@each.isValid', function() {
      var formItems = this.get('formItems');
      var allValid = true;
      formItems.forEach(function(formItem) {
        if(formItem.type !== 'recaptcha' && formItem.isValid !== true) {
          allValid = false;
        }
      });
      return allValid;
  }),

  recaptchaItem : Ember.computed('formItems', function() {
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
      Ember.$.ajax({
        type : 'POST',
        url: 'https://teken.geenpeil.nl/pechtold/submit',
        data : JSON.stringify(message),
        contentType : 'application/json',
        error : function(e) {
          console.error('error:',e);
          alert('Er is een probleem opgetreden bij het versturen.');
        },
        success : function(r) {
          console.log('response:',r);

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
