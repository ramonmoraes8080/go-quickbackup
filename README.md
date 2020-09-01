# Backup

CLI that backup a list of files and folders into a .zip file and upload it
somewhere using YAML based configuration

## Install

## Usage

Base YAML config for the examples saved at `~/.backup.yaml`:

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

- Backing up config files to `~/backups`

`backup` or `backup -s my_configs -to my_backups`

- Backing up config files to the pendrive at `/run/media/johnnymnemonic/my-pendrive/`

`backup -to pendrive` or `backup -s my_configs -to pendrive`

- Backing up etc files to the pendrive at `/run/media/johnnymnemonic/my-pendrive/`

`backup -s my_etc -to pendrive`

- Backing up config files to the pendrive at `/run/media/johnnymnemonic/my-pendrive/`

`backup -to pendrive` or `backup -s my_configs -to pendrive`

## Backends

| Name             | Upload | Download  | List |
|------------------|--------|-----------|------|
| Local Filesystem |    X   |           |      |
| Google Drive     |        |           |      |
| AWS S3           |        |           |      |

## License

Apache License 2.0

Copyright (c) 2020 Ramon Moraes
