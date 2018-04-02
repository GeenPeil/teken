#!/usr/bin/env bash

#############################
# Built for ubuntu/xenial64 #
#############################


# Update all the things first
sudo apt-get update

# Install node
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install Ember-Cli
sudo npm install -g ember-cli

# Install dependencies
cd /vagrant/geenpeil-ember && npm install
