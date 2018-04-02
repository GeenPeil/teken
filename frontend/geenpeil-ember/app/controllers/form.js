import Controller from '@ember/controller';
import { inject } from '@ember/controller';
import { computed } from '@ember/object';

export default Controller.extend({

  applicationController: inject('application'),

  /*
  * Section handling
  */

  section : computed('sectionNumber', function() {
    var sectionNumber = this.get('sectionNumber');
    return this.get('applicationController').getSection(sectionNumber-1);
  }),

  totalSections : computed('', function() {
    return this.get('applicationController.model.form.sections.length');
  }),

  preferences : computed('model.form', function() {
    return this.get('applicationController.model.form.preferences');
  }),

  formItems : computed('section', function() {
    return this.get('section').map(function(id) {
      return this.get('applicationController').formItemForId(id);
    }.bind(this));
  }),

  isFirstSection : computed('sectionNumber', function() {
    return this.get('sectionNumber') === 1;
  }),

  isLastSection : computed('sectionNumber', function() {
    return this.get('sectionNumber') === this.get('totalSections');
  }),

  /*
  * Input validation
  */

  noErrors : computed('formItems.@each.isValid', function() {
    var stopOnEmpty = this.get('applicationController.model.form.preferences.stopOnEmpty');

    if(stopOnEmpty) {
      return this.get('formItems').findBy('isValid',false) === undefined && this.get('formItems').findBy('isValid',undefined) === undefined;
    }
    else {
      return !!this.get('formItems').findBy('isValid',false) === false;
    }

  }),

  /*
  * Actions
  */

  actions : {
    cancel : function() {
      var notFirst = !this.get('isFirstSection');

      if(notFirst) {
        this.transitionToRoute('form',this.get('sectionNumber')-1);
      }
      else {
        location.reload();
      }
    },

    continue : function() {
      if(this.get('noErrors')) {
        var notLast = !this.get('isLastSection');

        if(notLast) {
          this.transitionToRoute('form',this.get('sectionNumber')+1);
        }
        else {
          this.transitionToRoute('check');
        }
      }
      else {
        alert('U moet alle velden geldig ingevuld hebben om verder te kunnen gaan.');
      }

    }
  }

});
