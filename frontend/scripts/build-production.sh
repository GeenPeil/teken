#!/usr/bin/env bash

vagrant ssh -c "cd /vagrant/geenpeil-ember && ember build --environment production"
