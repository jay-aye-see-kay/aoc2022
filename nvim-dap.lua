local dap = require("dap")

dap.adapters.delve = {
	type = "server",
	port = "${port}",
	executable = {
		command = "dlv",
		args = { "dap", "-l", "127.0.0.1:${port}" },
	},
}

-- https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md
dap.configurations.go = {
	{
		type = "delve",
		name = "Debug",
		request = "launch",
		program = "${file}",
	},
	{
		type = "delve",
		name = "Debug test", -- configuration for debugging test files
		request = "launch",
		mode = "test",
		program = "${file}",
	},
	-- works with go.mod packages and sub packages
	{
		type = "delve",
		name = "Debug test (go.mod)",
		request = "launch",
		mode = "test",
		program = "./${relativeFileDirname}",
	},
}

require("dapui").setup({})
require('dap-go').setup()
require("nvim-dap-virtual-text").setup({})

vim.keymap.set("n", "<space>dR", require("dap").clear_breakpoints, { desc = "clear breakpoints" })
vim.keymap.set("n", "<space>db", require("dap").toggle_breakpoint, { desc = "toggle breakpoint" })
vim.keymap.set("n", "<space>dc", require("dap").continue, { desc = "continue" })
vim.keymap.set("n", "<space>dr", require("dap").repl.open, { desc = "repl" })
vim.keymap.set("n", "<space>di", require("dap").step_into, { desc = "step into" })
vim.keymap.set("n", "<space>do", require("dap").step_over, { desc = "step over" })
vim.keymap.set("n", "<space>dt", require("dap").step_out, { desc = "step out" })

vim.keymap.set("n", "<space>du", require("dapui").toggle, { desc = "toggle" })
vim.keymap.set("n", "<space>dc", require("dap").run_to_cursor, { desc = "run to cursor" })

