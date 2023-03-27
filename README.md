# golib A Repository of useful functions

## Description
golib is a collation of helper functions I have collected, written and regularly use in my projects, I have a rule if i find myself copying 
and pasting from previous projects then it gets tidied up, documented and added to the library. that way i can easily find it for future use.
Feel free to use be aware that things may get moved around and renamed. At some point when i am happy with it i'll make it a full release 
and take care not to break compatibility but for now Use at you own risk.

If you would like to contribute i am more than happy for the help.



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
| str      | String manipulations functions |
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
- Add some examples for now take a look at the tests
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
    - Moved string to str stop clashing with standard strings library
    - Documented the functions
    - Updated the documents
  
## license
golib is Apache 2.0 licensed.
