### 组合模式
允许你将对象组织成树状结构来表现 整体 部分 层次结构，组合让客户能够以一致的方式处理个别对象以及对象组合

### 展现方式
这里的数据和迭代器中的数据相同，都是菜单和菜单项的遍历和展示

meun和meunitem 都继承了 meuncomponentinterface 的所有方法，也就是实现了 meuncomponent的接口,这些方法中有一个是
print,当对这些数据进行遍历的时候，会遍历所有的这些数据