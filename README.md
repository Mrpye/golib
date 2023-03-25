# golib A Repository of useful functions

## Description
This package contains function that I have collected and found useful and use a lot.


---

## Requirements
* go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install golib

---

## Project folders
Below is a description golib project folders and what they contain
|   Folder        | Description  | 
|-----------|---|
| net    | network and internet functions  |
| compression      | Compression functions tar|
| console      | Terminal function|
| convert      | Type and data Conversion  |
| dir      | Directory functions |
| encrypt      | Encryption functions for file and string |
| file      | Function performing actions of files |
| log      | Log functions mainly around formatting output |
| map_utils      | Utility for doing stuff with maps |
| math      | Simple Math function |
| path      | Path function  |
| string      | String manipulations functions |
---

## Installation and Basic usage
This will take you through the steps to install and get golib up and running.
<details>
<summary>1. Install</summary>

Once you have installed golang you can run the following command to get golib
```bash
go get github.com/Mrpye/golib
```
</details>

<details>
<summary>2. Include in your project</summary>

```go
    include github.com/Mrpye/golib
```
</details>

---


## Todo: 
- Just keep collectinf functions
---

## Change Log
### v0.1.0
- First build

### v0.1.1
- Added convert lib and unit tests

### v0.2.0
- Restructured the lib 
- Lots of unit tests
- Improved the convert

### v0.2.1
- Fix the log.ActionLog formatting
- added:
    - ActionLogDT
    - ActionLogDateOK
    - ActionLogDateFail

### v0.2.2
- Fix the log.ActionLog formatting
- added:
    - Added Type checks
    - Base64EncString made to can take a  []byte or string
    - Added ReadFile 
    - Logs updated to handle different colors for border
    - added CheckNotBlank in string
    - GzipBase64 function 

## license
golib is Apache 2.0 licensed.
