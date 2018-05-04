import Component from '@ember/component';
import { computed } from '@ember/object';
import { observer } from '@ember/object';
import { run } from '@ember/runloop';
import $ from 'jquery';

export default Component.extend({

  classNames : ['signature-field'],

  width : 360,

  height : 160,

  scale : 1.0,

  canvas : computed('showEditor', function() {
    return document.getElementById('can'); //FIXME - ugly
  }),

  ctx : computed('canvas', function() {
    return this.get('canvas').getContext("2d");
  }),

  showEditor : false,

  actions : {

    clicked : function() {
      if(this.get('formItem.isValid')) {
        if(this.erase()) {
          this.toggleProperty('showEditor');
        }
      }
      else {
        this.send('toggleEditor');
      }
    },

    toggleEditor : function() {
      this.toggleProperty('showEditor');
    },

    clear : function() {
      this.get('ctx').clearRect(0, 0, this.get('width'), this.get('height'));
    }

  },

  save : function() {
    var dataURL = this.get('canvas').toDataURL();

    // TODO - more better validation of image content
    // Checking for non falsey strings in response to some browser privacy features blocking toDataURL by returning a function
    if(typeof dataURL === 'string' && dataURL) {
      this.set('formItem.isValid',true);
      this.set('formItem.value',dataURL);
    }
  },

  erase : function() {
    var m = confirm("Wilt u de handtekening opnieuw zetten?");
    if(m) {
      this.set('formItem.isValid',undefined);
      this.set('formItem.value',undefined);
    }
    return m;
  },

  onValueChanged : observer('formItem.value', function() {
    var value = this.get('formItem.value');
    if(value) {
      var imageElement = this.$('.image-preview');
      if(imageElement) {
        imageElement.css('background-image',value);
      }
    }
  }),

  onShowEditor : observer('showEditor', function() {
    if(this.get('showEditor')) {
      window.scrollTo(0,0);
      run.next(this,function() {
        this.setupCanvas();
        $('body').addClass('noScroll');
        $('.signature-editor').bind('touchmove', function(e){e.preventDefault()});
      }.bind(this),1);
    }
    else {
      $('body').removeClass('noScroll');
      $('.signature-editor').unbind('touchmove');
    }
  }),

  setupCanvas : function() {
    var canvas = this.get('canvas');
    var ctx = this.get('ctx');

    // Measure the width of the screen
    var baseWidth = this.get('width');
    var maxScale  = 2.0;
    var actualWidth = $(window).width();

    var scale = (actualWidth / baseWidth) * 0.90;
    scale = scale > maxScale ? maxScale : scale;
    this.set('scale');
    $(canvas).css('transform', 'scale(' + scale + ')');

    // Apply existing image if found
    var imageUrl = this.get('formItem.value');
    if(imageUrl) {
      var img = new Image();
      img.onload = function(){
        ctx.drawImage(img,0,0);
      };
      img.src = imageUrl;
    }

    var flag = false,
        prevX = 0,
        currX = 0,
        prevY = 0,
        currY = 0,
        dot_flag = false;

    var thiz = this;

    // Setup mouse event listeners
    canvas.addEventListener("mousemove", function (e) {
        findxy('move', e);
    }, false);
    canvas.addEventListener("mousedown", function (e) {
        findxy('down', e);
    }, false);
    canvas.addEventListener("mouseup", function (e) {
        findxy('up', e);
    }, false);
    canvas.addEventListener("mouseout", function (e) {
        findxy('out', e);
    }, false);

    // Setup touch event listeners
    canvas.addEventListener("touchmove", function (e) {
        findxy('move', e.changedTouches[0]);
    }, false);
    canvas.addEventListener("touchstart", function (e) {
        findxy('down', e.changedTouches[0]);
    }, false);
    canvas.addEventListener("touchend", function (e) {
        findxy('up', e.changedTouches[0]);
    }, false);

    function draw() {
        ctx.beginPath();
        ctx.moveTo(prevX, prevY);
        ctx.lineTo(currX, currY);
        ctx.strokeStyle = "black";
        ctx.lineWidth = 2;
        ctx.stroke();
        ctx.closePath();
    }

    function findxy(res, e) {
        var eventX = e.clientX;
        var eventY = e.clientY;

        //scale the events
        var canvasCenterX = canvas.offsetLeft + (canvas.width / 2);
        eventX = canvasCenterX + ((eventX - canvasCenterX)/scale);
        eventY = canvas.offsetTop + ((eventY - canvas.offsetTop)/scale);

        if (res == 'down') {
            prevX = currX;
            prevY = currY;
            currX = eventX - canvas.offsetLeft;
            currY = eventY - canvas.offsetTop;

            flag = true;
            dot_flag = true;
            if (dot_flag) {
                ctx.beginPath();
                ctx.fillStyle = "black";
                ctx.fillRect(currX, currY, 2, 2);
                ctx.closePath();
                dot_flag = false;
            }
        }
        if (res == 'up') {
            flag = false;

            thiz.save();
        }
        if (res == 'out') {
          flag = false;
        }
        if (res == 'move') {
            if (flag) {
                prevX = currX;
                prevY = currY;
                currX = eventX - canvas.offsetLeft;
                currY = eventY - canvas.offsetTop;
                draw();
            }
        }

    }
  }

});
