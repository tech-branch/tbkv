package tbkv

type Item struct {
	Key   string
	Value string
}

type ItemRequest struct {
	Item Item
	Ok   chan bool
}

type Request struct {
	Key    string
	Result chan Result
}

type Result struct {
	Value string
	Ok    bool
}
