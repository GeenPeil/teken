import Controller from '@ember/controller';
import Object from '@ember/object';
import { observer } from '@ember/object';

export default Controller.extend({

  model : null,

  onInit : function() {
    this.set('formData',Object.create());

    window.onbeforeunload = function() {
      if(!this.get('model.form.sent')) {
          return "Als u de pagina verlaat of herlaad worden de door u tot nu toe ingevulde gegevens gewist. Bevestig om nu te wissen.";
      }
    }.bind(this);
  },

  onModelChanged : observer('model', function() {
    var formMap = {};
    this.get('model.form.fields').forEach(function(field) {
      formMap[field._id] = field;
    });
    this.set('model.form.map',formMap);
  }),

  formItemForId : function(id) {
    return this.get('model.form.fields').findBy('_id',id);
  },

  getSection : function(sectionIndex) {
    return this.get('model.form.sections')[sectionIndex];
  },

  getSections : function() {
    return this.get('model.form.sections');
  }

});
