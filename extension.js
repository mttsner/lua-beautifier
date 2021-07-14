const { languages, Range, TextEdit } = require("vscode");
const { readFileSync } = require('fs');
const { resolve } = require('path');
require('./wasm_exec.js');

module.exports = {
	activate: context => {
		const go = new Go()
		const file = readFileSync(resolve(__dirname, "main.wasm"))
		const buf = new Uint8Array(file)

		WebAssembly.instantiate(buf, go.importObject)
		.then(result => {
			go.run(result.instance)
			context.subscriptions.push(languages.registerDocumentFormattingEditProvider("lua", {
				provideDocumentFormattingEdits: document => [TextEdit.replace(new Range(document.lineAt(0).range.start, document.lineAt(document.lineCount - 1).range.end), Beautify(document.getText()))]
			}))
		})
	},
	deactivate: () => { }
};