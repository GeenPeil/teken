import Ember from 'ember';

export default Ember.Controller.extend({

  model : null,

  formData : null,

  onInit : Ember.on('init', function() {
    this.set('formData',Ember.Object.create());

    window.onbeforeunload = function() {
      return "Als u de pagina verlaat of herlaad worden de door u tot nu toe ingevulde gegevens gewist. Bevestig om nu te wissen.";
    };
  }),

  onModelChanged : Ember.observer('model', function() {
    var formMap = {};
    this.get('model.form.fields').forEach(function(field) {
      formMap[field._id] = field;
    });
    this.set('model.form.map',formMap);
  }),

  formItemForId : function(id) {
    return this.get('model.form.fields').findBy('_id',id);
  },

  /** DEPRECATED
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
  },
  */

  getSection : function(sectionIndex) {
    return this.get('model.form.sections')[sectionIndex];
  },

  getSections : function() {
    return this.get('model.form.sections');
  }

});
