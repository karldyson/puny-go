# puny

## Introduction

This is a simple program that converts between ASCII and unicode.

It's mainly written for converting punycode, and so adding CLI options to force to_ascii or to_unicode
are currently missing.

It assumes strings starting xn-- are ASCII for conversion to unicode, and everything else is unicode
for conversion to ASCII.

## Usage

* --terse
  outputs the conversion bare

* --version
  output the code version

* --revision
  output more verbose code revision information

## ToDo

* add CLI options to force to_ascii or to_unicode

## Bug Reporting

Please use the issues feature, or, of course, offer to contribute.

## Licence and Copyright

This code is Copyright (c) 2024 Karl Dyson.

All rights reserved.

## Warranty

There's no warranty that this code is safe, secure, or fit for any purpose.

I'm not responsible if you don't read the code, check its suitability for your
use case, and wake up to find it's eaten your cat...
