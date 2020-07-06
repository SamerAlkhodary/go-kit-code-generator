package gen

type fileGenerator interface {
	generateFile(poutputPath string)
	generateCode()
	run(outputPath string)
	GetFileName() string
}
