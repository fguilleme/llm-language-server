# LLM Language Server

Implementation of a language server for APIs such as Ollama or Codestral, using Go and the Language Server protocol.

The goal is to have a similar handler as Github Copilot LSP and make it easier to handle multiple providers when working on completion plugins.

NEOVIM inline completion support
Since 0.12+ neovim has builtin inline completion

in {runtimepath/lsp} add ll.lua containing

return {
cmd = { "llm-language-server-1.0.1" },
filetypes = { "lua", "c", "cpp", "python" },
root_markers = { ".git" },

    init_options = {
        provider = "ollama",
        params = {
            url = "http://localhost:11434/api",
            model = "qwen2.5-coder:0.5b",
            model_params = {
                num_ctx = 32768,
                num_predict = -1,
                temperature = 0,
                top_p = 0.95,
                max_tokens = 64,
                stop = { "<||fim_pad|>", "<|endoftext|>", "\n\n", "```" },
            },
        },
    },
    settings = {},

}

Activate this lsp with
:lua vim.lsp.enable('llm')

and activate inline completion with
:lua vim.lsp.inline_completion.enable()

:LspInfo should show that inline completion is enabled in the active buffer

Depending on your configuration all this could be done automatically as well as
mapping <TAB> to accept completion
