package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"sigs.k8s.io/yaml"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read file: %s\n", os.Args[1])
		os.Exit(1)
	}
	var output bytes.Buffer
	//var f interface{}
	var f map[string]interface{}
	//var f secretData
	ferr := yaml.Unmarshal(content, &f)
	if ferr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error unmarshalling manifest content: %q \n%s\n", ferr, content)
		os.Exit(1)
	}

	delete(f, "sops")
	//fmt.Printf("--- f:\n%v\n\n", f)

	d, err := yaml.Marshal(&f)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %q", err)
	}

	//fmt.Println(string(d))
	fmt.Println(reflect.TypeOf(f["stringData"]))
	fmt.Println()
	output.Write(d)
	output.WriteString("\n---\n")
	_, _ = fmt.Fprintf(os.Stdout, output.String())

}

/*package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

func main() {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
}*/
