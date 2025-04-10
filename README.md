# Victor

Open a blank Markdown File in your editor of choice from the command line that uses a template and saves to the desired location noted in the config file.

_Inspired by_ [Hugo](https://gohugo.io/) _to make wanting to write frictionless._

# How to Install

- Make sure you have GO installed on your machine
- Clone this repository: 
    - `git clone git@github.com:Bones1335/blog-journal-tool.git`
- Install the missing GO package dependencies (if you're missing them):
    - `go mod tidy`
- Finally, build the tool:
    - `go build`

# How to Use

- This tool can use any editor called from the command line by writing its command to the `editor` struct value found in the `config.go` file.
- The save location is based on your home directory, so you can have the tool save your Markdown document anywhere within by replacing the `blog` and `journal` struct values, found in the `config.go` file.
- Type `./blog-journal-tool new blog` or `./blog-journal-tool new journal` to launch your editor of choice to start writing an `index.md` file with an auto-generated template ready to go. This is what the template looks like:
```
---
title: 'Testing File'
date: %v
url: 
categories: 
    - example 1
    - example 2
tags:
    - example 3
    - example 4
---
```
- It's a YAML header where you can change the following metadata to suite your needs. You write your document/Markdown after the `---` otherwise the tool will freakout. 
- Notice the `url` value is empty. You don't have to fill that out as it will auto-generate based on your title, creating the directory that your `index.md` file will be saved to.
    - For example, the auto-generated url for this test metadata would be `/testing-file`.
- The `date` is auto-generated for the date and time you open a new document.
- Once you've saved whatever you've written, the tool will auto-generate the directory to save your `index.md` file to using the `url` created from your title and save it to the `blog` or `journal` location indicated in the config and depending on the argument you passed in.
    - For example, this test file will save to `$HOME/Documents/blog/2025/testing-file/index.md`.
