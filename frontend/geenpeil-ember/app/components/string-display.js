import Component from '@ember/component';
import { computed } from '@ember/object';
import { htmlSafe } from '@ember/string';

export default Component.extend({

  classNames : ['string-display'],

  parts : computed('forItem',function() {
    var value = this.get('formItem.value'),
        display = this.get('formItem.display'),
        properties = this.get('properties');

    if(display) {
      var valueParts = [];
      var parts = display.split(' ');
      var endIndex = 0;
      var startIndex = 0;
      for(var i=0;i<parts.length;i++) {
        endIndex += parts[i].length;

        var s = value.substr(startIndex,endIndex-startIndex);

        var xOffset = startIndex*properties.characterWidth;

        valueParts.push({
          value : s,
          style : htmlSafe(`position:absolute; left:${xOffset}px;`)
        });

        endIndex += 1;
        startIndex = endIndex;
      }

      return valueParts;
    }
    // single line
    else {
      return [{value:value,style:""}];
    }

  })

});
