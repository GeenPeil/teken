#!/usr/bin/env bash

#############################
# Built for ubuntu/trusty64 #
#############################


# Update all the things first
sudo apt-get update

# Install g++ for Node compilation
sudo apt-get install -y g++

# Install node
echo 'export PATH=/home/vagrant/local/bin:$PATH' >> ~/.profile
. ~/.profile
mkdir ~/local
mkdir ~/local/bin
mkdir ~/node-latest-install
cd ~/node-latest-install
curl http://nodejs.org/dist/node-latest.tar.gz | tar xz --strip-components=1
./configure --prefix=/home/vagrant/local
make install

# Install latest npm
curl -L https://www.npmjs.org/install.sh | sh

# Install Bower
npm install -g bower

# Install Ember-Cli
npm install -g ember-cli
