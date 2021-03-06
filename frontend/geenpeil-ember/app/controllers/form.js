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

  form : computed('',function() {
    return this.get('applicationController').get('model.form');
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

  errors : computed('formItems.@each.isValid', function() {

    // Being careful with return types which are either Ember.Array or undefined

    var errors = [];

    var inValidItems = this.get('formItems').filterBy('isValid',false);
    var undefinedItems = this.get('formItems').filterBy('isValid',undefined);

    if(inValidItems) {
      errors = errors.concat(inValidItems);
    }
    if(undefinedItems) {
      errors = errors.concat(undefinedItems);
    }

    return errors
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
        var errorFieldNames = this.get('errors').map(function(error) { return error.name })

        // Log error with Piwik
        if(window._paq != undefined) {
          window._paq.push(['trackEvent', 'Form', this.get('sectionNumber'), 'Error', 'Invalid ' + errorFieldNames.join(', ')]);
        }

        var title = this.get('form.text.alert_title_generic');
        var textSingle = this.get('form.text.form_single_invalid');
        var textMultiple = this.get('form.text.form_multiple_invalid');

        if(errorFieldNames.length < 2) {
          swal(title, textSingle + errorFieldNames, 'warning');
        }
        else {
          swal(title, textMultiple + errorFieldNames.join(', ') + '.', 'warning');
        }
      }

    }
  }

});
