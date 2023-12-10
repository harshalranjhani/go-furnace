# go-furnace

## Overview

A simple boilerplate generator which generates multiple templates in a concurrent manner to save time!

## Prerequisites

- Node.js and npm installed

## Configuration: Electron templates

To generate a electron template , you need to configure sudoers to allow npm to be executed without a password prompt. Follow these steps:

1. Open the sudoers file using the `visudo` command:

   ```bash
   sudo visudo

2. Add the following line to allow running npm install electron without a password prompt:
   
   ```bash
   your_username ALL=(ALL) NOPASSWD: /usr/bin/npm install electron

Replace your_username with your actual macOS username.

3. Save the changes and exit the editor
