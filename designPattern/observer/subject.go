package observer

type Subject interface {
	registryobserver(observer)
	deleteobserver(observer)
	notificateobserver()
	setrealtimeData(data string)
}

//天气发布者
type WeaterDate struct {
	Date string
	Observers []observer
}

func NewWeaterDate()WeaterDate{
	return WeaterDate{}
}

func(wd *WeaterDate)registryobserver(ob observer){
	wd.Observers = append(wd.Observers,ob)
}

func(wd *WeaterDate)deleteobserver(ob observer){
	for index,observer := range wd.Observers{
		if ob.getname() == observer.getname(){
			wd.Observers = append(wd.Observers[0:index],wd.Observers[index+1:]...)
		}
	}

}

func (wd *WeaterDate)setrealtimeData(data string){
	wd.Date = data
	wd.notificateobserver()
}

func(wd *WeaterDate)notificateobserver(){
	for _,observer:= range wd.Observers{
		observer.update(wd.Date)
	}
}

