package types

type PlainStackFrame struct {
	FrameId         string `parquet:"fieldid=0,logical=String" json:"frameId"`
	Type            string `parquet:"fieldid=1,logical=String" json:"type"`
	Label           string `parquet:"fieldid=2,logical=String" json:"label"`
	From            string `parquet:"fieldid=3,logical=String" json:"from"`
	To              string `parquet:"fieldid=4,logical=String" json:"to"`
	ContractCreated string `parquet:"fieldid=5,logical=String" json:"contractCreated"`
	Value           string `parquet:"fieldid=6,logical=String" json:"value"`
	Input           string `parquet:"fieldid=7,logical=String" json:"input"`
	Error           string `parquet:"fieldid=8,logical=String" json:"error"`
	ChildrenCount   int32  `parquet:"fieldid=9" json:"childrenCount"`
}
