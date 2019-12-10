package retag

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/yyyiue/protobuf/protoc-gen-go/generator"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	generator.RegisterPlugin(new(retag))
}

type retag struct {
	gen         *generator.Generator
	tags        map[string]string
	fieldMaxLen int
	tagMaxLen   int
}

// Name returns the name of this plugin, "settag"
func (r *retag) Name() string {
	return "retag"
}

// Init initializes the plugin.
func (r *retag) Init(gen *generator.Generator) {
	r.gen = gen
}

func (r *retag) P(args ...interface{}) { r.gen.P(args...) }

// Generate generates code for the services in the given file.
func (r *retag) Generate(file *generator.FileDescriptor) {
	r.getStructTags(*file.Name)
	r.retag()
}

// GenerateImports generates the import declaration for this file.
func (r *retag) GenerateImports(file *generator.FileDescriptor) {}

func (r *retag) getStructTags(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	r.tags = make(map[string]string)
	var comment bool
	reader := bufio.NewReader(file)
	msgNameStack := NewStack()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		//skip empty line in message
		if len(line) <= 0 {
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(string(line)), "/*") {
			comment = true
		}

		if comment && strings.Contains(string(line), "*/") {
			comment = false
			continue
		}

		if comment {
			continue
		}

		//fmt.Println("------", string(line))
		if strings.HasPrefix(strings.TrimSpace(string(line)), "message") {
			if msgNameStack.GetPOP() != "" {
				msgNameStack.PUSH(msgNameStack.GetPOP() + "_" + strings.Fields(string(line))[1])
			} else {
				msgNameStack.PUSH(strings.Fields(string(line))[1])
			}
			continue
		}

		if msgNameStack.GetPOP() != "" && strings.TrimSpace(string(line))[0] == '}' {
			msgNameStack.POP()
			continue
		}

		if msgNameStack.GetPOP() != "" {
			if strings.HasPrefix(strings.TrimSpace(string(line)), "//") {
				continue
			}

			k, v := getFieldAndTag(string(line), msgNameStack.GetPOP())

			if len(strings.Split(k, ".")) <= 1 {
				continue
			}

			// panic(string(line))
			//  int64 number_aa = 1;   //`json:"number123,string"`

			// panic(k)
			//ChildOrder.number_aa
			r.tags[k] = v

			if len(strings.Split(k, ".")[1]) > r.fieldMaxLen {
				r.fieldMaxLen = len(strings.Split(k, ".")[1])
			}

			tags := strings.Fields(v)
			for _, tag := range tags {
				if len(strings.Split(tag, ":")[1])-2 > r.tagMaxLen {
					r.tagMaxLen = len(strings.Split(tag, ":")[1]) - 2
				}
			}
		}
	}

	//for k, v := range r.tags {
	//	fmt.Println(k, v)
	//}
	/**
	ChildOrder.Number->json:"number123,string"
	ChildOrder.clazz_number->json:"clazzNumber,int"
	ChildOrder.clazz_category->json:"clazzCategory1,int"
	**/
}

// line 字段所在proto文件中的一行
// msgName 字段所属message
func getFieldAndTag(line string, msgName string) (field string, tag string) {
	fts := strings.Split(line, "//")
	if len(fts) <= 1 {
		return "", ""
	}

	tag = fts[1]
	fs := strings.Fields(fts[0])
	fsl := len(fs)
	field = msgName + "."
	for i := 0; i < fsl; i++ {
		if i == fsl-1 {
			field += fs[i]
			break
		}

		if fs[i+1] == "=" {
			flag := false
			for k, c := range fs[i] {
				if k == 0 {
					field += strings.ToUpper(string(c))
					continue
				}
				if c == '_' {
					flag = true
					continue
				}
				if flag {
					field += strings.ToUpper(string(c))
					flag = false
				} else {
					field += string(c)
				}
			}
			break
		}
	}

	tag = strings.TrimSpace(tag)
	tag = strings.Trim(tag, "`")
	tag = trimInside(tag)

	return
}

func trimInside(s string) string {
	for {
		if strings.Contains(s, "  ") {
			s = strings.Replace(s, "  ", " ", -1)
		} else {
			break
		}
	}

	return s
}

