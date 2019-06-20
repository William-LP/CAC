# CAC
Config As Code, to build my config anywhere in a click

### Soft to install 
* Chocolatey
* Chrome
* Git
* Git Bash
* Launchy
* Office Suite
* Putty
* Python
* Ungit
* Visual Studio

## Chocolatey
In an **administrator** cmd :
```
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
```
Chocolatey Package :
``` 
choco install launchy 
choco install git 
choco install vscode 
choco install putty
choco install googlechrome
choco install python
choco install nodejs
npm install -g ungit
```
One line command :
``` 
choco install launchy git vscode putty googlechrome python nodejs -y && "c:\Program Files\nodejs\npm" install -g ungit
```
## Launchy
### Softs to bind
* Chrome
* Git Bash (admin)
* Putty
* Visual Studio
* Shortcuts
* Word
* Excel
* PowerPoint

### Websites to bind (using chrome)
* [Amazon](https://www.amazon.fr/)
* [Facebook](https://www.facebook.com/)
* [Gdrive](https://www.google.com/drive/)
* [Github](https://github.com/)
* [Gmail](https://mail.google.com/)
* [Keep](http://keep.google.com)
* [Le Bon Coin](https://leboncoin.fr)
* [Messenger](https://messenger.com)
* [Stackoverflow](https://stackoverflow.com)

