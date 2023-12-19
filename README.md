# Worklog Application

## Overview

The Worklog application is a command-line tool developed in Go, which provides the functionality to list files in the current directory and write the details of files to a remote API. The tool leverages the Cobra library for creating a structured command-line interface.

## How this entire repo was generated

![How the entire repo was created](worklog.build.webm)

## Requirements

- Go programming language (1.x+)
- A functional Go environment
- External Go packages:
  - Cobra: `github.com/spf13/cobra`

## Installation

First, clone the repository to your local machine:

```sh
git clone https://github.com/your_username/worklog.git
```

Next, navigate to the root directory of the project:

```sh
cd worklog
```

To download the necessary Go dependencies, run:

```sh
go mod tidy
```

## Building the Application

To build the application, you can use the Go build command:

```sh
go build -o worklog
```

This command will compile the program and produce an executable named `worklog`.

## Usage

The application has two primary commands:

- `list`: Lists the files in the current directory with their sizes in a human-readable format.
- `write`: Sends the details of files in the current directory as a JSON payload to a predefined API endpoint.

### Listing Files  

To list the files, run the following command:

```sh
./worklog list
```

This will output the names and sizes of files located in the directory where the command was executed.

### Writing File Details to an API

Before using the `write` command, ensure to set the correct API endpoint by modifying the `apiEndpoint` variable in the `WriteFiles` function inside the source code.

Once the API endpoint is set, run:

```sh
./worklog write
```

The above command will serialize the file details into a JSON payload and make an HTTP POST request to the specified API endpoint.

## Contributing

Contributions are welcome. Please feel free to create issues or pull requests on the project's GitHub repository.

## License

Please add your preferred license here.
