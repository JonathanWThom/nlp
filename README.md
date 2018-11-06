# NLP

Natural Language Processing in Go using [Google's Cloud Natural Language API](https://cloud.google.com/natural-language/docs/reference/libraries)

Currently supports [analyzing syntax features](https://cloud.google.com/natural-language/docs/analyzing-syntax).
See [https://github.com/JonathanWThom/nlp/blob/master/sample_output.txt](https://github.com/JonathanWThom/nlp/blob/master/sample_output.txt) for sample output.

### Installation

1. Install Go
2. `go get -u github.com/jonathanwthom/nlp`
3. Follow these instructions in [this guide](https://cloud.google.com/natural-language/docs/quickstart-client-libraries) in order to set up a Google Cloud project.
4. Set `export GOOGLE_APPLICATION_CREDENTIALS=[path_to_your_service_account_key_json]`
5. `go install`

### Usage

Pass a string (in quotes) after the `-t` flag to the executable:

```
nlp -t "The quick brown fox jumped over the lazy dog."
```

It will print to stdout by default. Add the `-o` to write to a file:

```
nlp -t "And a river runs through it." -o quote.txt
```

### License

MIT
