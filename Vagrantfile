# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

  config.vm.box = "bento/ubuntu-16.04"
  config.vm.network "forwarded_port", guest: 9999, host: 9999, host_ip: "127.0.0.1"
  config.vm.network "private_network", ip: "192.168.33.10"
  
  config.vm.define "MyLinux", primary: true do |linuxmachine|
	linuxmachine.vm.hostname = "MyLinux"
	linuxmachine.vm.provision "shell", inline: <<-SHELL
		
		
		
		
	SHELL
end

end
