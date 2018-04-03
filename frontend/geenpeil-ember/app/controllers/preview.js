import Controller from '@ember/controller';
import { inject } from '@ember/controller';
import { computed } from '@ember/object';

export default Controller.extend({

  applicationController: inject('application'),

  model : null,

  actions : {

    close : function() {
      window.history.back();
    }

  },

  properties : computed('model.form', function() {
    return this.get('model.form.properties');
  }),

  containerWidth : computed('model.form.properties.width', function() {
    return 'width: '+this.get('properties.width')+'px;';
  }),

  completedFormItems : computed('',function() {
    return this.get('applicationController').get('model.form.fields');
  })
});
