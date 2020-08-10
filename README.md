# terminstall

One command-line tool to install all of your apps on new machines based on a `yml` config file! No more waiting waiting hours or forgetting which apps to install, just create a install file and you are good to go!

## Usage: Install file

Terminstall reads all your desired apps from `yml/yaml` files with the following basic format:

```yml
version: "1.0"
install_providers: true
providers:
  - name: brew # name used to install the apps using this provider
    base_command: brew install # base command to run for installing the apps when this provider is used
    install_provider_commands: # install and configure the provider
      - /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
      - brew update

  - name: brewcask
    base_command: brew cask install

machine_setup:
  continue_on_error: true # continue or stop upon an app installation error
  installations:
    - provider: brewcask # alias for the base_command: brew cask install
      apps:
        - slack
        - alfred
        - vscode
        - goland

    - provider: brew # alias for the base_command: brew install
      apps:
        - aws
        - gcloud
        - docker
        - kubectl
        - k9s
```

### providers

Providers are commands that will provide the software/apps that you want. A `macOS` example of a provider would be `homebrew`. Hence, the install file allows one to setup many providers which will be used to install their apps!

- `providers`: list of providers which will be configured and used to install the apps
- `providers.name`: name of the provider to be used on the `machine_setup.installations` section to install the apps
- `providers.install_provider_commands`: list of commands used to install/setup the provider
- `providers.base_command`: base command to install the apps! Each app name will be the last argument of the base command

Example: `base_command slack`, if `base_command = brew install`, then `brew install slack`

### machine_setup

Here one can add all the apps that each provider will install:

- `machine_setup.continue_on_errors`: if one of the installs fail, stop the whole process or ignore and continue with other installs
- `machine_setup.installations`: list of apps to install for each configured provider

For the given example yaml: the following command will be run for the `brewcask` provider:

```bash
brew cask install slack
brew cask install alfred
brew cask install vscode
...
```

For the `brew` provider, the following commands will be run:

```bash
brew install aws
brew install gcloud
brew install docker
...
```

## Running the project

Run the following command on the terminal:

```bash
make
```

## Running tests of the project

Run the following command on the terminal:

```bash
make tests
```
