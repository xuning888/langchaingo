package vectorstores

// Option is a function that configures an Options.
type Option func(*Options)

// Options is a set of options for similarity search and add documents.
type Options struct {
	NameSpace      string
	ScoreThreshold float64
	Filters        map[string]any
}

// WithNameSpace returns an Option for setting the name space.
func WithNameSpace(nameSpace string) Option {
	return func(o *Options) {
		o.NameSpace = nameSpace
	}
}

func WithScoreThreshold(scoreThreshold float64) Option {
	return func(o *Options) {
		o.ScoreThreshold = scoreThreshold
	}
}

// Vector searches can be limited bsaed on metadata filters. Searches with  metadata
// filters retrieve exactly the number of nearest-neighbors results that match the filters. In
// most cases the search latency will be lower than unfiltered searches
// See https://docs.pinecone.io/docs/metadata-filtering
func WithFilters(filters map[string]any) Option {
	return func(o *Options) {
		o.Filters = filters
	}
}
