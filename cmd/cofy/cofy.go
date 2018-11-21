package main

func main() {
	rootCMD.AddCommand(versionCMD)
	rootCMD.Execute()
}
