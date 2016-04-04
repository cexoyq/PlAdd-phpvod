# PlAdd-phpvod
遍历目录下的影片，然后添加到phpvod的数据库

1.pv_video表里：
	cid表示影片的分类ID，必须要填，否则不会在WEB上显示电影
	subject表示影片的名称

2.pv_urls表里：
	rul表示影片的播放地址
	vid是影片在pv_video表里的vid
	pid与其它的一致，用5

cid:
35	日语
34	俄语
33	韩语
32	德语
31	法语
30	国语
43	英语


http://c.300y.com.cn/vod/xfplay.exe 
"INSERT pv_video SET cid=43,nid=1,author='xiongyi',authorid=1,subject=?"
"INSERT pv_urls SET vid=?,pid=5,url=?"
