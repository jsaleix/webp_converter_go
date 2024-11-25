# WEBP_CONVERTER_GO

Application to convert webp files to jpg.

## Using executable

You can build the executable yourself with Go:
```bash
make build
# Or according to your OS
make build-win
```

## Usage

The application will list every webp files in the current folder and generate the processed jpg files in a newly created `/result` folder.

You can run it directly by double-clicking if you're using an executable. 
But you can also use it from the command terminal in this way:

```bash
go run .
# Or if compiled
./converter.exe
```

## Misc when using the terminal

Expected output with two webp files in the current directory:

```bash
Done
Processed files: 2
Output dir. C:\Users\Impera\Documents\webp_converter_cli\bin\result
```
