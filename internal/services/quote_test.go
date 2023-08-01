package services

import (
	"reflect"
	"testing"
)

func TestQuoteService_GetRandomQuote(t *testing.T) {
	type fields struct {
		quotes []Quote
	}
	tests := []struct {
		name   string
		fields fields
		want   Quote
	}{
		{
			name: "Get the one quote",
			fields: fields{
				[]Quote{
					{
						"Author 1",
						"Quote 1",
					},
				},
			},
			want: Quote{
				"Author 1",
				"Quote 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &QuoteService{
				quotes: tt.fields.quotes,
			}
			if got := qs.GetRandomQuote(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
