package cmd

// Json formatting for export and import
type Wrap struct {
	Data []Item `json:"data"`
}
type Item struct {
	Path  string `json:"path"`
	Pairs []Pair `json:"pairs"`
}
type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
