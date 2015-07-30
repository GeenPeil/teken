#!/usr/bin/env bash

#############################
# Built for ubuntu/trusty64 #
#############################


# Update all the things first
sudo apt-get update

# Install g++ for Node compilation
sudo apt-get install -y g++

echo 'export PATH=$HOME/local/bin:$PATH' >> ~/.bashrc
. ~/.bashrc
mkdir ~/local
mkdir ~/local/bin
mkdir ~/node-latest-install
cd ~/node-latest-install
curl http://nodejs.org/dist/node-latest.tar.gz | tar xz --strip-components=1
./configure --prefix=$HOME/local
make install

# Reload
. ~/.profile

curl -L https://www.npmjs.org/install.sh | sh

# Reload
. ~/.profile

# Install Bower
npm install -g bower

# Install Ember-Cli
npm install -g ember-cli

# Setup Ember project
cd /vagrant/geenpeil-ember
npm install
bower install
ember build --environment production
