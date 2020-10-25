# accessDbLoader
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)

A fast way to convert SQL DDL and inserts into Access database file. 
This tool was created to solve the problem of loading multiple SQL inserts (or executing multiple SQL DDL statements), what is not supported in MS Access. 
Right now it supports creating a new MS Access .accdb file from an SQL script.

## Usage
Clone the repo and build the module with `go build accessDbLoader`.

In the future I plan to include .exe version for direct download in releases
