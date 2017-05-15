# -*- mode: ruby -*-
# vi: set ft=ruby :

ENV["LC_ALL"] = "en_US.UTF-8"

Vagrant.configure("2") do |config|

  config.hostmanager.enabled = true
  config.hostmanager.ignore_private_ip = false
  config.hostmanager.include_offline = true

  config.vm.provider "virtualbox" do |v|
    v.customize ["modifyvm", :id, "--memory", "4096"]
    v.customize ["modifyvm", :id, "--cpus", "4"]
  end

  config.vm.box = "boxcutter/centos73"
  config.vm.hostname = "foremanserver.lab.local.dev"
  config.vm.network :private_network, ip: "10.0.20.10"
  config.hostmanager.aliases = %w(foremanserver)

  config.vm.provision "shell", inline: <<-SHELL
    yum -y install https://yum.puppetlabs.com/puppetlabs-release-pc1-el-7.noarch.rpm
    yum -y install epel-release
    yum -y install https://yum.theforeman.org/releases/1.14/el7/x86_64/foreman-release.rpm
    yum -y install foreman-installer
    systemctl restart network
    foreman-installer --foreman-admin-password azerty >> foreman-installer.log
  SHELL

end