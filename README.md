# grub-dfct
grub-dfct, a Grub Default File Configuration Tool

## What this does?
This tool modifes /etc/default/grub file changing grub configs, it also allows to reconfigure grub.cfg so the changes would be applied
## Usage
Just run grub-dfct (or whatever the binary is called) and it should write the help dialog
## Building
The apps is written in golang and is using cobra libary for cli managment, if you have already installed golang from your package manager:
```
git clone https://github.com/Hydriam/grub-dfct
cd grub-dfct
go build -o grub-dfct
```
