Vagrant and Proxy?
===

Install vagrant-proxyconf plugin

        vagrant plugin install vagrant-proxyconf

Edit Vagrantfile

        vim Vagrantfile

Config

        if Vagrant.has_plugin?("vagrant-proxyconf")
            config.proxy.http     = "http://192.168.0.2:3128/"
            config.proxy.https    = "http://192.168.0.2:3128/"
            config.proxy.no_proxy = "localhost,127.0.0.1,.example.com"
        end

