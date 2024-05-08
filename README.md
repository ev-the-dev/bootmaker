# Bootmaker
WIP CLI/TUI tool to quickly generate module, and associated files, boilerplate

Written in Go

## Installation
```sh
$ npm i -g @ev-the-dev/bootmaker
```
## Usage
```sh
$ bootmaker
```
> TL;DR
>
> - Type a module name and press **"Enter"/"Return"**
> - Naviagte checklist with **arrow keys** or **j/k**
> - Toggle checklist items with **"Enter"**
> - Submit response by pressing **"q"** or **"ctrl+c"**

Once you run the binary, bootmaker, you'll be prompted to input the name
of the module you wish to create. Currently the only format(s) supported
for names are lowercase, hyphen delimited, words.
**Example: "invoices" or "invoice-items"**

After pressing **"Enter"**, or **"Return"**, you will then see a checklist of
module related files that can be generated. By default they are all selected,
however, you can navigate using **arrow keys** or **vim keybinds (jk)**.
To toggle a selection press **"Enter"**.

When you're ready to submit your selection press **"q"** or **"ctrl+c"** to
exit the TUI and have the process continue to generation.

## Possible Future Improvements
- Ability to create boilerplate files inside of an already existing module
- Conditionally generate DTOs based off of user's checked items
- Conditionally include appropriate imports, controllers, services, etc.
  according to user's checked items
- Provide config functionality for user to determine project structure they'd
  like to use as a blueprint for boilerplate generation
- ???
- Suggestions?
