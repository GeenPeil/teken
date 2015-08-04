import DS from 'ember-data';

export default DS.Model.extend({

  timestamp : DS.attr('string'),

  complete : DS.attr('boolean'),

  acceptedTos : DS.attr('boolean'),

  voornaam : DS.attr('string'),

  tussenvoegsel : DS.attr('string'),

  achternaam : DS.attr('string'),

  straat : DS.attr('string'),

  huisnummer : DS.attr('string'),

  postcode : DS.attr('string'),

  woonplaats : DS.attr('string'),

  geboortedatum : DS.attr('string'),

  geboorteplaats : DS.attr('string'),

  handtekening : DS.attr('string')

});
