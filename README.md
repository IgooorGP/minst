# terminstall

One command-line tool to install all of your apps on new machines with a single command!

## Setup

First idea of `Terminstall` yml configuration file:

```yml
version: "1.0"
install_providers: "yes"
providers:
  - name: brew
    install_provider_commands:
      - /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
      - brew update
    install_apps_command: brew <app_name>

  - name: brewcask
    install_apps_command: brew cask <app_name>

machine_setup:
  continue_on_error: "yes"
  installations:
    - provider: brewcask
      apps:
        - slack
        - alfred
        - vscode
        - pycharm

    - provider: brew
      apps:
        - aws
        - gcloud
        - docker
        - kubectl
        - k9s
```
