package dev

/**
* AppState
 */
type AppState struct {
	Nodes       []Node  `json:"nodes"`
	NodeIndex   int     `json:"nodeIndex"`
	Value       float64 `json:"value"`
	BufferIndex int     `json:"bufferIndex"`
	Clearing    bool    `json:"clearing"`
	HideName    bool    `json:"hideName"`

	Config      Config
	Helpers     Helpers
	NameFactory NameFactory
}

func NewAppState() *AppState {
	state := &AppState{}
	state.Init()
	return state
}

func (state *AppState) Init() {
	state.Config = *NewConfig()
	state.NameFactory = *NewNameFactory(&state.Config)
	state.Helpers = Helpers{Config: &state.Config}

	state.HideName = true
	state.Value = 0.0
	state.NodeIndex = 0
	state.BufferIndex = 0
	state.Nodes = []Node{}
	node := Node{}
	node.Init(state.NameFactory.GetName())
	state.Nodes = []Node{node}
}

// Setter
func (state *AppState) SetCurrentBuffer(buffer string) {
	state.Nodes[state.NodeIndex].Buffer = buffer
}

func (state *AppState) AddBuffer(value string) {
	index := state.GetCurrentNode().Add(value)
	state.BufferIndex = index
}

func (state *AppState) ToggleHideName() {
	state.HideName = !state.HideName
}

// Getter
func (state *AppState) IsEmpty() bool {
	return len(state.Nodes) == 1 && state.Nodes[0].IsEmpty()
}

func (state *AppState) GetCurrentNode() *Node {
	if len(state.Nodes) == 0 {
		return nil
	}
	return &state.Nodes[state.NodeIndex]
}

// Actions
func (state *AppState) AddNode() {
	node := Node{Buffer: "", Result: 0.0}
	node.Init(state.NameFactory.GetName())
	state.Nodes = append(state.Nodes, node)
	state.NodeIndex = len(state.Nodes) - 1
	state.BufferIndex = 0
}

func (state *AppState) ToClear() {
	state.Clearing = true
}

func (state *AppState) CancelClear() {
	state.Clearing = false
}

func (state *AppState) Reset() {
	state.Init()
	state.Clearing = false
}

func (state *AppState) NodeToRecord(node *Node) []string {
	return []string{
		node.Name,
		string(node.Operator),
		node.Buffer,
		state.Helpers.FormatFloat(node.Result),
	}
}
