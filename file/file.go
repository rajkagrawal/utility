package file

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
)
// os.O_WRONLY|os.O_APPEND :: If this is ued then file is used to append
// os.O_WRONLY|os.O_APPEND|os.O_TRUNC : if this is used then it first removes all the contents and then adds/appends


func IsFileExist(path string) bool {
	_, err := os.Open(path)
	if err != nil {
		if _, ok := err.(*os.PathError);ok {
			return false
		}
	}
	return true
}

func CopyFile(src , destination string) error {
	reader , err := os.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()
	writer , err := os.Create(destination)
	if err != nil {
		return err
	}
	defer writer.Close()
	_, err = io.Copy(writer,reader)
	if err!=nil{
		return err
	}
	err = writer.Sync()
	return err
}
func UncompressFile(pathOfCompressedFile , targetFileName string)error{
	// Open compressed file
	gzipFile, err := os.Open(pathOfCompressedFile)
	if err != nil {
		return err
	}

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
	return err
	}
	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer
	outfileWriter, err := os.Create(targetFileName)
	if err != nil {
		return err
	}
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	return err
}
func CompressFile(path string, compressFileName string ) error {
	file, err := os.Open(path)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	// Read file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	file.Close()
	write, err := os.Create(compressFileName)
	if err != nil {
		return err
	}
	x := gzip.NewWriter(write)
	defer x.Close()
	_, err = x.Write(data)
	return err

}