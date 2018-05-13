
export function fixValue(formItem) {

  // Only work on validated values
  if(!formItem.isValid) {
    return formItem.value;
  }

  switch(formItem._id) {
    case('postcode') :
      return fixPostcode(formItem);
    case('geboortedatum') :
      return fixDate(formItem);
    case('handtekening') :
      return fixImage(formItem);
    default :
      return formItem.value;
  }

}

export function fixPostcode(formItem) {
  var value = formItem.value;
  value = value.replace(/ /g,''); // remove spaces
  value = value.substring(0,4) + ' ' + value.substring(4);
  return value;
}

export function fixDate(formItem) {
  formItem.value;
  var parts = formItem.value.split(formItem.separator);
  var partLengths = formItem.display.split(' ').map(function(string) { return string.length; });
  partLengths.forEach(function(length, index) {
    var diff = length - parts[index].length;
    var pad = '';
    while(diff) {
      pad += '0';
      diff--;
    }
    parts[index] = pad + parts[index];
  });
  return parts.join(formItem.separator);
}

export function fixImage(formItem) {
  return formItem.value.replace(/^data:image\/(png|jpg);base64,/, '');
}
