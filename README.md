# accessDbLoader
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=S-Maciejewski_accessDbLoader&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=S-Maciejewski_accessDbLoader)
[![Access DB Loader - Build artifacts](https://github.com/S-Maciejewski/accessDbLoader/workflows/Access%20DB%20Loader%20-%20Build%20artifacts/badge.svg?branch=main)](https://github.com/S-Maciejewski/accessDbLoader/actions)

A fast way to convert SQL DDL and inserts into Access database file. 
This tool was created to solve the problem of loading multiple SQL inserts (or executing multiple SQL DDL statements), what is not supported in MS Access. 
Right now it supports creating a new MS Access .accdb file from an SQL script.

## Installation
Go to releases and download an executable for your system:
- `adbLoader.exe` for Windows 
- `adbLoader` for linux (amd64)

## Usage
The adbLoader is a command line application and it takes following arguments:
- `-db-path <path>`  a relative path to output (newly generated) Access database file (by default it's './result.accdb')
- `-sql-path <path>` a relative path to the file containing SQL DDL and insert statements
- `-h` displays usage help

