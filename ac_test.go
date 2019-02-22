package ac

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

func TestAC_Filter(t *testing.T) {

	dict := map[string]int{
		"尼玛":1,
		"尼玛比":1,
		"叼你":1,
		"狗比":1,
		"38":1,
		"SB":1,
		"猪":1,
		//"，":2,
	}


	ac := FromDict(dict)
	input := "叼你啊,尼玛比是个猪啊,狗比，38就-是SB                                                                                                                                                                                                                                                                                                                   "
	t.Log(len(input))
	out := ac.Filter(input)
	t.Log(len(out),out)
}

func TestFromFile(t *testing.T) {
	ac := FromFile("dict.txt")

	input := "一辈子的孤单是得呀，候鸟de 倒霉命运；特工小子就是你"
	output := ac.Filter(input)
	t.Log("len:",len(output),"out:",output)

	//t.Log(http.ListenAndServe("localhost:10000", nil))

}


var globalInput string
func TestMain(m *testing.M) {

	globalInput = prepareData("dict.txt",10)
	os.Exit(m.Run())
}

func TestMutilStr(t *testing.T) {
	ac := FromFile("dict.txt")


	t.Log("input 字符数： ",len([]rune(globalInput)),"字节数： ",len(globalInput),"循环次数： ",1e5)

	
	for i:=0;i<1e5;i++{
		ac.Filter(globalInput)
		//output := ac.Filter(globalInput)
		//t.Log("len:",len(output),"out:",output)

	}



}


func prepareData(file string,len int) string {
	dict := make(map[int]string, 1115)
	f, err := os.OpenFile(file, os.O_RDONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	i := 0
	for {
		l, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		piece := bytes.Split(bytes.TrimSpace(l), []byte("\t"))
		key := string(piece[0])

		dict[i] = key
		i++
	}

	str:= ""
	i = 0
	for _,v := range dict {
		if i > 100 {
			break
		}

		if i%2 == 0 {
			str += v
		} else {
			str+= "#"
		}
		i++
	}

	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	//for i:=0 ;i<len;i++{
	//	//r := rnd.Intn(len(dict)*2)
	//	//if r >= len(dict) {
	//	//	str += "#"
	//	//} else {
	//	//	str += dict[r]
	//	//}
	//
	//	if i%2 == 0 {
	//		str += dict[i]
	//	} else {
	//		str+= "#"
	//	}
	//
	//
	//}
	return str
}

func BenchmarkAC_Filter(b *testing.B) {

	b.StopTimer()
	input := prepareData("dict.txt" ,100)
	b.Log("input 字符数： ",len([]rune(input)),"字节数： ",len(input),"循环次数： ",b.N)

	ac := FromFile("dict.txt")

	b.StartTimer()


	for i:=0;i<b.N;i++{
		ac.Filter(input)

	}
	//output := ac.Filter(input)
	//
	//b.Log("len:",len(output),"out:",output)
}


func BenchmarkAC_FilterUgly(b *testing.B) {

	b.StopTimer()
	input := prepareData("dict-ugly-out.txt",100)
	b.Log("input 字符数： ",len([]rune(input)),"字节数： ",len(input),"循环次数： ",b.N)

	ac := FromFile("dict-ugly-out.txt")

	b.StartTimer()



	for i:=0;i<b.N;i++{
		ac.Filter(input)

	}
	//output := ac.Filter(input)

	//b.Log("len:",len(output),"out:",output)



}

//func TestPProf(t *testing.T) {
//	input := prepareData("dict-ugly-out.txt")
//	ac := FromFile("dict-ugly-out.txt")
//	for {
//
//		t := time.After(1*time.Minute)
//
//		select {
//		case <-t :
//			break
//		default:
//
//		}
//
//
//	}
//}



//func BenchmarkNil(b *testing.B) {
//	b.StopTimer()
//
//	time.Sleep(1*time.Second)
//	b.StartTimer()
//
//	for i:=0; i< b.N ; i++{
//		time.Sleep(1*time.Second)
//
//	}
//
//
//}

func TestFixFile(t *testing.T) {
	f, err := os.OpenFile("dict-ugly.txt", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	// output
	outputFile, err := os.Create("dict-ugly-out.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	writer := bufio.NewWriterSize(outputFile, 1024)

	// process loop
	//var line []byte
	//for err = nil; err != io.EOF; line, err := r.ReadSlice('\n') {
	//	fmt.Println(string(line))
	//	line = append(line,bt...)
	//	fmt.Println(string(line))
	//	writer.Write(line)
	//	line = nil
	//}

	//bt := []byte{'\t','1','\n'}

	dict := []string{}
	for {

		 line, err := r.ReadSlice('\n')
		 if err == io.EOF {
			break
			}
		l := bytes.TrimSpace(line)

		//fmt.Println(string(append(l,bt...)))
		dict = append(dict,string(l))
	}

	for _,v := range dict {
		v += "\t1\n"
		writer.Write([]byte(v))
	}
	writer.Flush()

}