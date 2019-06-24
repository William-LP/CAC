# Chocolatey's installation
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

$packageList = Get-Content "$($PSScriptRoot)\package.txt"

# Packages' installation using chocolatey
foreach($package in $packageList) {
    choco install $package -y
}

# Ungit's installation using npm (choclatey repo isn't up to date sadly)
npm install -g ungit

# Replace launchy.init with home made config
Copy-Item -Path "launchy.ini" -Destination "$env:AppData/Launchy/launchy.ini" -force
if ($?) {
    Write-Host "Launchy's .ini config updated"
} else {
    Write-Host "Error while updating Launchy's .ini config"
    Write-Host "Maybe worth checking privileges' settings or doing it by hand (refer to https://github.com/William-LP/CAC/blob/master/launchy.ini)"
}

# Rescanning the disk using the new .ini file
& "${env:ProgramFiles(x86)}\Launchy\Launchy.exe" /rescan