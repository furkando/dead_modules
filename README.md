# Dead Modules

Dead Modules is a CLI tool written in Go to search for and remove old `node_modules` directories within a working directory. It provides an interactive terminal interface to list, select, and delete these directories, helping you clean up disk space and manage your project dependencies more effectively.

## Features

- **Interactive CLI**: Provides a user-friendly terminal interface using the `tview` package.
- **Directory Scanning**: Recursively searches for `node_modules` directories within the working root directory.
- **Selection and Deletion**: Allows users to select directories and delete them interactively.
- **Dynamic Sorting**: Lists directories sorted by their last modified date.
- **Status Updates**: Displays status messages and updates in real-time during the search and deletion processes.

## Installation

1. ** Add brew tap**

   ```sh
   brew tap furkando/tap
   ```

2. **Install the Application**

   ```sh
   brew install dead_modules
   ```

3. **Run the Application**

   ```sh
   dead_modules
   ```

## Manual Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/furkando/dead_modules.git
   cd dead_modules
   ```

2. **Build the Project**
   Ensure you have Go installed. If not, download and install it from [golang.org](https://golang.org/dl/).

   ```sh
   go build -o dead_modules
   ```

3. **Run the Application**
   ```sh
   ./dead_modules
   ```

## Usage

### Starting the Application

Run the application in your terminal:

```sh
./dead_modules
```

### Navigating the Interface

- **Search for Directories**: The application starts by prompting you to search for `node_modules` directories. Press `Enter` to begin the search.
- **View Results**: Directories are listed in a table, sorted by their last modified date.
- **Select Directories**: Use the arrow keys to navigate the list. Press `Space` to select or deselect a directory.
- **Delete Selected Directories**: Press `Enter` to delete the selected directories. The status will update to show `[DELETING]` and `[DELETED]` messages.
- **Quit the Application**: Press `q` to quit the application at any time.

### Status Indicators

- **[yellow]Searching...**: Indicates that the application is currently searching for `node_modules` directories.
- **[yellow][Deleting]**: Indicates that the application is currently deleting the specified directory.
- **[green][DELETED]**: Indicates that the specified directory has been successfully deleted.
- **[red]Error**: Indicates an error occurred during the deletion process.

## Code Structure

- **main.go**: Entry point of the application.
- **ui/ui.go**: Handles the user interface setup and interactions.
- **ui/table.go**: Manages the table display and updates.
- **search/search.go**: Contains the logic for searching `node_modules` directories.
- **delete/delete.go**: Contains the logic for deleting the selected directories.

## Example

Below is a sample terminal session using Dead Modules CLI Tool:

```sh
$ dead_modules
```

- On startup, you will see the prompt:

  ```
  Dead Modules v1.0.0

  Press Enter to start searching for node_modules...
  ```

- After pressing `Enter`, the search begins and results are displayed in a table.
- Use the arrow keys to navigate, `Space` to select/deselect, and `Enter` to delete.

## Debugging

To enable debugging mode, add the `-debug` flag when running the application:

```sh
./dead_modules -debug
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have suggestions or bug reports.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
