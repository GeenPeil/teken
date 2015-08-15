import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';


moduleForComponent('recaptcha-widget', 'Integration | Component | recaptcha widget', {
  integration: true
});

test('it renders', function(assert) {
  assert.expect(2);

  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{recaptcha-widget}}`);

  assert.equal(this.$().text(), '');

  // Template block usage:
  this.render(hbs`
    {{#recaptcha-widget}}
      template block text
    {{/recaptcha-widget}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
