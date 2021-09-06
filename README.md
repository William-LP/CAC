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
* Visual Studio
* WinSCP
* MySQL Workbench
* Light Shot

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
choco install adobereader
choco install winscp
choco install mysql.workbench
choco install lightshot.install
```
One line command :
``` 
choco install launchy git vscode putty googlechrome python adobereader winscp mysql.workbench lightshot.install -y"
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
* WinSCP
* MySQL Workbench

### Websites to bind (using chrome)
* [Amazon](https://www.amazon.fr/)
* [Facebook](https://www.facebook.com/)
* [Gdrive](https://www.google.com/drive/)
* [Github](https://github.com/)
* [Gmail](https://mail.google.com/)
* [Gmaps](https://www.google.com/maps)
* [Keep](http://keep.google.com)
* [Le Bon Coin](https://leboncoin.fr)
* [Messenger](https://messenger.com)
* [Stackoverflow](https://stackoverflow.com)
* [YouTube](https://youtube.com)

### To do

* Quit the go-thing; not really usefull honestly
* Script the websites' linking within Launchy
* Put the hand into the launchy.ini - might have good stuff to do in there
* Download photo & set them as dynamic background

