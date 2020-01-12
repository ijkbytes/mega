package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
)

type Pagination struct {
	current int // 当前页码
	count   int // 记录总数
	perPage int // 每页条目数
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
	last := int(math.Ceil(float64(p.count) / float64(p.perPage)))

	var ret []int

	if p.count == 0 || last == 0 {
		return ret
	}

	if last <= 10 || p.current <= 5 {
		for i := 1; i <= last && i <= 10; i++ {
			ret = append(ret, i)
		}
		return ret
	}

	for i := p.current - 5; i <= p.current+4; i++ {
		ret = append(ret, i)
	}
	return ret
}

func (p *Pagination) Next() int {
	last := int(math.Ceil(float64(p.count) / float64(p.perPage)))

	if p.current == last {
		return 0
	}
	return p.current + 1
}

func (p *Pagination) Prev() int {
	if p.current < 1 {
		return 0
	}
	return p.current - 1
}

func (p *Pagination) Last() int {
	return int(math.Ceil(float64(p.count) / float64(p.perPage)))
}

func (p *Pagination) First() int {
	return 1
}

func (p *Pagination) Current() int {
	return p.current
}
