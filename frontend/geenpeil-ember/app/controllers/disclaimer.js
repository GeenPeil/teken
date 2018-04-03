import Controller from '@ember/controller';
import { inject } from '@ember/controller';
import { computed } from '@ember/object';

export default Controller.extend({

  applicationController: inject('application'),

  form : computed('', function() {
    return this.get('applicationController').get('model.form');
  }),

  consent : false

});
