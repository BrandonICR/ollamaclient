# Ollama Client

This client package provides to the consumers an easy way to interact with a ollama server using Go.

## Get Started

Import this client to your project!, just execute:

> go get github.com/BrandonICR/ollamaclient

## Example

If you have running ollama in your local system or in a known remote server, you can try with
the example implementation, only execute in your terminal:

> go run _example/example.go model=deepseek-r1:14b domain=http://localhost:11434

If no arguments are provided by default model=deepseek-r1:14b and domain=http://localhost:11434 are taken.

Once you have run the previous command, you can see the logs in [log](./_example/log)

## Help

If you have not dowloaded Ollama yet in your system see https://ollama.com/