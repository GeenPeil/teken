import Ember from 'ember';

export default Ember.Controller.extend({

  applicationController: Ember.inject.controller('application'),

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

  actions : {
    cancel : function() {
      var notFirst = !this.get('isFirstSection');

      console.log('cancel',notFirst);

      if(notFirst) {
        this.transitionToRoute('form',this.get('sectionNumber')-1);
      }
      else {
        //TODO - ask user for permission to dump data
        //TODO - dump data
        this.transitionToRoute('home');
      }
    },

    continue : function() {
      var notLast = !this.get('isLastSection');

      if(notLast) {
        this.transitionTo('form',this.get('sectionNumber')+1);
      }
      else {
        this.transitionToRoute('check');
      }
    }
  }

});
