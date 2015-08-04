import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';


moduleForComponent('generic-field', 'Integration | Component | generic field', {
  integration: true
});

test('it renders', function(assert) {
  assert.expect(2);

  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{generic-field}}`);

  assert.equal(this.$().text(), '');

  // Template block usage:
  this.render(hbs`
    {{#generic-field}}
      template block text
    {{/generic-field}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
