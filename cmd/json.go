package cmd

// Json formatting for export and import
type Wrap struct {
	Data []Pair `json:"data"`
}
type Pair struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
