# Chocolatey's installation
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# Packages' installation using chocolatey
foreach($package in Get-Content .\package.txt) {
    choco install $package
}

# Ungit's installation using npm (choclatey repo isn't up to date sadly)
npm install -g ungit