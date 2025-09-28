package lsp

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializationOptions struct {
	Provider string `json:"provider"`
	Params   any    `json:"params"`
}

type InitializeParams struct {
	ClientInfo            ClientInfo            `json:"clientInfo"`
	InitializationOptions InitializationOptions `json:"initializationOptions"`
}

type TextDocumentSyncKind struct {
	OpenClose bool `json:"openClose"`
	Change    int  `json:"change"`
}

type inlineCompletionKind struct {
}

type ServerCapabilities struct {
	TextDocumentSync TextDocumentSyncKind `json:"textDocumentSync"`
	InlineCompletion inlineCompletionKind `json:"inlineCompletionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

func NewInitializeResponse(id int) ResponseMessage {
	return ResponseMessage{
		Message: Message{
			JsonRPC: "2.0",
		},
		ID: id,
		Result: InitializeResult{
			Capabilities: ServerCapabilities{TextDocumentSync: TextDocumentSyncKind{
				OpenClose: true,
				Change:    2,
			}},
			ServerInfo: ServerInfo{
				Name:    "llm-language-server",
				Version: "1.0.0-0",
			},
		},
	}
}
