Inventory
===

- The format for inventory is an INI format

        [webservers]
        foo.example.com
        bar.example.com

        [dbservers]
        one.example.com
        two.example.com
        three.example.com

- Default Ansible’s inventory file

        /etc/ansible/hosts

- Vagrant will generate an inventory file

        .vagrant/provisioners/ansible/inventory/vagrant_ansible_inventory
