# WEBP_CONVERTER_CLI

Command-line interface to convert webp files to jpg.

## Usage

The application will create two folders ```/input``` and ```/output```.
Place all your webp files in ```/input``` and launch the app. It will generate the processed jpg files in ```/output``` .

```bash
go run .
# Or if compiled
./converter.exe
```

Expected output with two webp files in ```/input``` :

```bash
Done
Processed files: 2
Output dir. C:\Users\Impera\Documents\webp_converter_cli\bin\output
```

You can obtain the paths for ```/input``` and ```/output``` by using the **--help** option:

```bash
go run . --help
# Or if compiled
./converter.exe --help
```

```bash
Put your webp files in: C:\Users\Impera\Documents\webp_converter_cli\bin\input
Jpg files will be generated in: C:\Users\Impera\Documents\webp_converter_cli\bin\output
```

## Misc

### Using binary
You can build the binary yourself with Go:

    make build-win

### Adding to path
The best way to use this type of CLI is to add it to your path so that you can simply use it in your terminal when you need it:

    $ webp

How to do it:

 - [on Linux](https://askubuntu.com/a/322773)
 - [on MacOS](https://medium.com/codex/adding-executable-program-commands-to-the-path-variable-5e45f1bdf6ce)
 - [on Windows](https://medium.com/@kevinmarkvi/how-to-add-executables-to-your-path-in-windows-5ffa4ce61a53)
