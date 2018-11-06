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

	text := `
		He was an old man who fished alone in a skiff in the Gulf Stream and he had
		gone eighty-four days now without taking a fish. In the first forty days a
		boy had been with him. But after forty days without a fish the boy's parents
		had told him that the old man was now definitely and finally salao, which is
		the worst form of unlucky, and the boy had gone at their orders in another
		boat which caught three good fish the first week. It made the boy sad to see
		the old man come in each day with his skiff empty and he always went down to
		help him carry either the coiled lines or the gaff and harpoon and the sail
		that was furled around the mast. The sail was patched with flour sacks and,
		furled, it looked like the flag of permanent defeat.
	`

	syntax, err := analyzeSyntax(ctx, client, text)
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("TEXT:\n%s\n\n", text)

	index := 0

	for i := 0; i < len(syntax.Sentences); i++ {
		sentence := syntax.Sentences[i]
		text := sentence.Text
		content := text.Content
		sentenceBegin := int(text.BeginOffset)
		sentenceEnd := sentenceBegin + len(content) - 1
		fmt.Printf("SENTENCE: %s\n\n", content)

		for index < len(syntax.Tokens) && int(syntax.Tokens[index].Text.BeginOffset) <= sentenceEnd {
			token := syntax.Tokens[index]
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
			fmt.Printf("Lemma: %s\n\n", token.Lemma)

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
