import Component from '@ember/component';
import { computed } from '@ember/object';
import { htmlSafe } from '@ember/string';

export default Component.extend({

  classNames : ['signature-display'],

  tagName : 'div',

  attributeBindings: ['style'],

  style : computed('formItem', function() {
    let width = this.get('formItem.width');
    let height = this.get('formItem.height');
    let backgroundImage = this.get('formItem.value');
    return htmlSafe(`width: ${width}px; height: ${height}px; background-image: url(${backgroundImage});`)
  })

});
