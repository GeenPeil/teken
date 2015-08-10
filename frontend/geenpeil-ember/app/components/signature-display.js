import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['signature-display'],

  tagName : 'div',

  onDidInsertElement : Ember.on('didInsertElement', function() {
    this.$().css({ width: this.get('formItem.width') });
    this.$().css({ height: this.get('formItem.height') });
    this.$().css( 'background-image','url(' + this.get('formItem.value') + ')' );
  })

});
