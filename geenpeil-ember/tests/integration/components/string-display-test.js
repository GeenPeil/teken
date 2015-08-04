import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';


moduleForComponent('display-string', 'Integration | Component | display string', {
  integration: true
});

test('it renders', function(assert) {
  assert.expect(2);

  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{display-string}}`);

  assert.equal(this.$().text(), '');

  // Template block usage:
  this.render(hbs`
    {{#display-string}}
      template block text
    {{/display-string}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
