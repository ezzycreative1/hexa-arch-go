# hexa-arch-go

### Pre Commit
To install pre-commit package:
* __Mac OS__
  ```console
  brew install pre-commit
  ```
* __Windows & Linux__ </br>
  for windows make sure you have python installed, linux is already available at the beginning of the install.
  ```console
  python -m pip install pre-commit
  ```
* __Check Installation__ </br>
  to make sure the installation is successful
  ```console
  pre-commit --version
  ```
Install pre-commit on local repository
```console
pre-commit install
```
Static Check Package (Go) </br>
  * installation
    ```console
    go install honnef.co/go/tools/cmd/staticcheck@latest
    ```
  * add in linux PATH ```~/.profile``` and for windows in Advanced System Settings
    ```sh
    # file: ~/.profile
    PATH="$HOME/go/bin:$PATH"
    ```
Run pre-commit to the whole rule
```console
pre-commit run --all-files
```