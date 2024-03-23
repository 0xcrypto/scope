# 🔭 Scope

Scope is a small tool that parses target files provided by different bug bounty platforms and creates separate .txt files based on their type.

## Currently Supported Formats

| Format | Platform | Support | Remarks |
|--------|----------|---------|---------|
| csv | HackerOne | ✅ | Fully supported |
| json | Burpsuite | 🅿️ | Planned but might take some time |
| html | HackerOne | ⏺️ | In development using headless browser |
| html | Intigriti | 🅿️ | Planned to be done after HackerOne |
| html | Bugcrowd | 🅿️ | Planned to be done after HackerOne |
| html | YesWeHack | 🅿️ | Planned to be done after HackerOne |

## Installation

```sh
go install -v github.com/0xcrypto/scope@latest
```

## Usage

```sh
scope inputs.csv

# With stdin
cat inputs.csv | scope /dev/stdin
```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).