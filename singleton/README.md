# 单例模式

主要有两种，一种是饿汉式，在包init方法中直接初始化；另一种懒汉式则是在使用前调用GetInstance()方法时再初始化。
饿汉式优点是在使用时几乎无消耗，缺点则是在启动时消耗会增加，相当于我不管你用不用，我先帮你准备好了；而懒汉式则恰恰相反，你不用我只会静静地待着，不会强行怼你脸上。

如果不是初始化需要耗时较长的情况下，我个人更倾向与懒汉式。