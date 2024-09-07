package main

func GenerateAssembly(sourceCode string) (string, error) {
	prompt := "Convert the following C code to assembly:\n" + sourceCode
	assemblyCode, err := QueryChatGPT(prompt)
	if err != nil {
		return "", err
	}
	return assemblyCode, nil
}
