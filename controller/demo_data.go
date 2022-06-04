package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "https://v3-web.douyinvod.com/f6d503547a149f176a3f5f25213e3609/629b1f91/video/tos/cn/tos-cn-ve-15c001-alinc2/ea307f6cf9ee48f2bccae1d3f34d062c/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=2722&bt=2722&cs=0&ds=3&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcwIOVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=NzZmZjgzZTRpNzs1ZWg6aUBpM2xvcWQ6Zjs1ZDMzNGkzM0BeL15eMTReXzYxMzEuMjAuYSNycDNpcjRfaTJgLS1kLTBzcw%3D%3D&l=20220604160157010151211086431C8F97",
		CoverUrl:      "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/94c75520b32341a9a950bf90eccee648_1653222266~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=%2B95ZBFgrr1iwPIne8JMOlhfxhX8%3D&from=4257465056&se=false&biz_tag=feed_cover&l=20220604160157010151211086431C8F97",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            3,
		Author:        DemoUser,
		PlayUrl:       "https://v26-web.douyinvod.com/d95a355041f77983f7a18a12ae1aa2c7/629b1fde/video/tos/cn/tos-cn-ve-15c001-alinc2/1653a991409d4fb097542e1132c09db7/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1294&bt=1294&cs=0&ds=3&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcwIOVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=MzxoZjQ5ZWk5Omc2Nmc6Z0BpajplcTY6Zjs0ZDMzNGkzM0BfNV80NS81X14xYzRjLl5eYSNwaW5zcjQwa19gLS1kLWFzcw%3D%3D&l=20220604160157010151211086431C8F97",
		CoverUrl:      "https://p6-pc-sign.douyinpic.com/tos-cn-p-0015/8f3906b08e884842b6c571b50a54c942_1654004716~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=Ma4YKsyM5UMuN%2BZ0skLXEfYzYEI%3D&from=4257465056&se=false&biz_tag=feed_cover&l=20220604160157010151211086431C8F97",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            4,
		Author:        DemoUser,
		PlayUrl:       "https://v26-web.douyinvod.com/19543971ae0d3dca1a54eb4c59faf774/629b211a/video/tos/cn/tos-cn-ve-15c001-alinc2/fa6d9865dbd143409496b25a7f7c209b/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=3120&bt=3120&cs=0&ds=6&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcW4OVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=Nmc2ZzY2ZzlpZTY5ZGc8ZUBpM3M0bTk6ZnE2ZDMzNGkzM0BgYDUtXmFiNTIxMzBgX2JhYSNqL3NlcjRvcjZgLS1kLWFzcw%3D%3D&l=2022060416082801014020714310F26E9B",
		CoverUrl:      "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/6c01c63be37a4a398745723005726f3f_1653751697~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=bsM93JfLkgw%2FkFrZIpICpWNIYwM%3D&from=4257465056&se=false&biz_tag=feed_cover&l=2022060416082801014020714310F26E9B",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            5,
		Author:        DemoUser,
		PlayUrl:       "https://v3-web.douyinvod.com/10a60bd2785721d695b6e0bd4c472368/629b21ef/video/tos/cn/tos-cn-ve-15c001-alinc2/79ca6bc9b35c42998092f414faf7b7ac/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1239&bt=1239&cs=0&ds=3&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcW4OVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=N2g2ODk8ZDhoOWhoOmg2ZkBpajlnZDo6Zm9wZDMzNGkzM0AuLl5eLTExX2ExL2M2LjY2YSNpYDRncjRnbjZgLS1kLS9zcw%3D%3D&l=2022060416082801014020714310F26E9B",
		CoverUrl:      "https://p6-pc-sign.douyinpic.com/tos-cn-p-0015/151299fcbe3542629a34af1b3f8df92f_1653828985~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=CJQ%2F5EhLCTHNuqLs3RA6BQttX%2FQ%3D&from=4257465056&se=false&biz_tag=feed_cover&l=2022060416082801014020714310F26E9B",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            6,
		Author:        DemoUser,
		PlayUrl:       "https://v26-web.douyinvod.com/78c3e6711123e8bb227b4fdaadb2501d/629b2121/video/tos/cn/tos-cn-ve-15c001-alinc2/9849038e197047f9bfe8be8c74f9ce18/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1342&bt=1342&cs=0&ds=6&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcW4OVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=ZjxnOTZkZ2c4aDc2N2g0NUBpanl3Zzk6ZjM4ZDMzNGkzM0AtNjQxLjZjX18xNWIxNWFiYSMvLi5vcjRva2BgLS1kLS9zcw%3D%3D&l=2022060416082801014020714310F26E9B",
		CoverUrl:      "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/e5f51167e9b046ae9ba5a87112d9d9a4_1654152168~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=tvcPnq66SN4j5HEwY2wYfSpb%2BE8%3D&from=4257465056&se=false&biz_tag=feed_cover&l=2022060416082801014020714310F26E9B",
		FavoriteCount: 12,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            7,
		Author:        DemoUser,
		PlayUrl:       "https://v26-web.douyinvod.com/82258c9495e57d482ae542b42e42a3eb/629b213e/video/tos/cn/tos-cn-ve-15c001-alinc2/27f244658a74465892eced88c865a3a1/?a=6383&ch=5&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=191&bt=191&cs=0&ds=2&ft=X1nbLXvvBQOBULry_8Z.PRDFVKlcW4OVF2bLqWICfiZm&mime_type=video_mp4&qs=0&rc=PGQ0Ozo0ZmlpN2U2Zmg2ZUBpajU3amY6ZmhyZDMzNGkzM0BjYzU0LV8vXjIxMGJhYzJgYSNiNi9ecjRfLjFgLS1kLS9zcw%3D%3D&l=2022060416082801014020714310F26E9B",
		CoverUrl:      "https://p3-pc-sign.douyinpic.com/tos-cn-p-0015/4f8686d14ab74c82a7f4ebb5b7ffb467_1653178612~tplv-dy-360p.jpeg?x-expires=1655539200&x-signature=b%2FjLR3ww9HQ0WXheoI59jvHD2G4%3D&from=4257465056&se=false&biz_tag=feed_cover&l=2022060416082801014020714310F26E9B",
		FavoriteCount: 1,
		CommentCount:  2,
		IsFavorite:    true,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
