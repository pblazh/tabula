local M = {}

M.setup = function()
	vim.api.nvim_create_autocmd("FileType", {
		desc = "Execute tabula on csv files",
		pattern = "csv",
		group = vim.api.nvim_create_augroup("tabula", { clear = true }),
		callback = function(opts)
			local csvview = require("csvview")

			vim.cmd("set foldmethod=marker")
			vim.cmd("set foldlevel=0")

			if not csvview.is_enabled(0) then
				csvview.enable(0)
			end

			vim.o.autoread = true

			vim.api.nvim_create_autocmd({ "BufWritePost", "InsertLeave" }, {
				buffer = opts.buf,
				group = vim.api.nvim_create_augroup("tabula_save", { clear = true }),
				callback = function()
					local stderr = vim.uv.new_pipe()

					local chunks = {}

					function OnExit(code, signal)
						if code == 0 then
							vim.schedule(function()
								vim.cmd("e %")
								-- vim.cmd("1messages")
							end)
						else
							local out = table.concat(chunks, "")
							print(out)
							vim.schedule(function()
								vim.cmd("e %")
								-- vim.cmd("1messages")
							end)
						end
					end

					vim.cmd(":write!")

					vim.uv.spawn("tabula", {
						stdio = { nil, nil, stderr },
						args = { "-a", "-u", vim.api.nvim_buf_get_name(0) },
					}, OnExit)

					if not stderr then
						return
					end

					stderr:read_start(function(err, chunk)
						if err then
							print(err)
						elseif chunk then
							table.insert(chunks, chunk)
						end
					end)
				end,
			})
		end,
	})
	--
end

-- Map a command to the function
-- vim.api.nvim_command('command! HelloWorld lua require("tabula").Tabula()')

return M
