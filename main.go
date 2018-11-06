// https://cloud.google.com/natural-language/docs/morphology#parsing_a_syntactic_analysis_response
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	var text, output string
	flag.StringVar(&text, "t", "", "Text to be analyzed")
	// 	flag.String(output, "o", "File to write to, defaults to stdout")
	flag.Parse()

	if text == "" {
		log.Fatalf("Must enter text with flag -t [your_text_here]")
	}

	w := os.Stdout
	//	var file *os.File

	if output != "" {
		file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("File cannot be created")
		}
		defer file.Close()
		// 	w = bufio.NewWriter(file)
	}

	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	syntax, err := analyzeSyntax(ctx, client, text)
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Fprintf(w, "TEXT:\n%s\n\n", text)

	index := 0

	for i := 0; i < len(syntax.Sentences); i++ {
		sentence := syntax.Sentences[i]
		text := sentence.Text
		content := text.Content
		sentenceBegin := int(text.BeginOffset)
		sentenceEnd := sentenceBegin + len(content) - 1
		fmt.Fprintf(w, "SENTENCE: %s\n\n", content)

		for index < len(syntax.Tokens) && int(syntax.Tokens[index].Text.BeginOffset) <= sentenceEnd {
			token := syntax.Tokens[index]
			fmt.Fprintf(w, "Content: %s\n", token.Text.Content)
			fmt.Fprintf(w, "Part of Speech (Tag): %s\n", token.PartOfSpeech.Tag)
			fmt.Fprintf(w, "Part of Speech (Number): %s\n", token.PartOfSpeech.Number)
			fmt.Fprintf(w, "Part of Speech (Person): %s\n", token.PartOfSpeech.Person)
			fmt.Fprintf(w, "Part of Speech (Gender): %s\n", token.PartOfSpeech.Gender)
			fmt.Fprintf(w, "Part of Speech (Case): %s\n", token.PartOfSpeech.Case)
			fmt.Fprintf(w, "Part of Speech (Tense): %s\n", token.PartOfSpeech.Tense)
			fmt.Fprintf(w, "Part of Speech (Aspect): %s\n", token.PartOfSpeech.Aspect)
			fmt.Fprintf(w, "Part of Speech (Mood): %s\n", token.PartOfSpeech.Mood)
			fmt.Fprintf(w, "Part of Speech (Voice): %s\n", token.PartOfSpeech.Voice)
			fmt.Fprintf(w, "Part of Speech (Reciprocity): %s\n", token.PartOfSpeech.Reciprocity)
			fmt.Fprintf(w, "Part of Speech (Proper): %s\n", token.PartOfSpeech.Proper)
			fmt.Fprintf(w, "Part of Speech (Form): %s\n", token.PartOfSpeech.Form)
			fmt.Fprintf(w, "Dependency Edge (Head Token Index): %d\n", token.DependencyEdge.HeadTokenIndex)
			fmt.Fprintf(w, "Dependency Edge (Label): %s\n", token.DependencyEdge.Label)
			fmt.Fprintf(w, "Lemma: %s\n\n", token.Lemma)

			index++
		}
	}
}

func analyzeSyntax(ctx context.Context, client *language.Client, text string) (*languagepb.AnnotateTextResponse, error) {
	return client.AnnotateText(ctx, &languagepb.AnnotateTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		Features: &languagepb.AnnotateTextRequest_Features{
			ExtractSyntax: true,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
}
