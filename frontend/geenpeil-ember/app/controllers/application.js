import Ember from 'ember';

export default Ember.Controller.extend({

  model : null,

  formData : null,

  onInit : Ember.on('init', function() {
    this.set('formData',Ember.Object.create());
  }),

  formItemForId : function(id) {
    return this.get('model.form.fields').findBy('_id',id);
  },

  formItemIndex : function(id) {
    return this.get('model.form.fields').indexOf(id);
  },

  nextFormItem : function(formItem) {
    var formItemIndex = this.formItemIndex(formItem);
    if(formItemIndex >= this.get('model.form.fields.length')) {
      return null;
    }
    else {
      return this.get('model.form.fields')[formItemIndex+1];
    }
  },

  previousFormItem : function(formItem) {
    var formItemIndex = this.formItemIndex(formItem);
    if(formItemIndex <= 0) {
      return null;
    }
    else {
      return this.get('model.form.fields')[formItemIndex-1];
    }
  }

});
