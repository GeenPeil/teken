import Controller from '@ember/controller';
import Object from '@ember/object';
import { observer } from '@ember/object';

export default Controller.extend({

  model : null,

  queryParams: ['ref'],

  ref : '',

  onRefChanged : observer('ref', function() {
    // TODO - track ref
  }),

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
