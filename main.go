package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"sigs.k8s.io/yaml"
)

func main() {
	ksopsEnv, _ := strconv.ParseBool(os.Getenv("KSOPS_ENV"))
	if ksopsEnv {
		fmt.Println("this is true")
	} else {
		fmt.Println("this is false")
	}

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
	for k, v := range f {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		if k == "stringData" || k == "data" {
			//fmt.Println(reflect.TypeOf(v))
			tempList := make(map[string]string)
			for a, _ := range v.(map[string]interface{}) {
				tempList[a] = "SECRET"
			}
			f[k] = tempList
		}
	}

	d, err := yaml.Marshal(&f)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %q", err)
	}

	//fmt.Println(string(d))
	fmt.Println(reflect.TypeOf(d))

	n := testA(10)
	fmt.Println(n)

	output.Write(d)
	output.WriteString("\n---\n")
	_, _ = fmt.Fprintf(os.Stdout, output.String())

}

func testA(x int) int {
	x = x * 2
	return x
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
