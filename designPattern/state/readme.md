### 状态模式
糖果机的事情，每个状态对应不同的状态类，而不是用数字来代替。我们为这些状态类实现不同的方法，让不同状态在不同操作下有不同的动作

### 定义
状态模式允许对象在内部状态发生变化的时候改变它的行为，对象看起来好像修改了它的类

### 实现
我们将糖果机不同的状态实例化，每个状态都有所有的转换方法，我们转换这个糖果机状态的时候，对应的转换方法也会发生改变。好像这个糖果机变成了其他类的东西