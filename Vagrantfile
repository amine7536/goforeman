# -*- mode: ruby -*-
# vi: set ft=ruby :

ENV["LC_ALL"] = "en_US.UTF-8"

Vagrant.configure("2") do |config|

  config.vm.provider "virtualbox" do |v|
    v.customize ["modifyvm", :id, "--memory", "4096"]
    v.customize ["modifyvm", :id, "--cpus", "4"]
  end

  config.vm.box = "centos/7"
  config.vm.hostname = "foremanserver.lab.local.dev"
  config.vm.network :private_network, ip: "10.0.20.10"

  config.vm.provision "shell", inline: <<-SHELL
    yum -y install https://yum.puppetlabs.com/puppet5/puppet5-release-el-7.noarch.rpm
    yum -y install http://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
    yum -y install https://yum.theforeman.org/releases/1.17/el7/x86_64/foreman-release.rpm
    yum -y install foreman-installer
    foreman-installer --foreman-admin-password azerty >> foreman-installer.log
  SHELL

end
