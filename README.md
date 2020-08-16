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

MIT License

Copyright (c) 2020 Ramon Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
