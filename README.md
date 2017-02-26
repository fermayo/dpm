# dpm

Install development tools locally to your project using docker containers

[![asciicast](https://asciinema.org/a/f056w0w93x0b2tpgppkzegw50.png)](https://asciinema.org/a/f056w0w93x0b2tpgppkzegw50)


## Installation

Install Docker. Then:

    curl -L "https://github.com/fermayo/dpm/releases/download/0.2.0/dpm-$(uname -s)-$(uname -m)" -o /usr/local/bin/dpm; chmod +x /usr/local/bin/dpm

And add `~/.dpm` to your shell path. For example, for `bash`:

    echo "export PATH=$PATH:$HOME/.dpm" >> ~/.bashrc

If you want project commands defined with `dpm` to override system commands (if you trust your projects!), use the following instead:

    echo "export PATH=$HOME/.dpm:$PATH" >> ~/.bashrc


## Usage

### Defining commands for your project

Add a file called `dpm.yml` to your project root defining the commands you want to use for development. For example:

```yaml
commands:
  go:
    image: golang:1.7.5
    context: /go/src/github.com/fermayo/dpm

  glide:
    image: dockerepo/glide
```

These are container definitions with a syntax similar to Compose. The following defaults will apply:
* The entrypoint of the container defaults to the command name
* The project root is mounted read/write to `context` (default: `/run/context`)
* The working directory defaults to `context` (default: `/run/context`)
* Containers are deleted after execution (`--rm`)
* They are run in interactive mode (`-i`) and with tty enabled (`-t`)

You can override these or define any other container attributes.

Currently only the following attributes are supported: `image`, `entrypoint`, `context`.

### Installing commands

Execute:

    dpm install
    
and it will create all command aliases in `.dpm/`. Run it every time you update `dpm.yml`.


### Using project commands

From the project root, run the following to enable its installed commands:

    dpm activate

Then, just execute them as if they were installed in your OS:

    $ go version
    go version go1.7.5 linux/amd64

You can also list which commands are available by running:

    dpm list


### Uninstalling commands

To remove all commands from the current project, just run:

    dpm uninstall
