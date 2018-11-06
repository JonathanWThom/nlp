// https://cloud.google.com/natural-language/docs/morphology#parsing_a_syntactic_analysis_response
package main

import (
	"context"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	syntax, err := analyzeSyntax(ctx, client, "The quick brown fox jumped over the lazy dog.")
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Println(syntax.Sentences[0].Text.Content)
	for i := 0; i < len(syntax.Tokens); i++ {
		token := syntax.Tokens[i]
		fmt.Printf("Content: %s\n", token.Text.Content)
		fmt.Printf("Part of Speech (Tag): %s\n", token.PartOfSpeech.Tag)
		fmt.Printf("Part of Speech (Number): %s\n", token.PartOfSpeech.Number)
		fmt.Printf("Part of Speech (Person): %s\n", token.PartOfSpeech.Person)
		fmt.Printf("Part of Speech (Gender): %s\n", token.PartOfSpeech.Gender)
		fmt.Printf("Part of Speech (Case): %s\n", token.PartOfSpeech.Case)
		fmt.Printf("Part of Speech (Tense): %s\n", token.PartOfSpeech.Tense)
		fmt.Printf("Part of Speech (Aspect): %s\n", token.PartOfSpeech.Aspect)
		fmt.Printf("Part of Speech (Mood): %s\n", token.PartOfSpeech.Mood)
		fmt.Printf("Part of Speech (Voice): %s\n", token.PartOfSpeech.Voice)
		fmt.Printf("Part of Speech (Reciprocity): %s\n", token.PartOfSpeech.Reciprocity)
		fmt.Printf("Part of Speech (Proper): %s\n", token.PartOfSpeech.Proper)
		fmt.Printf("Part of Speech (Form): %s\n", token.PartOfSpeech.Form)
		fmt.Printf("Dependency Edge (Head Token Index): %d\n", token.DependencyEdge.HeadTokenIndex)
		fmt.Printf("Dependency Edge (Label): %s\n", token.DependencyEdge.Label)
		fmt.Printf("Lemma: %s\n\n\n", token.Lemma)
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
