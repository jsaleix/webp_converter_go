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

The application will create two folders `/input` and `/output`.
Place all your webp files in `/input` and run the app. It will generate the processed jpg files in `/output` .

You can run it directly by double-clicking if you're using an executable. 
But you can also use it from the command terminal in this way:

```bash
go run .
# Or if compiled
./converter.exe
```

## Misc when using the terminal

Expected output with two webp files in `/input` :

```bash
Done
Processed files: 2
Output dir. C:\Users\Impera\Documents\webp_converter_cli\bin\output
```

You can obtain the paths for `/input` and `/output` by using the **--help** option:

```bash
go run . --help
# Or if compiled
./converter.exe --help
```

```bash
Put your webp files in: C:\Users\Impera\Documents\webp_converter_cli\bin\input
Jpg files will be generated in: C:\Users\Impera\Documents\webp_converter_cli\bin\output
```

