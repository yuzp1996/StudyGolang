# 装饰者模式

## 作用
一旦熟悉了装饰的技巧，就可以在不修改任何底层代码的情况下，给你的对象赋予新的职责

装饰者模式动态的将责任附加到对象身上，若要扩展功能，装饰者提供了比继承更有弹性的替代方案

## 原则
### 开放-关闭原则
类应该对拓展开放，对修改关闭。 我们的目标是允许类容易扩展，在不修改现有代码的情况下，就可以搭配新的行为

## 如何做
可通过提供扩展的方法来保护代码免于被修改

不可能所有的代码都这么做，将注意力集中在设计中最有可能改变的地方，然后应用开放-关闭原则

* 拿一个咖啡对象
* 用摩卡对象装饰它
* 用奶泡对象装饰它
* 调用cost放放，并依赖委托将调料的价钱加上去