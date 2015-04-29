Port Forward
===

**Enabling application access in real world**

        vim Vagrantfile

        config.vm.network "forwarded_port", guest: 8080, host: 8080

