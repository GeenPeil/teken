#!/usr/bin/env bash

vagrant ssh dev -c "cd /vagrant/geenpeil-ember && ember build -environment=production"
