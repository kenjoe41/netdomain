# Netdomain
Netdomain is a command-line tool that allows you to search for subdomains of a specific domain using the [Netlas](https://netlas.io/) API.

## Installation
To install Netdomain, make sure you have [Go](https://golang.org/) installed on your system. Then, run the following command:
```bash
go install github.com/kenjoe41/netdomain@latest
```

## Usage
To use Netdomain, you will need an API key from [Netlas](https://netlas.io/). You can sign up for an API key on their website.

Once you have your API key, you can run Netdomain using the following command:
```bash
netdomain -d <domain> -apikey <API_KEY>
```
Where `<domain>` is the domain you want to search for subdomains of, and `<API_KEY>` is your Netlas API key.

### Optional Flags
- `-of <OUTPUT_FILE>`: Specify an output file for the results. If this flag is not provided, the results will be printed to stdout.
- `-silent`: Run in silent mode, which suppresses stdout output.

## Examples
Search for subdomains of `example.com` and print the results to stdout:
```bash
netdomain -d example.com -apikey YOUR_API_KEY
```
Search for subdomains of `example.com` and write the results to a file called `output.txt`:
```bash
netdomain -d example.com -apikey YOUR_API_KEY -of output.txt
```
Search for subdomains of `example.com` and write to file in silent mode (no stdout):
```bash
netdomain -d example.com -apikey YOUR_API_KEY -of output.txt -silent
```
## Contributing

1. Fork the repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

Netdomain is released under the MIT License. See [LICENSE](https://github.com/kenjoe41/netdomain/blob/main/LICENSE) for more details.
