运行命令：go run main.go ./conf/config.ini

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：应该error直接抛给上层。理由如下： dao层做一些信息的查询时，sql.ErrNoRows只是一种结果形式，不是连接超时等panic错误。dao层把错误类型返回model层，model层对返回类型做判断断言判断，如果是sql.ErrNoRows，wrap一下信息打印日志到终端，返回业务层的数据也是应该正常的返回结果 （里面的数据可能不正常，看日志排查问题）
