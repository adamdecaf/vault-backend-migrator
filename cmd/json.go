package cmd

// todo: Store the vault address to help warn on export/import against the same vault

// Json formatting for export and import
type Wrap struct {
	Data []Pair `json:"data"`
}
type Pair struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
