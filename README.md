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

The application will create two folders `/input` and `/output`.
Place all your webp files in `/input` and run the app. It will generate the processed jpg files in `/output` .

You can run it directly by double-clicking or use it from the command terminal in this way:

```bash
./converter.exe
# Or
./converter
```

## Misc when using the terminal

Expected output with two webp files in `/input` :

```bash
Done
Processed files: 2
Output dir. C:\Users\Impera\Documents\webp_converter\bin\output
```

You can obtain the paths for `/input` and `/output` by using the **--help** option:

```bash
./converter.exe --help
# Or
./converter --help
```

```bash
Put your webp files in: C:\Users\Impera\Documents\webp_converter\bin\input
Jpg files will be generated in: C:\Users\Impera\Documents\webp_converter\bin\output
```

## If you want to use it with Go

Use the ```--dev``` flag if you intend to use it without building an executable.

```bash
go run . --dev
go run . --dev --help
```