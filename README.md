# Svelte Unused Files Analyzer
This Go program is designed to analyze `.svelte` files in a directory, identifying which files are used and which are not. It provides an option to search recursively and can output the list of unused files to a text file.

## Flags

- `-d`: Specifies the directory to search for `.svelte` files. Defaults to the `/src` directory.
- `-i`: Specifies the input file containing a list of files to ignore. Defaults to `ignore_files.txt`.
- `-o`: Specifies the output file for the list of unused files. Defaults to `unused_files.txt`.
- `-v`: Enables verbose output. (This will also disable the progress display.)
- `-np`: Disables the progress display (useful when redirecting output).
- `-r`: Enables recursive search (to also find files only used in unused files).

## Usage
  
  Download the latest release from the [releases page](https://github.com/FerrarioChristian/svelte-unused-components/releases) and place it in the package directory, the program will analyze the `/src` folder.\
  Then run the executable from the command line with the flags you need.\
  The program will output the list of unused files to the specified output file (defaults to `unused_files.txt`).

  ```sh
  ./svelte-unused-components -r
  ```

## Build and Run

1. **Clone the repository**:
   ```sh
   git clone https://github.com/FerrarioChristian/svelte-unused-components.git
   cd svelte-unused-components
   ```
2. **Build the program**:
   ```sh
    go build .
    ```
3. **Run the program**:
    ```sh
    ./svelte-unused-components -r
    ```

## License
Distributed under the MIT License. See `LICENSE.txt` for more information.
