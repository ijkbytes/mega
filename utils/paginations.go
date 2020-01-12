package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
)

type Pagination struct {
	current int   // 当前页码
	count   int   // 记录总数
	perPage int   // 每页条目数
	first   int   // 首页页码
	last    int   // 最后一页页码
	prev    int   // 上一页页码
	next    int   // 下一页页码
	pages   []int // 页码列表
}

func GetPage(c *gin.Context) int {
	ret, _ := strconv.Atoi(c.Query("p"))
	if ret < 1 {
		ret = 1
	}

	return ret
}

func NewPagination(current, perpage, count int) *Pagination {
	p := &Pagination{
		count:   count,
		current: current,
		perPage: perpage,
	}

	if count > 0 {
		p.first = 1
	}

	p.last = int(math.Ceil(float64(p.count) / float64(p.perPage)))

	if p.current < 1 {
		p.prev = 0
	} else {
		p.prev = p.current - 1
	}

	if p.current == p.last {
		p.next = 0
	} else {
		p.next = p.current + 1
	}

	if p.count == 0 || p.last == 0 {
		p.pages = nil
	} else if p.last <= 10 || p.current <= 5 {
		for i := 1; i <= p.last && i <= 10; i++ {
			p.pages = append(p.pages, i)
		}
	} else {
		for i := p.current - 5; i <= p.current+4; i++ {
			p.pages = append(p.pages, i)
		}
	}

	return p
}

func (p *Pagination) PageUrl(url string, page int) string {
	if strings.Contains(url, "?") {
		return fmt.Sprintf("%s&p=%d", url, page)
	}

	return fmt.Sprintf("%s?p=%d", url, page)
}

// 页码列表
func (p *Pagination) Pages() []int {
	return p.pages
}

func (p *Pagination) Next() int {
	return p.next
}

func (p *Pagination) Prev() int {
	return p.prev
}

func (p *Pagination) Last() int {
	return p.last
}

func (p *Pagination) First() int {
	return p.first
}

func (p *Pagination) Current() int {
	return p.current
}
