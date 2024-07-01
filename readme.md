# wcgo

`wcgo` is a Go-based minimal utility of the Unix `wc` command. It provides functionality to count lines, words,
bytes, and characters in a given file. The project includes a Makefile for building and running the program easily.

## Installation

1. **Clone the repository:**
    ```bash
   
    cd wcgo
    ```

2. **Ensure you have Go and Make installed:**
    ```bash
    go version
    make --version
    ```

## Usage

### Building the Program

Build the program using the provided Makefile:

```bash
make build
```

3 **This will build the program for the target ARCH in the build directory.You can run the program with the desired
options:**
- To count lines, words, bytes, and characters in a file:
```bash
./build/wcgo <filename>
```
- To count lines:
```bash
./build/wcgo -l <filename>
```
- To count only words
```bash
./build/wcgo -w <filename>
```
- To count only bytes:
```bash
./build/wcgo -c <filename>
```


