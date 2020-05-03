package Iterator

import (
	"strconv"
)

type MeunIterator interface {
	HasValue() bool
	GetValue() MeunItem
}

type Iterator interface {
	CreateIterator()MeunIterator
}

type MeunItem struct {
	description string
	name string
}
type BreakfastMeun struct {
	Items []MeunItem
}

func NewBreakfastMeun()*BreakfastMeun{
	return &BreakfastMeun{
		Items: []MeunItem{
			{
				description: "好吃的果子",
				name:        "果子",
			},
			{
				description:"好吃的豆浆",
				name:"豆浆",
			},
		},
	}
}
func (meun BreakfastMeun)CreateIterator()MeunIterator  {
	return &BreakfastMeuniterator{index:0,Items:meun.Items}

}

type BreakfastMeuniterator struct {
	index int
	Items []MeunItem
}



func(meun *BreakfastMeuniterator)HasValue()bool{
	if meun.index < len(meun.Items){
		return true
	}
	return false
}

func (meun *BreakfastMeuniterator)GetValue()MeunItem{
	result := meun.Items[meun.index]
	meun.index += 1
	return result
}




type LunchMeun struct {
	Items map[string]MeunItem
}

func NewLunchMeun()*LunchMeun{
	return &LunchMeun{
		Items: map[string]MeunItem{"0":MeunItem{
			description: "好吃的鱼香肉丝",
			name:        "鱼香肉丝",
		},"1":MeunItem{
			description: "好吃的凉皮",
			name:        "凉皮",
		}},
	}
}

func (meun LunchMeun)CreateIterator()MeunIterator  {
	return &LunchMeunIterator{Items:meun.Items,index:0}

}

type LunchMeunIterator struct {
	index int
	Items map[string]MeunItem
}

func(meun *LunchMeunIterator)HasValue()bool{
	if meun.index < len(meun.Items){
		return true
	}
	return false
}

func (meun *LunchMeunIterator)GetValue()MeunItem{
	result := meun.Items[strconv.Itoa(meun.index)]
	meun.index++
	return result
}


