//create a connection to the server, using IPC as a transport
// also include all preview/proposed LSP features

import {
  createConnection,
  ProposedFeatures,
  TextDocuments,
  TextDocumentSyncKind,
  type InitializeParams,
  type InitializeResult,
} from "vscode-languageserver";
import { TextDocument } from "vscode-languageserver-textdocument";

const connection = createConnection(ProposedFeatures.all);

//simple text document manager

const documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

connection.onInitialize((params: InitializeParams) => {
  const result: InitializeResult = {
    capabilities: {
      textDocumentSync: TextDocumentSyncKind.Incremental,
    },
  };

  return result;
});

documents.onDidChangeContent((change) => {
  connection.window.showInformationMessage(
    "onDidChangeContent: " + change.document.uri
  );
});

//make the text document manager listen on the connection
// for open. change and close text document events

documents.listen(connection);

//listen on the connection
connection.listen();
