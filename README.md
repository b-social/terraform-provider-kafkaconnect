Terraform Provider for Kafka Connect
====================================

A [Kafka Connect](https://docs.confluent.io/current/connect/index.html)
provider for [Terraform](https://www.terraform.io/) allowing connectors to
be managed by Terraform.

Installation
------------

### Requirements

*terraform-provider-kafkaconnect* is based on 
[Terraform](https://www.terraform.io) and as such, you need 
[Terraform](https://www.terraform.io/downloads.html) >=0.10.0

### Installation from binaries (recommended)

The recommended way to install *terraform-provider-kafkaconnect* is use the 
binary distributions from the 
[Releases](https://github.com/b-social/terraform-provider-kafkaconnect/releases)
page. The packages are available for Linux and macOS.

Download and uncompress the latest release for your OS. This example uses the 
linux binary.

```sh
> wget https://github.com/b-social/terraform-provider-kafkaconnect/releases/download/0.1.0-rc.5/terraform-provider-kafkaconnect_v0.1.0-rc.5_linux_amd64.tar.gz
> tar -xvf terraform-provider-kafkaconnect*.tar.gz
```

Now copy the binary to the Terraform's plugins folder. If this is your first 
plugin, the directory may need to be created first.

```sh
> mkdir -p ~/.terraform.d/plugins/
> mv terraform-provider-kafkaconnect*/terraform-provider-kafkaconnect ~/.terraform.d/plugins/
```

### Installation from sources

If you wish to compile the provider from source code, you'll first need 
[Go](http://www.golang.org) installed on your machine (version >=1.10 is 
*required*). You'll also need to correctly setup a 
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding 
`$GOPATH/bin` to your `$PATH`.

Clone this repository to: 
`$GOPATH/src/github.com/b-social/terraform-provider-kafkaconnect`

```sh
> mkdir -p $GOPATH/src/github.com/b-social
> git clone https://github.com/b-social/terraform-provider-kafkaconnect.git $GOPATH/src/github.com/b-social/terraform-provider-kafkaconnect
```

Enter the provider directory and build the provider

```sh
> cd $GOPATH/src/github.com/b-social/terraform-provider-kafkaconnect
> ./build provider:build
```

Now copy the compiled binary to Terraform's plugins folder. If this is your 
first plugin, the directory may need to be created first.

```sh
> mkdir -p ~/.terraform.d/plugins/
> cd $GOPATH/bin
> cp terraform-provider-kafkaconnect ~/.terraform.d/plugins/
```

Usage
-----

This plugin currently consists of a single resource `kafkaconnect_connector`.
The resource supports arbitrary connectors, whether custom or bundled with
Kafka Connect.

To define a connector, first define the provider:

```hcl-terraform
provider "kafkaconnect" {
  url = "https://kafkaconnect.example.com"
}
```

Then, define a connector:

```hcl-terraform
resource "kafkaconnect_connector" "file_stream_source" {
  name = "local-file-source"
  class = "FileStreamSource"
  maximum_tasks = "1"

  configuration = {
    file = "/tmp/test.txt"
    topic = "connect-test"
  }
}
```

The `name`, `class` and `maximum_tasks` arguments are required. If custom key
or value converters are required, the resource also supports 
`key_converter_class`, `key_converter_configuration`, `value_converter_class`
and `value_converter_configuration` arguments, where `*_class` is a string 
referencing the class to use and `*_configuration` is a map of additional
configuration to provide to the converter. Any additional connector 
configuration can be passed in the `configuration` map.

Development
-----------

After checking out the repo, run `dep ensure` to install dependencies. Then, 
run `./build provider:test:unit` to run the tests. 

Contributing
------------

Bug reports and pull requests are welcome on GitHub at 
https://github.com/b-social/terraform-provider-kafkaconnect. This project is 
intended to be a safe, welcoming space for collaboration, and contributors are 
expected to adhere to the 
[Contributor Covenant](http://contributor-covenant.org) code of conduct.

License
-------

The gem is available as open source under the terms of the 
[MIT License](http://opensource.org/licenses/MIT).
