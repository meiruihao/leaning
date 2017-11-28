package main

import (
	"fmt"
	_ "go/types"
	//"io/ioutil"
	//"net/http"
	"math"
	"errors"
	//"reflect"
	//"time"
	"time"
	//"sync/atomic"
	//"sort"
	//"os"
	"os"
	//"crypto/rand"
	//s "strings"
	//"regexp"
	//"encoding/json"
	//"strconv"
	//"net/url"
	//"net"
	//"io/ioutil"
	//"flag"
	//"strings"
	//"os/exec"
	//"syscall"
	"log"
	"net/http"
	"io"
)

//定义一个结构体 类似与实体类
type Car struct {
	color string
	size  int
	logo string
}

//定义一个求阶层方法
func Factori(n uint64)(result uint64){
	if(n > 0){
		result = n * Factori(n-1)
		return result
	}
	return 1
}

//定义一个求斐波那契方法
func feibonaqi(n int) int{
	if(n < 2){
		return n;
	}
	return feibonaqi(n - 2) + feibonaqi(n-1)
}

//定义接口

type Phone interface {
	call()
}

type SANSANGPhone struct {

}

type IPhone struct {

}

func (sPhone  SANSANGPhone)  call(){
	fmt.Println("SANSANGPhoe call")
}

func(iphone IPhone) call(){
	fmt.Println("IPhone call")
}

func plus(a int,b int) int{
	return a + b
}

func sum(nums ...int)  {
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Print(total)
}

type geomtry interface {
	area() float64
	perim() float64
}

type rect struct {
	width,height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64  {
	return r.width*r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle)area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle)  perim() float64{
	return 2 * math.Pi * c.radius
}