func (r *retag) retag() {
	if len(r.tags) <= 0 {
		return
	}

	readbuf := bytes.NewBuffer([]byte{})
	// r.gen.Buffer.Bytes() 生成器生成的 Go 代码
	readbuf.Write(r.gen.Buffer.Bytes())
	buf := bytes.NewBuffer([]byte{})

	reader := bufio.NewReader(readbuf)
	var comment bool
	msgNameStack := NewStack()
	for {
		// 不断读取生成的 Go 代码
		line, _, err := reader.ReadLine()
		if err != nil {
			buf.WriteString("\n")
			break
		}

		if strings.HasPrefix(strings.TrimSpace(string(line)), "/*") {
			comment = true
		}

		if comment && strings.Contains(string(line), "*/") {
			comment = false
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if comment {
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if r.needRetag(strings.TrimSpace(string(line))) {
			if msgNameStack.GetPOP() != "" {
				msgNameStack.PUSH(msgNameStack.GetPOP() + "_" + strings.Fields(string(line))[1])
			} else {
				msgNameStack.PUSH(strings.Fields(string(line))[1])
			}

			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if msgNameStack.GetPOP() != "" && strings.TrimSpace(string(line))[0] == '}' {
			msgNameStack.POP()
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if msgNameStack.GetPOP() != "" {
			if strings.HasPrefix(strings.TrimSpace(string(line)), "//") {
				buf.Write(line)
				buf.WriteString("\n")
				continue
			}

			//panic(string(line))
			//NumberAa int64   `protobuf:"varint,1,opt,name=number_aa,json=numberAa,proto3" json:"number_aa,omitempty"`
			fields := strings.Fields(strings.TrimSpace(string(line)))
			key := msgNameStack.GetPOP() + "." + fields[0]
			//panic(key)
			//ChildOrder.NumberAa
			tag := r.tags[key]
			newline := resetTag(string(line), fields[0], tag, r.fieldMaxLen, r.tagMaxLen)
			buf.WriteString(newline)
			buf.WriteString("\n")
			continue
		}
		buf.Write(line)
		buf.WriteString("\n")
	}

	r.gen.Buffer.Reset()
	// data 为要生成的 .go 文件内容
	data := buf.Bytes()
	r.gen.Buffer.Write(data)
}

func (r *retag) needRetag(line string) bool {
	for k := range r.tags {
		ks := strings.Split(k, ".")
		sub := "type " + ks[0] + " struct"
		if strings.HasPrefix(line, sub) {
			return true
		}
	}

	return false
}

func resetTag(line string, field string, tag string, maxlenField, maxlenTag int) string {
	//reset default json
	res := strings.Trim(strings.TrimRight(strings.TrimRight(line, "\n"), " "), "`")
	if strings.Contains(line, "json:") && strings.Contains(tag, "json:") {
		r := regexp.MustCompile(` json:"[\w]*,omitempty"`)
		res = r.ReplaceAllString(res, "")
		//substr := " json:\"" + field + ",omitempty\""
		//res = strings.Replace(res, substr, "", -1)
	}

	fs := strings.Fields(res)
	for i := 2; i < len(fs); i++ {
		if i == 2 {
			res = strings.Replace(res, fs[i], "`", -1)
			fs[i] = fs[i][1:]
		} else {
			res = strings.Replace(res, fs[i], "", -1)
		}
	}

	fs = append(fs, strings.Fields(tag)...)

	for i := 2; i < len(fs); i++ {
		if i == 2 {
			format := "%-" + strconv.Itoa(len(`protobuf:"bytes,xxx,opt,name=`)+maxlenField) + "s "
			res += fmt.Sprintf(format, fs[i])
		} else if i != len(fs)-1 {
			format := "%-" + strconv.Itoa(len(fs[i])-len(strings.Trim(strings.Split(fs[i], ":")[1], "\""))+maxlenTag) + "s  "
			res += fmt.Sprintf(format, fs[i])
		} else {
			res += fs[i]
		}
	}

	res += "`"
	return res
}
