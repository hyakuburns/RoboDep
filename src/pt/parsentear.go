package pt

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var GREEN string = "\033[38;5;184m"
var ORANGE string = "\033[38;5;202m"
var GRAY string = "\033[38;5;109m"
var RED string = "\033[38;5;160m"
var RESET string = "\033[0m"

//go:embed templates
var tmplFiles embed.FS

func Instructions() {
	fmt.Printf("Please use the program correctly\n")
	fmt.Printf(GREEN + "up" + RESET + " for reading a file and getting the dependencies\n")
	fmt.Printf(GREEN + "git [url]" + RESET + " for adding a url to the deps and cloning\n")
	fmt.Printf(GREEN + "hg [url]" + RESET + " for adding a url to the deps and cloning\n")
}
func fileError(fpath string, err error) {
	if err != nil {
		fmt.Printf("Error handling file %s: %v\n", fpath, err.Error())
		os.Exit(1)
	}
}

func cmdError(stdmix []byte, err error) {
	if err != nil {
		fmt.Println(string(stdmix))
		fmt.Printf("Error executing clone command: %v\n\n", err.Error())
	}
}

func tmplError(err error) {
	if err != nil {
		fmt.Printf("Error with the template: %v\n", err.Error())
	}
}

func sanitizeString(dirty *string) {
	*dirty = strings.ReplaceAll(*dirty, "\n", "")
	*dirty = strings.TrimSpace(*dirty)
}
func FileExistence(patheroo string) bool {
	_, err := os.Stat(patheroo)
	if true == os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func SplitUpLine(stringaroo string) []string {
	strings.Trim(stringaroo, " ")
	strArr := strings.Split(stringaroo, " ")
	return strArr
}
func CheckWS(stringaroo string) bool {
	cleaned := strings.TrimSpace(stringaroo)
	if cleaned == "" {
		return true
	} else {
		return false
	}
}

func cloneGit(path string) {
	clonecmd := []string{"git", "-C", "./robodep", "clone"}
	cmd := exec.Command(clonecmd[0], clonecmd[1], clonecmd[2], clonecmd[3], path)
	stdmix, err := cmd.CombinedOutput()
	cmdError(stdmix, err)

}
func cloneHG(path string) {
	clonecmd := []string{"hg", "clone", "--cwd", "./robodep"}
	cmd := exec.Command(clonecmd[0], clonecmd[1], clonecmd[2], clonecmd[3], path)
	stdmix, err := cmd.CombinedOutput()
	cmdError(stdmix, err)
}

func createDepF(file *os.File) {
	fmt.Printf("Creating dep.robo\n")
	t, err := template.ParseFS(tmplFiles, "templates/dep.robo.tmpl")
	tmplError(err)
	err = t.Execute(file, "")
	tmplError(err)

}

func ParseDeps(fileLoc string) {
	//TODO: make it so it checks for the file and then creates it adding whats needed if it doesnt exist, preferably use a template
	//FIXME:lelelelelellele
	var file *os.File
	var err error
	if FileExistence(fileLoc) {
		file, err = os.Open(fileLoc)
		fileError(fileLoc, err)

	} else {
		file, err = os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0700)
		createDepF(file)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var dstart bool
	for scanner.Scan() {
		if CheckWS(scanner.Text()) {
			continue
		}

		splitStr := SplitUpLine(scanner.Text())
		if splitStr[0] == "depstart" {
			dstart = true
			continue
		}
		if dstart == true && len(splitStr) > 1 {
			tokenize := strings.SplitAfter(splitStr[1], "/")
			name := len(tokenize) - 1
			fex := FileExistence("./robodep/" + tokenize[name])
			if !fex {
				fmt.Printf("Cloning %s...\n", tokenize[name])
				switch splitStr[0] {
				case "git":
					cloneGit(splitStr[1])
				case "hg":
					cloneHG(splitStr[1])
				default:
					continue

				}
			} else {
				fmt.Printf(ORANGE+"Skipping"+RESET+" %s, directory already exists\n", tokenize[name])
			}
		}

	}
}

func gitCheckEQ(combinedArgs string, f *os.File) bool {

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		trimScan := strings.TrimSpace(scanner.Text())
		if trimScan == combinedArgs {
			return true
		}
	}
	return false
}

func GitAddRepo(fileLoc string, argv []string) {
	combinedArgs := argv[1] + " " + argv[2]
	tokenize := strings.SplitAfter(argv[2], "/")
	name := len(tokenize) - 1

	f, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0700)
	fileError(fileLoc, err)
	defer f.Close()
	wbool := gitCheckEQ(combinedArgs, f)
	if wbool {
		fmt.Printf(GREEN+"Cloning"+RESET+" %s...\n", tokenize[name])
		cloneGit(argv[2])
	} else {
		_, err = f.WriteString("git " + argv[2] + "\n")
		if err != nil {
			fmt.Println(RED + "Error" + RESET + " writing to dep.robo")
		}
		fmt.Printf(GREEN+"Cloning"+RESET+" %s...\n", tokenize[name])
		cloneGit(argv[2])
	}

}

func hgCheckEQ(combinedArgs string, f *os.File) bool {

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		trimScan := strings.TrimSpace(scanner.Text())
		if trimScan == combinedArgs {
			return true
		}
	}
	return false
}

func HgAddRepo(fileLoc string, argv []string) {
	combinedArgs := argv[1] + " " + argv[2]
	tokenize := strings.SplitAfter(argv[2], "/")
	name := len(tokenize) - 1
	fbool := false
	if !FileExistence(fileLoc) {
		fbool = true
	}
	f, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0700)
	if fbool {
		createDepF(f)
	}
	fileError(fileLoc, err)
	defer f.Close()
	wbool := hgCheckEQ(combinedArgs, f)
	if wbool {
		fmt.Printf(GREEN+"Cloning"+RESET+" %s...\n", tokenize[name])
		cloneHG(argv[2])
	} else {
		_, err = f.WriteString("hg " + argv[2] + "\n")
		if err != nil {
			fmt.Println(RED + "Error" + RESET + " writing to dep.robo")
		}
		fmt.Printf(GREEN+"Cloning"+RESET+" %s...\n", tokenize[name])
		cloneHG(argv[2])
	}

}
