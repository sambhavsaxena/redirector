# redirector
🎲 A custom load balancer program written in Golang that works on Round Robin algorithm to serve requests to prevent network congestion and redirection.

<div align="center">

https://github.com/sambhavsaxena/redirector/assets/76242518/785873a0-9e4d-4e47-9f94-0dc247d23d6e

</div>


- **Usage:** A minimal program to balance and redirect server requests.
- **Prevents single point of failure:** Ensures that SPoFs are eliminated in case of HTTP requests.
- **Custom domains:** The redirect URL can be fully customized.
- **Low complexity:** The application is relatively simple to understand and work with.
- **No limit on server URLs:** The number of proxy servers can be added single-handedly.
- **Improves overall reliability:** Improves the reliability of the network by efficient distribution of requests.

## Prerequisites

Golang core compiler has to be installed on the machine. Along with that, Golang CLI toolkit is necessary for compiling the source for the dependent architecture.

## Installation

Follow the steps in order to create an executable of your architecture:

- Fork and clone [this](https://github.com/sambhavsaxena/redirector) repository to make an instant copy of the content.
- Alternatively, you can download the source and set it up with Github Desktop.
- Open the root folder in the code editor of your preference, and run the following commands:

```
 go build
 go run main.go
```

This [file](https://github.com/sambhavsaxena/redirector/blob/main/golang-loadbalancer) is a pre-installed binary for my x86-64 Arch based distribution, and similar architectures.


## Documentation

Check out the [official docs](https://go.dev/doc/) for a quick overview of the structure used. 

You can improve it by sending pull requests to on a seperate branch.

## Contributing

The main purpose of this repository is to continue evolving Golang core, making it faster and easier to use. Development of Golang happens in the open on GitHub, and we are grateful to the community for contributing bug fixes and improvements. Read below to learn how you can take part in improving Golang based apps.

## Code of Conduct

Redirector has adopted a Code of Conduct that we expect every participants to adhere to. Please read [the full text](https://go.dev/conduct) so that you can understand what actions will and will not be tolerated.

## Community

Interact and find more about the Go community [here](https://forum.golangbridge.org/).
