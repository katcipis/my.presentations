Provisioning
===

Vagrant has built-in support for automated provisioning.
Using this feature, Vagrant will automatically
install software when you vagrant up.

        config.vm.provision "ansible" do |ansible|
          ansible.playbook = "provisioning/playbook.yml"
          ansible.sudo = true
        end

