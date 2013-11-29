package main


import(
	"strings"
	"fmt"
	"path/filepath"
	"os"
)

func source(files []string) <-chan string {
    out := make(chan string, 1000)
    go func() {
        for _, filename := range files {
            out <- filename
        }
        close(out)
    }()
    return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
    out := make(chan string, cap(in))
    go func() {
        for filename := range in {

            if len(suffixes) == 0 {
                out <- filename
                continue
            }
            ext := strings.ToLower(filepath.Ext(filename))
            for _, suffix := range suffixes {

                if ext == suffix {
					
                    out <- filename
                    break
                }
            }
        }
        close(out)
    }()
    return out
}

func sink(in <-chan string) {
    for filename := range in {
        fmt.Println(filename)
    }
}

func filterSize(minSize int,maxSize int, in <-chan string) <-chan string {
    out := make(chan string, cap(in))
    go func() {
        for filename := range in {
			
			fileinfo,err := os.Stat(filename)
			if err != nil{
				continue
			}
			if fileinfo.Size() >= int64(minSize) && fileinfo.Size() <= int64(maxSize){
                out <- filename
                continue
            }
        }  
        close(out)
    }()
    return out
}

func handleCommandLine()(int,int,[]string,[]string){
	a1 := []string{".go"}
	a2 := []string{"c:/GoSrc/src/test/concurrent/main.go"}
	return 0,10000,a1,a2
}

func main(){
	minSize, maxSize, suffixes, files := handleCommandLine()

	channel1 := source(files)
	channel2 := filterSuffixes(suffixes, channel1)
	channel3 := filterSize(minSize, maxSize, channel2)
	sink(channel3)
}









