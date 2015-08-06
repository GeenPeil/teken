import Ember from 'ember';

export default Ember.Component.extend({

  classNames : ['signature-field'],

  width : 360,

  height : 160,

  canvas : Ember.computed('', function() {
    return document.getElementById('can');
  }),

  ctx : Ember.computed('', function() {
    return this.get('canvas').getContext("2d");
  }),

  actions : {

    erase : function() {
              console.log('context',this.get('ctx'));
        var m = confirm("Wilt u de handtekening opnieuw zetten?");
        if (m) {
            this.get('ctx').clearRect(0, 0, this.get('width'), this.get('height'));
            document.getElementById("canvasimg").style.display = "none";
        }
    }

  },

  save : function() {
    console.log('save');
    var dataURL = this.get('canvas').toDataURL();
    this.set('value',dataURL);
  },

  onDidInsertElement : Ember.on('didInsertElement', function() {
    var canvas = this.get('canvas');
    var ctx = this.get('ctx');
    var w = canvas.width;
    var h = canvas.height;

    // Apply existing image if found
    var imageUrl = this.get('value');
    if(imageUrl) {
      var img = new Image;
      img.onload = function(){
        ctx.drawImage(img,0,0); // Or at whatever offset you like
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
        findxy('move', e)
    }, false);
    canvas.addEventListener("mousedown", function (e) {
        findxy('down', e)
    }, false);
    canvas.addEventListener("mouseup", function (e) {
        findxy('up', e)
    }, false);
    canvas.addEventListener("mouseout", function (e) {
        findxy('out', e)
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
      // console.log('findXy',res,e);
        if (res == 'down') {
            prevX = currX;
            prevY = currY;
            currX = e.clientX - canvas.offsetLeft;
            currY = e.clientY - canvas.offsetTop;

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
        if (res == 'up' || res == "out") {
            flag = false;

            thiz.save();

        }
        if (res == 'move') {
            if (flag) {
                prevX = currX;
                prevY = currY;
                currX = e.clientX - canvas.offsetLeft;
                currY = e.clientY - canvas.offsetTop;
                draw();
            }
        }
    }
  })

});
