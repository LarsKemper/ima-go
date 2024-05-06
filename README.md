<h3>ima-go</h3>
<p>Simple ascii image generator to learn go</p>

---

## Overview

- [Examples](https://github.com/LarsKemper/ima-go#examples)
- [About this repository](https://github.com/LarsKemper/ima-go#about-this-repository)
- [Installation](https://github.com/LarsKemper/ima-go#installation)
- [Run ima-go](https://github.com/LarsKemper/ima-go#run-project)

## Examples

### "F**k you, Nvidia"

|                     Original                     |                          Ascii                           |
|:------------------------------------------------:|:--------------------------------------------------------:|
| ![original](/assets/examples/linus.png?raw=true) |   ![ascii](/assets/results/result-linus.png?raw=true)    |

### Colorful image

|                      Original                      |                         Ascii                         |
|:--------------------------------------------------:|:-----------------------------------------------------:|
| ![original](/assets/examples/windows.png?raw=true) | ![ascii](/assets/results/result-windows.png?raw=true) |

## About this repository

### Top-level layout

This repository's contents is divided across following primary sections:

- `/assets` contains assets used for general purposes (you can find example images here)
- `/cmd` contains the main executable
- `/internal` contains the internal packages

## Installation

Clone the repository

```sh
$ git clone https://github.com/LarsKemper/ima-go.git
```

Switch to "ima-go" folder

```sh
$ cd ima-go
```

## Run ima-go

Run using [go](https://go.dev/) cli

```sh
$ go run ./cmd/ima-go/main.go
```