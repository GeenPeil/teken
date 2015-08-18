import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

  /*
  * Section handling
  */

  section : Ember.computed('sectionNumber', function() {
    var sectionNumber = this.get('sectionNumber');
    return this.get('applicationController').getSection(sectionNumber-1);
  }),

  totalSections : Ember.computed('', function() {
    return this.get('applicationController.model.form.sections.length');
  }),

  formItems : Ember.computed('section', function() {
    return this.get('section').map(function(id) {
      return this.get('applicationController').formItemForId(id);
    }.bind(this));
  }),

  isFirstSection : Ember.computed('sectionNumber', function() {
    return this.get('sectionNumber') === 1;
  }),

  isLastSection : Ember.computed('sectionNumber', function() {
    return this.get('sectionNumber') === this.get('totalSections');
  }),

  /*
  * Input validation
  */

  noErrors : Ember.computed('formItems.@each.isValid', function() {
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
        alert('Niet alle velden zijn geldig. Kijk bij het veld wat er mis is met de invoer.');
      }

    }
  }

});
