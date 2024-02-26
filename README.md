# Website Monitor

This is a simple command-line application that performs website monitoring using HTTP requests. The application reads a list of websites from a file and tests their availability by sending HTTP requests. The results of the monitoring are logged to a file.

## Dependencies

The following dependencies were used in this application:

- Go: A programming language that provides a set of features for building efficient and scalable applications.
- bufio: A Go package that provides buffered I/O functionality.
- http: A Go package that provides HTTP client and server implementations.
- io: A Go package that provides basic I/O primitives.
- ioutil: A Go package that provides I/O utility functions.
- os: A Go package that provides a platform-independent interface to operating system functionality.
- strconv: A Go package that provides string conversion functions.
- strings: A Go package that provides string manipulation functions.
- time: A Go package that provides functionality for measuring and displaying time.

## Getting Started

To run the application, follow these steps:

1. Clone the repository to your local machine.
2. Open a terminal window and navigate to the project directory.
3. Run the command `go run main.go` to start the application.
4. Follow the prompts in the application to choose a command from the menu.

## Usage

The application provides the following commands:

- 1: Start monitoring websites by sending HTTP requests.
- 2: Display the monitoring logs.
- 0: Exit the application.

When the user chooses the "Start monitoring" command, the application reads the list of websites from a file named "sites.txt". It then sends HTTP requests to each website and records the results to a file named "log.txt". The application performs the monitoring for a specified number of times and with a specified delay between each round of monitoring.

When the user chooses the "Display logs" command, the application reads the contents of the file "log.txt" and displays them on the screen.

## Conclusion

This application demonstrates how to use Go to perform website monitoring using HTTP requests. By following the steps outlined above, you should be able to get the application up and running on your local machine in no time.
