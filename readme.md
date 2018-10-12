# KunhnMunkress
带权二分图最大匹配Kuhn-Munkres算法go语言实现。是将 http://csclab.murraystate.edu/~bob.pilgrim/445/munkres.html 中的c#实现改为go实现，同时参考了https://github.com/zarrabeitia/munkres_sparse

## Installation

```
go get -u github.com/dazheng/kuhnmunkres
```
## 注意
* 是按最小值求的最大匹配，如果是最大值的请转成最小值
* 传入的二维数组每个都必须有值，没有的填整型最大值

## 参考资料
* 此算法的详细解释见：http://csclab.murraystate.edu/~bob.pilgrim/445/munkres.html
* 匈牙利算法(Kuhn-Munkres)算法： https://blog.csdn.net/zsfcg/article/details/20738027