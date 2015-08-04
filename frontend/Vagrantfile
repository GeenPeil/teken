# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  # Ember port
  config.vm.network :forwarded_port, host: 4200, guest: 4200

  # Ember live reload port
  config.vm.network :forwarded_port, host: 35729, guest: 35729

  # Configure VM hardware usage
  config.vm.provider "virtualbox" do |v|
    v.memory = 512
    v.cpus = 1
  end

  # The base box with all the required software installed
  config.vm.define "base" do |base|

    # Specify the Vagrant box to build on
    base.vm.box = "ubuntu/trusty64"

    # Specify bootstrapping
    base.vm.provision :shell, path: "bootstrap.sh", privileged: false

  end

end