func measure(g geomtry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type tixing struct {
	shangdi,xiadi,gao float64
}

func (t tixing) area() float64  {
	return (t.shangdi+t.xiadi)*t.gao / 2
}

func (t tixing) perim() float64  {
	return (t.shangdi+t.xiadi) * 2
}

/**
异常处理
 */
func f1(arg int)(int,error)  {
	if arg == 42 {
		return -1,errors.New("can't work with 42")
	}
	return arg+3,nil
}

type argError struct {
	arg int
	prop string
}

func (e *argError)Error() string  {
	return fmt.Sprintf("%d - %S",e.arg,e.prop)
}

func f2(arg int)(int ,error)  {
	if arg == 42 {
		return -1,&argError{arg,"cant work it"}
	}
	return arg + 3,nil
}

func f(from string)  {
	for i := 0;i < 3 ; i++  {
		fmt.Println(from,":",i)
	}
}

func ping(pings chan <- string ,msg string)  {
	pings <- msg
}

func pong(pings <- chan string , pongs chan  <- string)  {
	msg := <- pings
	pongs <- msg
}

func worker(id int,jobs <- chan int,results chan <- int)  {
	for j:= range jobs{
		fmt.Println("worker",id,"started job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)
		results <- j * 2
	}
}

type ByLength []string

func (s ByLength) Len() int{
	return len(s)
}

func (s ByLength) Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool  {
	return len(s[i]) < len(s[j])
}


func createFile(p string) *os.File  {
	fmt.Println("creating...")
	f,err := os.Create(p)
	if(err != nil){
		panic(err)
	}
	return f
}

func write(f  *os.File)  {
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}

func closeFile(f  *os.File)  {
	fmt.Println("closing...")
	f.Close()
}

var p = fmt.Println

func main() {
	fmt.Println("Hello,GO!")

	//var  name  =  "菜鸟";
	//fmt.Println(name);
	//var a int = 1;
	//var b int  = 2;
	//fmt.Println(a + b);
	//if(a == b){
	//	fmt.Println("a 等于 b")
	//}else {
	//	fmt.Println("a 不等于 b")
	//}

	//var n [10]int
	//var i,j int
	//for i = 0; i < 10; i++  {
	//	n[i] = i+ 100;
	//}
	//
	//for j = 0; j < 10; j++  {
	//	fmt.Println("ssss",j,n[j]);
	//}

	//var benzi Car
	//var bmw Car
	//
	//benzi.color = "red"
	//benzi.logo = "benzi"
	//benzi.size = 2
	//
	//
	//bmw.color = "blue"
	//bmw.size = 3
	//bmw.logo = "bmw"
	//
	//fmt.Print(bmw)

	//var pepples []string = make([]string,3)
	//
	//pepples[0] = "p1"
	//pepples[1] = "p2"
	//pepples[2] = "p3"
	//
	//fmt.Println(pepples)
	//
	//pepples2 := pepples[0:2]
	//fmt.Println(pepples2)

	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	//nums := []int{2 , 3 , 4}
	//sum := 0
	//for _, num := range nums {
	//	sum += num
	//}
	//fmt.Println("sum:", sum)


	/* 创建 map */
	//var testMap = make(map[string] string)
	//testMap["age"] = "12"
	//testMap["sex"] = "MAN"
	//testMap["weight"] = "65"

	//遍历map person 为 key testMap[person] 为 value

	//for person := range testMap{
	//	fmt.Println("人的：",person,"是",testMap[person])
	//}

	/* 打印 map */
	//person , isExist := testMap["sex"]
	//if(isExist){
	//	fmt.Println("person of sex is",person)
	//}else {
	//	fmt.Println("person of sex is not present")
	//}

	/* 删除元素 */
	//delete(testMap,"sex")
	//fmt.Println("key for sex is deleted")
	//
	//fmt.Println("new map is ",testMap)

	//调用一个求阶层方法
	//var i int = 15
	//fmt.Println("15的阶层是：",Factori(uint64(i)))

	//调用一个求斐波那契方法
	//var  i int
	//for i = 0; i < 20;i++ {
	//	fmt.Print("\t", feibonaqi(i))
	//}

	//var sum int = 17
	//var count int = 5
	//var mean float32
	//
	//mean = float32(sum)/float32(count)
	//
	//fmt.Println(mean)



	//var phone Phone
	//
	//phone = new(SANSANGPhone)
	//phone.call()
	//
	//phone = new (IPhone)
	//phone.call()


	//获取的百度首页的HTML源文件
	//response,_:=http.Get("https://www.ljldata.com")
	//defer response.Body.Close()
	//body,_:=ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	//numbers := []int{0,1,2,3}
	//for i:= range numbers  {
	//	fmt.Println("Slice item",i,"is",numbers[i])
	//}

	//countryMap := map[string] string {"France":"Paris","China":"Beijing"}
	//for contry := range countryMap {
	//	fmt.Print(contry,countryMap[contry])
	//}

	//for country ,capital :=range countryMap{
	//	fmt.Print(country,"aaaaa",capital)
	//}

	//const name  =  "abd"
	//fmt.Print(name)


	// for 循环
	//i:= 3
	//for i< 10 {
	//	fmt.Println(i)
	//	i = i + 1
	//}

	//for j := 7;j <= 9 ; j++  {
	//	fmt.Println(j)
	//}

	//switch case
	//t := time.Now()
	//switch {
	//case t.Hour() <12:
	//	fmt.Print("is before noon")
	//default:
	//	fmt.Print("is after noon")
	//}
	//
	//switch t.Hour() {
	//case 15:
	//	fmt.Print("is now")
	//default:
	//	fmt.Print("is not now")
	//}

	//nums := []int{1,2,3}
	//sum := 0
	//for _, num := range nums {
	//	sum+=num
	//}
	//fmt.Print(sum)

	//for i, num :=range nums {
	//	if num == 3 {
	//		fmt.Print("index:", i)
	//	}
	//}


	//maps := map[string] string{"age":"1","sex":"man"}
	//for k,v := range maps{
	//	fmt.Printf(k,v)
	//}


	//fmt.Print(plus(1,2))

	//nums := []int{1,2,3,4}
	//sum(nums ...)


	//r:= rect{width:3,height:4}
	//c:= circle{radius:5}
	//measure(r)
	//measure(c)

	//t:=tixing{shangdi:2,xiadi:3,gao:5}
	//measure(t)

	//for _, i :=range []int{7 , 42} {
	//	if r,e := f2(i); e != nil {
	//		fmt.Println("f2 faild:",e)
	//	}else {
	//		fmt.Println("f2 faild:",r)
	//	}
	//}

	//for _, i := range []int{7,42}{
	//	if r,e := f2(i); e != nil{
	//		fmt.Println("f2 failed:",e)
	//	}else {
	//		fmt.Println("f2 worked:",r)
	//	}
	//}


	//f("sss")

	//go f("abc")

	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")

	//Go通道实例
	//mes := make(chan string)
	//go func(){mes <- "ping"}()
	//
	//msg := <- mes
	//fmt.Println(msg)

	//Go通道缓冲机制实例
	//mes := make(chan string,3)
	//mes <- "a"
	//mes <- "b"
	//mes <- "c"
	//
	//fmt.Println(<-mes)
	//fmt.Println(<-mes)
	//fmt.Println(<-mes)

	//Go通道路线实例
	//pings := make(chan string,1)
	//ping(pings,"passed message")
	//pongs := make(chan string,1)
	//pong(pings,pongs)
	//fmt.Print(<- pongs)


	//c := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(time.Second * 1)
	//	c <- "one"
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	c2 <- "two"
	//}()
	//
	//for i :=  0; i < 2 ;i++{
	//	select {
	//	case m1 := <- c:
	//		fmt.Println("reee",m1)
	//	case m2 := <- c2:
	//		fmt.Println("reee",m2)
	//	}
	//}


	//message := make(chan string)
	////signal := make(chan string)
	//
	//select {
	//case msg := <- message:
	//	fmt.Println("aaaa",msg)
	//	default:
	//	fmt.Println("no aaaa")
	//}

	/**
	定时器
	 */
	//timer := time.NewTimer(time.Second * 2)
	//<- timer.C
	//fmt.Println("timer 1")


	//jobs := make(chan int,100)
	//results := make(chan int,100)
	//for w:= 1; w<= 4 ;w++{
	//	go worker(w, jobs, results)
	//}
	//
	//for j := 1; j <= 5; j++ {
	//	jobs <- j
	//}
	//close(jobs)
	//
	//
	//for a := 1; a <= 5; a++ {
	//	<-results
	//}


	//requests := make(chan int,5)
	//for i := 1; i <= 5; i++{
	//	requests <- i
	//}
	//close(requests)
	//
	//limiter := time.Tick(time.Millisecond * 200)
	//
	//for req := range requests{
	//	<- limiter
	//	fmt.Println("request",req,time.Now())
	//}

	//var  ops uint64 = 0
	//for i := 0;i < 50 ; i++  {
	//	go func() {
	//		for{
	//			atomic.AddUint64(&ops,1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//opsFinal := atomic.LoadUint64(&ops)
	//fmt.Println("ops:", opsFinal)

	//排序
	//strs := []string{"c","b","a"}
	//sort.Strings(strs)
	//fmt.Print(strs)
	//
	//s := sort.StringsAreSorted(strs)
	//fmt.Println("sorted:",s)

	//fruits := []string{"peach", "banana", "kiwi"}
	//sort.Sort(ByLength(fruits))
	//fmt.Println(fruits)


	//错误处理
	//panic("a error")
	//_,err := os.Create("/temp/file")
	//if err != nil{
	//	panic(err)
	//}



	//延迟(defer)
	//f := createFile("defer.txt")
	//defer closeFile(f)
	//write(f)


	//var strs = []string{"a","b","c","d"}
	////fmt.Println(Index(strs,"a"))
	//
	//fmt.Println(Include(strs,"a"))

	//字符串函数
	//p("Contains:  ", s.Contains("test", "es"))
	//p("Count:     ", s.Count("test", "t"))
	//p("HasPrefix: ", s.HasPrefix("test", "te"))
	//p("HasSuffix: ", s.HasSuffix("test", "st"))
	//p("Index:     ", s.Index("test", "e"))
	//p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	//p("Repeat:    ", s.Repeat("a", 5))
	//p("Replace:   ", s.Replace("foo", "o", "0", -1))
	//p("Replace:   ", s.Replace("foo", "o", "0", 1))
	//p("Split:     ", s.Split("a-b-c-d-e", "-"))
	//p("ToLower:   ", s.ToLower("TEST"))
	//p("ToUpper:   ", s.ToUpper("test"))
	//p()


	//字符串格式化
	//p := poit{1,2}
	//fmt.Printf("%v\n",p)
	//
	//fmt.Printf("%+v\n", p)

	//正则表达式
	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//fmt.Println(match)



	//JSON实例
	//bolB, _ := json.Marshal(true)
	//fmt.Println(string(bolB))

	//intB, _ := json.Marshal(1)
	//fmt.Println(string(intB))

	//fltB,_ := json.Marshal(2.34)
	//fmt.Println(string(fltB))

	//slcD := []string{"apple","pear","peach"}
	//slcB,_ := json.Marshal(slcD)
	//fmt.Println(string(slcB))

	//mapD := map[string]int {"a":1,"b":2}
	//mapC,_ := json.Marshal(mapD)
	//fmt.Println(string(mapC))

	//res1D := &res1{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res1B, _ := json.Marshal(res1D)
	//fmt.Println(string(res1B))

	//res1D := &res2{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res1B, _ := json.Marshal(res1D)
	//fmt.Println(string(res1B))


	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//var dat map[string]interface{}
	//
	//if err := json.Unmarshal(byt, &dat); err != nil {
	//	panic(err)
	//}
	//fmt.Println(dat)
	//
	//num := dat["num"].(float64)
	//fmt.Println(num)
	//
	//strs := dat["strs"].([]interface{})
	//str1 := strs[0].(string)
	//fmt.Println(str1)



	//时间
	//p := fmt.Println
	//now := time.Now()
	//p(now)
	//
	//then := time.Date(2001,11,12,20,34,58,651387237,time.UTC)
	//p(then)
	//p(now.Year())
	//p(now.Month())
	//p(now.Day())
	//p(now.Hour())
	//p(now.Weekday())
	//p(now.Before(now))
	//p(now.After(now))
	//p(now.Equal(now))

	//时代（Epoch）
	//now := time.Now()
	//secs := now.Unix()
	//nanos := now.UnixNano()
	//fmt.Println(now)
	//
	//millis := nanos / 1000000
	//fmt.Println(secs)
	//fmt.Println(millis)
	//fmt.Println(nanos)

	//时间格式化/解析
	//p := fmt.Println
	//
	//t := time.Now()
	//p(t.Format(time.RFC3339))
	//p(t.Format("3:04PM"))
	//p(t.Format("Mon Jan _2 15:04:05 2017"))
	//form := "3 04 PM"
	//t2, e := time.Parse(form, "8 41 PM")
	//p(t2)
	//ansic := "Mon Jan _2 15:04:05 2006"
	//_, e = time.Parse(ansic, "8:41PM")
	//p(e)

	//数字解析
	//p := fmt.Println
	//f,_ := strconv.ParseFloat("1.2345",64)
	//p(f)
	//i,_ := strconv.ParseInt("123",0,64)
	//p(i)

	//d,_ := strconv.ParseInt("0x1c8",0,64)
	//p(d)

	//u,_ := strconv.ParseUint("789",0,64)
	//p(u)

	//k,_ := strconv.Atoi("abc")
	//p(k)


	//URL解析
	//p := fmt.Println
	//
	//s := "https://192.168.1.1:1234/business"
	//u,err := url.Parse(s)
	//if err != nil {
	//	panic(err)
	//}
	//p(u.Scheme)
	//p(u.Host)
	//
	//host,port,_ := net.SplitHostPort(u.Host)
	//p(host)
	//p(port)
	//p(u.Path)
	//
	//p(u.RawQuery)
	//m,_ := url.ParseQuery(u.RawQuery)
	//p(m)
	//p(m["k"][0])

	//base64编码
	//p := fmt.Println
	//data := "abc123!?SDASIH"
	//sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	//p(sEnc)
	//sDec,_ := b64.StdEncoding.DecodeString(sEnc)
	//p(sDec)
	//p(string(sDec))
	//
	//
	//uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	//p(uEnc)
	//uDec,_ := b64.URLEncoding.DecodeString(uEnc)
	//p(uDec)
	//p(string(uDec))


	//读取文件
	//p := fmt.Println
	//dat,err := ioutil.ReadFile("C:/Users/john/Desktop/1.txt")
	//check(err)
	//p(string(dat))

	//f , err := os.Open("C:/Users/john/Desktop/1.txt")
	//check(err)
	//b1 := make([]byte,5)
	//n1,err :=f.Read(b1)
	//check(err)
	//p("%d bytes:",n1,string(b1))
	//
	//o2, err := f.Seek(6, 0)
	//check(err)
	//b2 := make([]byte, 2)
	//n2, err := f.Read(b2)
	//check(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	//
	//wordPtr := flag.String("word", "foo", "a string")
	//p(*wordPtr)

	//环境变量
	//os.Setenv("FOO", "1")
	//p("FOO:", os.Getenv("FOO"))
	//p("BAR:", os.Getenv("BAR"))
	//
	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	fmt.Println(pair[0])
	//}

	//binary ,err := exec.LookPath("ls")
	//if err != nil{
	//	panic(err)
	//}
	//
	//args := []string{"ls","-a","-1","-h"}
	//env := os.Environ()
	//execErr := syscall.Exec(binary,args,env)
	//if execErr != nil{
	//	panic(execErr)
	//}

	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func check(e error)  {
	if e != nil{
		panic(e)
	}

}




type res1 struct {
	Page int
	Fruits []string
}

type res2 struct {
	Page int  `json:"page"`
	Fruits []string  `json:"fruits"`
	
}


type poit struct {
	x,y int
}

func Index(vs []string ,t string) int {
	for i, v := range vs{
		if(v == t){
			return i
		}
	}
	return -1
}

func Include(vs []string,t string) bool {
	return Index(vs,t) >= 0
}

func Any(vs []string,f func(string) bool) bool  {
	for _,v := range vs {
		if(!f(v)){
			return false
		}
	}
	return true
}

func All(vs []string,f func(string) bool) bool  {
	for _,v := range vs{
		if(!f(v)){
			return true
		}
	}
	return false
}




