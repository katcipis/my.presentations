Tips and tricks
===

config.vm.box_url

        config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"

config.vm.hostname

        config.vm.hostname = "newapp"

vb.customize

        vb.customize ["modifyvm", :id, "--name", "newapp"]
        vb.customize ["modifyvm", :id, "--memory", 512]
        vb.customize ["modifyvm", :id, "--cpus", 2]
        vb.customize ["modifyvm", :id, "--ioapic", "on"]

