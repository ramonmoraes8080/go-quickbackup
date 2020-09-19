# quickbackup

CLI that backup a list of files and folders into a .zip file and upload it
somewhere using YAML based configuration

## Install

## Usage

Base YAML config for the examples saved at `~/.quickbackup.yaml`:

```
defaults:
  location: my_backups
  schema: my_configs

backends:
  filesystem:
    path: ~/simple-backups
  google_drive:
    credential_json: ~/.config/backup/google-drive.json

locations:
  my_backups:
    backend: filesystem
    path: ~/backups
  pendrive:
    backend: filesystem
    path: /run/media/johnnymnemonic/my-pendrive/
  googledrive:
    backend: google_drive
    path: GoBackup

schemas:
  my_configs:
    - ~/.bashrc
    - ~/.vimrc
  my_etc:
    - /etc/X11/xorg.conf.d/
    - /etc/fstab
    - /etc/bashrc
```

- Backing up config files (`my_configs` schema) to `~/backups` folder

`quickbackup upload -s my_configs -l my_backups`

- Backing up config files (`my_configs` schema) to the pendrive at `/run/media/johnnymnemonic/my-pendrive/`

`quickbackup upload -s my_configs -l pendrive`

- Backing up etc files (`my_etc` schema) to the pendrive at `/run/media/johnnymnemonic/my-pendrive/`

`quickbackup upload -s my_etc -l pendrive`

- Since we defined the `defaults` section we can perform the back of the
`my_config` to the `~/backups` folder using only:

`quickbackup upload`

- List all the backups for `fedora` schema uploaded with `my_backups`

`quickbackup list -s fedora -l my_backups`

- Downloading/Recovering backups

`quickbackup download -s fedora -l my_backups`


## Backends

| Name             | Upload      | Download    | List        |
|------------------|-------------|-------------|-------------|
| Local Filesystem | Implemented | Implemented | Implemented |
| Google Drive     | Not Yet     | Not Yet     | Not Yet     |
| AWS S3           | Not Yet     | Not Yet     | Not Yet     |

## License

Apache License 2.0

Copyright (c) 2020 Ramon Moraes
