package utils

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"strings"
	"time"
)

var Misc = NewMisc()

type misc struct{}

func NewMisc() *misc {
	return &misc{}
}

// get map default
func (m *misc) GetMapDefault(mapValue map[string]interface{}, key string, def interface{}) interface{} {
	value, ok := mapValue[key]
	if ok {
		return value
	} else {
		return def
	}
}

// rand string
func (m *misc) RandString(strlen int) string {
	codes := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen := len(codes)
	data := make([]byte, strlen)
	rand.Seed(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63() + rand.Int63())
	for i := 0; i < strlen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

// rand int
func (m *misc) RandInt(strLen int) string {
	codes := "0123456789"
	codeLen := len(codes)
	data := make([]byte, strLen)
	rand.Seed(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63() + rand.Int63())
	for i := 0; i < strLen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

// get local ip
func (m *misc) GetLocalIp() string {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		return "localhost"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

// <LABEL_1203>
// total: <LABEL_843>
// page:  <LABEL_844>
// pagesize: <LABEL_1204>
// url: url<LABEL_1508>{page}<LABEL_594>
// args:
//  order <LABEL_845>，<LABEL_989>，<LABEL_1205>1-6<LABEL_1509>，<LABEL_14>
//  a_count   <LABEL_1206>a<LABEL_19>a<LABEL_1639>，<LABEL_1613>10<LABEL_1831>。
//
func (m *misc) Page(total, page, pagesize int, url string, args ...interface{}) string {
	order := []int{1, 2, 3, 4, 5, 6}
	a_count := 10
	if len(args) >= 1 {
		order = args[0].([]int)
	}
	if len(args) >= 2 {
		a_count = args[1].(int)
	}
	a_num := a_count
	first := "<LABEL_1640>"
	last := "<LABEL_1641>"
	pre := "<LABEL_1642>"
	next := "<LABEL_1643>"
	if a_num%2 == 0 {
		a_num++
	}
	pages := int(math.Ceil(float64(total) / float64(pagesize)))
	curpage := page
	if curpage > pages || curpage <= 0 {
		curpage = 1
	}
	body := `<span class="page_body">`
	prefix := ""
	subfix := ""
	start := curpage - ((a_num - 1) / 2)
	end := curpage + ((a_num - 1) / 2)
	if start <= 0 {
		start = 1
	}
	if end > pages {
		end = pages
	}
	if pages >= a_num {
		if curpage <= (a_num-1)/2 {
			end = a_num
		}
		if end-curpage <= (a_num-1)/2 {
			start -= int(math.Floor(float64(a_num)/float64(2))) - (end - curpage)
		}
	}
	for i := start; i <= end; i++ {
		if i == curpage {
			body += fmt.Sprintf(`<a class="page_cur_page" href="javascript:void(0);"><b>%d</b></a>`, i)
		} else {
			body += fmt.Sprintf(`<a href="%s">%d</a>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", i), 1), i)

		}
	}
	body += "</span>"
	if curpage > 1 {
		prefix = fmt.Sprintf(`<span class="page_bar_prefix"><a href="%s">%s</a><a href="%s">%s</a></span>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", 1), 1), first, strings.Replace(url, "{page}", fmt.Sprintf("%d", curpage-1), 1), pre)
	}
	if curpage != pages {
		subfix = fmt.Sprintf(`<span class="page_bar_subfix"><a href="%s">%s</a><a href="%s">%s</a></span>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", curpage+1), 1), next, strings.Replace(url, "{page}", fmt.Sprintf("%d", pages), 1), last)
	}
	info := fmt.Sprintf(`<span class="page_cur"><LABEL_1832>%d/%d<LABEL_1833></span>`, curpage, pages)
	id := fmt.Sprintf("gsd09fhas9d%d%d%d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
	gostr := fmt.Sprintf(`<script>function ekup(){if(event.keyCode==13){clkyup();}}function clkyup(){var num=document.getElementById('%s').value;if(!/^\d+$/.test(num)||num<=0||num>%d){alert('<LABEL_453>');return;};location='%s'.replace(/\{page\}/,document.getElementById('%s').value);}</script><span class="page_input_num"><input onkeyup="ekup()" type="text" id="%s" style="width:40px;vertical-align:text-baseline;padding:0 2px;font-size:10px;border:1px solid gray;"/></span><span class="page_btn_go" onclick="clkyup();" style="cursor:pointer;"><LABEL_1644></span>`, id, pages, url, id, id)
	totalstr := fmt.Sprintf(`<span class="page_total"><LABEL_1834>%d<LABEL_1835></span>`, total)
	pagenation := []string{totalstr, info, prefix, body, subfix, gostr}
	output := []string{}
	for _, v := range order {
		if v-1 < len(pagenation) && v-1 >= 0 {
			output = append(output, pagenation[v-1])
		}
	}
	if pages > 1 {
		return strings.Join(output, "")
	}
	return ""
}

// <LABEL_224>
func (m *misc) GetStrUnicodeIndex(str string, substr string) int {
	// <LABEL_148>
	result := strings.Index(str, substr)
	if result >= 0 {
		return m.GetStrUnicodeIndexByByteIndex(str, result)
	}
	return -1
}

// <LABEL_20>
func (m *misc) GetStrUnicodeIndexByByteIndex(str string, subStrByteIndex int) int {
	if subStrByteIndex > len(str)-1 {
		return -1
	}
	// <LABEL_48>[]byte
	prefix := []byte(str)[0:subStrByteIndex]
	// <LABEL_99>[]rune
	rs := []rune(string(prefix))
	// <LABEL_62>，<LABEL_63>
	result := len(rs)
	return result
}

// <LABEL_49>，<LABEL_1207>
func (m *misc) SubStrUnicode(str string, subStr string, preLen int, sufLen int) string {
	subStrRune := []rune(subStr)
	strRune := []rune(str)
	count := len(strRune)
	subStrUnicodeIndex := m.GetStrUnicodeIndex(str, subStr)
	startIndex := 0
	endIndex := count - 1
	if subStrUnicodeIndex-preLen > 0 {
		startIndex = subStrUnicodeIndex - preLen
	}
	if subStrUnicodeIndex+len(subStrRune)+sufLen < count-1 {
		endIndex = subStrUnicodeIndex + len(subStrRune) + sufLen
	}
	return string(strRune[startIndex:endIndex])
}

// <LABEL_49>，<LABEL_1207>
// subStrIndex <LABEL_299>
func (m *misc) SubStrUnicodeBySubStrIndex(str string, subStr string, subStrIndex int, preLen int, sufLen int) string {
	subStrRune := []rune(subStr)
	strRune := []rune(str)
	count := len(strRune)
	subStrUnicodeIndex := m.GetStrUnicodeIndexByByteIndex(str, subStrIndex)
	startIndex := 0
	endIndex := count - 1
	if subStrUnicodeIndex-preLen > 0 {
		startIndex = subStrUnicodeIndex - preLen
	}
	if subStrUnicodeIndex+len(subStrRune)+sufLen < count-1 {
		endIndex = subStrUnicodeIndex + len(subStrRune) + sufLen
	}
	return string(strRune[startIndex:endIndex])
}
