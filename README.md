# NLP

Natural Language Processing in Go using [Google's Cloud Natural Language API](https://cloud.google.com/natural-language/docs/reference/libraries)

### Installation

1. Install Go
2. Follow these instructions in [this guide](https://cloud.google.com/natural-language/docs/quickstart-client-libraries) in order to set up a Google Cloud project.
3. Set `export GOOGLE_APPLICATION_CREDENTIALS=[path_to_your_service_account_key_json]`
4. `go build`

### Usage

Pass a string (in quotes) after the `-t` flag to the executable:

```
./nlp -t "The quick brown fox jumped over the lazy dog."
```

It will print to stdout by default. Add the `-o` to write to a file:

```
./nlp -t "And a river runs through it." -o quote.txt
```

### License

MIT
