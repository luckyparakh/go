# URL Shorten
It shortens URL on the basis of base62 encoding.

## Features
The main features are:
1. It can saves URL data in file and in memory (on Go Slice).
2. It can list all present URLs.
3. It does not generate new short link for a URL, if it is already present in file or memory.
4. It also provide URL decode logic, that is not exposed on REST but can be used in future.

## Services
It exposes follwoing services on REST API:
- ListAllShortLink - List all short links present in memory
- CreateShortLink - Create and returns short link of a given URL. Also save URL info in memory.
- ListAllShortLinkFile - List all short links present in file.
- CreateShortLinkFile- Create and returns short link of a given URL. Also save URL info in file.

## Tech Stack
- Go 1.18
- MUX

## Installation
Download docker image from Docker Hub

`docker run <image name>` 

Note: It uses 4000 port.