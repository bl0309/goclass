package examplejson

import (
	"bufio"
	"fmt"
	"os"
)

func InitFileWrite() {

	data := []byte("People are good enough to understand what they do!\n")

	// data := ReadFile("./example_json/person_data.json")
	err := os.WriteFile("./example_json/writeText.txt", data, 0644)
	checkError(err)

	//if condition file existance

	f, err := os.Create("./example_json/writeText2.txt")
	checkError(err)
	defer f.Close()

	// data2 := []byte{12, 34, 53, 53, 22, 45, 66, 54}

	data2 := "People are fake!\n"
	n, err := f.Write([]byte(data2))
	checkError(err)
	fmt.Printf("Wrote %d byte\n", n)

	n2, err := f.WriteString("But humnan beings are real...\n")
	checkError(err)
	fmt.Printf("Wrote %d bytes\n", n2)
	f.Sync()

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("This line is added using bufio\n")
	checkError(err)
	fmt.Printf("Wrote using bufio %d bytes in new line\n", n3)
	w.Flush()

	// TextWriter w1 = new StreamWriter("./example_json/writeText2.txt", true);
	// w1.WriteLine("This one is new one.\n");
	// w1.Close();
}
