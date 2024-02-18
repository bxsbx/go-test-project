package main

import (
	"fmt"
	"github.com/schollz/progressbar"
	"github.com/schollz/progressbar/themes"
	"math/rand"
	"strings"
	"time"
)

func main() {
	bar := progressbar.New(100)

	//theme, _ := themes.NewDefault(1)
	theme := themes.New([]string{"*", " ", "{", "}"}...)
	bar.SetTheme(theme)

	//bar.Reset()
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(20 * time.Millisecond)
	}

	//MyBar(1000)
}

type ProgressBar struct {
	Sf            string  //进度符号
	CurTotal      int     //当前进度数
	Total         int     //总进度数
	Rate          float64 //当前进度百分比例
	SpendTime     int64   //已花费时间，单位为毫秒
	RemainderTime int64   //剩余时间，单位为毫秒
	SpendRate     int64   //平均速率，单位毫秒/度
}

func NewBar(sf string, total int) *ProgressBar {
	return &ProgressBar{
		Sf:    sf,
		Total: total,
	}
}

func (p *ProgressBar) Add(n int, spendTime int64) {
	p.CurTotal += n
	p.Rate = float64(p.CurTotal) / float64(p.Total) * 100
	p.SpendTime += spendTime
	p.SpendRate = p.SpendTime / int64(p.CurTotal)
	p.RemainderTime = int64(p.Total-p.CurTotal) * p.SpendRate
}

// 毫秒转时分秒
func (p *ProgressBar) SToHmS(ms int64) string {
	str := ""
	s := ms / 1000
	if ms%1000 > 0 {
		s++
	}
	m := s / 60
	if m > 0 {
		h := m / 60
		if h > 0 {
			str += fmt.Sprintf("%d时", h)
		}
		str += fmt.Sprintf("%d分", m%60)
	}
	str += fmt.Sprintf("%d秒", s%60)
	return str
}

func (p *ProgressBar) PrintProgressBar() {
	n := int(p.Rate)
	fs := strings.Repeat(p.Sf, n) + strings.Repeat(" ", 100-n)
	fmt.Printf("\r进度条：[%s]，当前进度：%d/%d，平均速率：%v，剩余时间：%s   ", fs, p.CurTotal, p.Total, p.SpendRate, p.SToHmS(p.RemainderTime))
}

func MyBar(total int) {
	progressBar := NewBar("█", total)
	for i := 0; i < total; i++ {
		randInt := int64(rand.Intn(30) + 20)
		time.Sleep(time.Duration(randInt) * time.Millisecond)
		progressBar.Add(1, randInt)
		progressBar.PrintProgressBar()
	}
}
