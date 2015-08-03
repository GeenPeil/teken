## pechtold - teken.geenpeil.nl server application

### API

pechtold draait een HTTP API waarmee nieuwe handtekeningen kunnen worden opgeslagen op de server, om op een later tijdstip uitgeprint te worden.

Voor het opslaan van de gegevens is slechts een simpele HTTP POST call nodig, met in de body JSON data. De JSON data is een object met de volgende velden:

 - `voornaam` (string) - "eerste officiele voornaam"
 - `tussenvoegsel` (string)
 - `achternaam` (string)
 - `geboortedatum` (string) - volgt formaat "dd-mm-yyyy"
 - `geboorteplaats` (string)
 - `straat` (string)
 - `huisnummer` (string) - huisnummer inclusief toevoeging
 - `postcode` (string)
 - `woonplaats` (string)
 - `handtekening` (string) - base64-encoded jpg of png (nog af te spreken)

TODO: captcha in de API meenemen

### Encryptie van handtekeningen en n.a.w. gegevens

TODO: encryptie specificeren

### Opslag van encrypted gegevens

TODO: opslag folder structuur specificeren
