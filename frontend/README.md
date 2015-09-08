GeenPeil Frontend
===

Getting Started
---

###Fast and easy way

1. Install Virtualbox
2. Install Vagrant
3. Download the base image here https://drive.google.com/open?id=0B3XfTcKvAVIQUmcwVXRZeURRUEU
4. Run 'vagrant up dev' from the frontend/ folder
5. Run 'bash scripts/develop.sh' start a development version on http://localhost:4200
6. Run 'bash scripts/build-production.sh' to make a production build into folder geenpeil-cordova/www/

 **Note:** the scripts can also be run with Batch on Windows if needed.


###From scratch

1. Install Virtualbox
2. Install Vagrant
3. Run 'vagrant up base'
4. Wait
5. Look inside scripts/ folder for useful build commands

 **Note:** update 'develop.sh' and 'build-production.sh' to say 'base' instead of 'dev' to use them in this case



###Roll your own

Install the following on your system:

1. Node ( https://nodejs.org/ )
2. Ember CLI ( http://www.ember-cli.com/user-guide/#getting-started )


Go into the folder geenpeil-ember/ and run:

1. 'npm install'
2. 'bower install'


Now you can run:

1. 'ember serve' to get a development build running on http://localhost:4200
2. 'ember build -environment=production' to make a production build into folder geenpeil-cordova/www/


**//TODO - Cordova and Android instructions**
