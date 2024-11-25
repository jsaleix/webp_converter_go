# WEBP_CONVERTER_GO

Application to convert webp files to jpg.

> ⚠︎ **The following commands are to be executed in the context of the ```code/``` subfolder.**

## Using executable

You can build the executable yourself with Go:
```bash
make build
# Or according to your OS
make build-win
```

## Usage

The application will list every webp files in the current folder and generate the processed jpg files in a newly created `/result` folder.

You can run it directly by double-clicking or use it from the command terminal in this way:

```bash
go run .
# Or on windows
./converter.exe
# Or on MacOS after running chmod +x converter
./converter
```

## Misc when using the terminal

Expected output with two webp files in the current directory:

```bash
Done
Processed files: 2
Output dir. C:\Users\Impera\Documents\webp_converter\bin\result
```
