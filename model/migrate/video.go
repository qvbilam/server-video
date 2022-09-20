package main

import (
	"gorm.io/gorm"
	"video/model"
)

func CreateVideo(db *gorm.DB) {
	db.Exec("truncate table category")
	CartoonModels := getCartoon()
	MovieModels := getMovie()
	TVModels := getTV()
	db.CreateInBatches(CartoonModels, 100)
	db.CreateInBatches(MovieModels, 100)
	db.CreateInBatches(TVModels, 100)
}

func getCartoon() []model.Video {
	models := []model.Video{
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "某科学的超电磁炮",
			Introduction:   "某科学的电磁炮",
			Icon:           "http://damowang.test.upcdn.net/video/images/we5310pi2yur2to539q2.gif",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "fate zero 第一季",
			Introduction:   "传说中，圣杯是能够实现拥有者愿望的宝物。为了追求圣杯的力量，7位魔术师各自召唤英灵，展开争夺圣杯的战斗，这就是圣杯战争。\n时间退回到第五次圣杯战争的10年前，第四次圣杯战争，参与者正是士郎他们的父辈。为了得到圣杯不择手段的卫宫切嗣，年轻时代的言峰绮礼，间桐家与远坂家的关系，同样身为王却意志不同的三位英灵。第四次圣杯之战就此爆发。",
			Icon:           "http://damowang.test.upcdn.net/video/dm/FZ/f101.png",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "fate stay night",
			Introduction:   "圣杯是传说中可实现持有者一切愿望的宝物，而为了得到圣杯的仪式就被称为圣杯战争。参加圣杯战争的7名由圣杯选出的魔术师被称为Master，与7名被称为Servant的使魔订定契约。他们是由圣杯选择的七位英灵，被分为七个职阶，以使魔的身份被召唤出来。能获得圣杯的只有一组，这7组人马各自为了成为最后的那一组而互相残杀。",
			Icon:           "http://damowang.test.upcdn.net/video/images/12wu80p2y83ir18qote5.jpeg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "进击的巨人 第一季",
			Introduction:   "进击的巨人是根据日本漫画家谏山创的漫画作品所改编的同名动画作品，主要讲述高墙内的人类为了生存而对抗巨人的故事。\n动画讲述了107年前（743）年，世界上突然出现了人类的天敌“巨人”，面临着生存危机而残存下来的人类逃到了一个地方，盖起了三重巨大的城墙。人们以自由为代价，在这隔绝的环境里享受了一百多年的和平，直到艾伦·耶格尔十岁那年，50米高的“超大型巨人”突然出现，以压倒性的力量破坏城门，其后瞬间消失，巨人们成群的冲进墙内捕食人类。\n艾伦亲眼看着人们以及自己的母亲被巨人吞食，怀着对巨人无法形容的憎恨，誓言杀死全部巨人。城墙崩坏的两年后，艾伦加入104期训练兵团学习和巨人战斗的技术。在训练兵团的三年，艾伦在同期训练兵里有着其他人无法比拟的强悍精神力，即使亲眼见过地狱也依然勇敢地向巨人挑战的艾伦，如愿以偿加入了向往已久的调查兵团。",
			Icon:           "http://damowang.test.upcdn.net/video/dm/进击的巨人/jin101.png",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "进击的巨人 第二季",
			Introduction:   "进击的巨人是根据日本漫画家谏山创的漫画作品所改编的同名动画作品，主要讲述高墙内的人类为了生存而对抗巨人的故事。\n动画讲述了107年前（743）年，世界上突然出现了人类的天敌“巨人”，面临着生存危机而残存下来的人类逃到了一个地方，盖起了三重巨大的城墙。人们以自由为代价，在这隔绝的环境里享受了一百多年的和平，直到艾伦·耶格尔十岁那年，50米高的“超大型巨人”突然出现，以压倒性的力量破坏城门，其后瞬间消失，巨人们成群的冲进墙内捕食人类。\n艾伦亲眼看着人们以及自己的母亲被巨人吞食，怀着对巨人无法形容的憎恨，誓言杀死全部巨人。城墙崩坏的两年后，艾伦加入104期训练兵团学习和巨人战斗的技术。在训练兵团的三年，艾伦在同期训练兵里有着其他人无法比拟的强悍精神力，即使亲眼见过地狱也依然勇敢地向巨人挑战的艾伦，如愿以偿加入了向往已久的调查兵团。",
			Icon:           "http://damowang.test.upcdn.net/video/dm/进击的巨人第二季/b21bb051f81986188bb69ec341ed2e738ad4e680.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "玉子超市",
			Introduction:   "玉子超市（たまこまーけっと）是由京都动画制作的原创动画。内容则是座落於某个小镇的兔子山商店街上，有一家日式饼店。店内住著一位十分喜欢饼类小吃的高中一年级女生——玉子，她所发生的日常生活趣事。",
			Icon:           "http://damowang.test.upcdn.net/video/dm/玉子超市/1.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "狼之子",
			Introduction:   "狼之子",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "你的名字",
			Introduction:   "故事发生的地点是在每千年回归一次的彗星造访过一个月之前，日本飞驒市的乡下小镇糸守町。在这里女高中生三叶每天都过着忧郁的生活，而她烦恼的不光有担任镇长的父亲所举行的选举运动，还有家传神社的古老习俗。在这个小小的城镇，周围都只是些爱瞎操心的老人。为此三叶对于大都市充满了憧憬。然而某一天，自己做了一个变成男孩子的梦。这里有着陌生的房间、陌生的朋友。而眼前出现的则是东京的街道。三叶虽然感到困惑，但是能够来到朝思暮想的都市生活，让她觉得神清气爽。另一方面在东京生活的男高中生立花泷也做了个奇怪的梦，他在一个从未去过的深山小镇中，变成了女高中生。两人就这样在梦中邂逅了彼此",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "机器人总动员",
			Introduction:   "公元2700年，地球早就被人类祸害成了一个巨大的垃圾场，已经到了 无法居住的地步，人类只能大举迁移到别的星球，然后委托一家机器人垃圾清理公司善后，直至地球的环境系统重新达到生态平衡。在人类离开之后，垃圾清理公司将机器人瓦力成批地输送到地球，并给他们安装了惟一的指令——垃圾分装，然而随着时间的推移，机器人一个接一个地坏掉，最后只剩下惟一的一个，继续在这个似乎已经被遗忘了的角落，勤勤恳恳地在垃圾堆中忙碌着，转眼就过去了几百年的时间，寂寞与孤独变成了围绕着他的永恒的主题。",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "赏金猎人：天国之扉",
			Introduction:   "这部被称为“二十世纪最后五年里最优秀的动画作品”出现在大银幕上确实已是一件让CB迷们兴奋的事情。这一次，SPIKE他们猎取的目标是一个叫VINCENT的恐怖分子。并且还和一个叫ELECTRA的女人牵扯着千丝万缕的关系……担任音乐还是菅野洋子，这位音乐奇才将爵士、朋克、重金属、乡村、民谣等各种类型音乐元素等融入到片中，使电影更加具有强大的感染力。",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "银魂:永远的万事屋",
			Introduction:   "德川幕府末期，来自外太空的天人降临，江户完全处于他们的统治之下。在此之后，天人文明高度发达，高楼大厦拔地而起，电车电话遍布大街小巷。在这座现代化都市的一隅，平素里懒洋洋但骨子里仍秉持着武士之魂的坂田银时经营着一家接受客人各种委托的万事屋。在存在感稀薄的志村新八和夜兔族幸存者神乐围绕下，万事屋度过繁忙而又麻烦多多的每一天。",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     1,
			Name:           "紫罗兰的永恒花园",
			Introduction:   "自动手记人偶这个名字名噪一时的时候，已经是相当一段时间前的事了。奥兰德博士发明了用于书写记录人声话语的机械。当初他仅仅为了爱妻而制造的这一机械，不知何时已普及至全世界，甚至设立了提供借出服务的机关。\n若是客人您有所期望的话，我可以去往任何地方。我是，自动手记人偶服务人员，薇尔莉特·伊芙加登。\n有着如同从故事书中出现般的样貌，金发碧眼的女子，保持着无机质的美，以玲珑剔透的声音如此说到",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}
	return models
}

func getMovie() []model.Video {
	models := []model.Video{
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "黄飞鸿之男儿当自强",
			Introduction:   "本片讲述了黄飞鸿应国际医学会的邀请，与梁宽、十三姨前赴广州，途中与白莲教结下纠纷。后得革命义士陆皓东之助，寄居英国领事馆，并结识革命领袖孙文。清朝大臣元述为对付革命义士，煽动白莲教团攻打英国领事馆，黄飞鸿为保存国运命脉，与元述对抗，并帮助孙文等人逃脱。",
			Icon:           "http://damowang.test.upcdn.net/video/images/etw17p55ry472oq9i62u.gif",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "闪光少女",
			Introduction:   "是由安乐影片有限公司出品发行的青春热血校园电影，由王冉执导，徐璐、彭昱畅领衔主演，刘泳希、韩忠羽、李诺、陈雨锶、乐思宏等主演的青春热血校园电影。该片讲述了神经少女陈惊与男闺蜜和小伙伴组成2.5次元乐队，大战西洋乐的励志青春故事。该片于2017年7月20日在中国上映。",
			Icon:           "http://damowang.test.upcdn.net/video/images/rui7pe1to2251q1y62w2.jpeg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "阿甘正传",
			Introduction:   "《阿甘正传》是由罗伯特·泽米吉斯执导的电影，由汤姆·汉克斯、罗宾·怀特等人主演，于1994年7月6日在美国上映。电影改编自美国作家温斯顿·格卢姆于1986年出版的同名小说，描绘了先天智障的小镇男孩福瑞斯特·甘自强不息，最终“傻人有傻福”地得到上天眷顾，在多个领域创造奇迹的励志故事。电影上映后，于1995年获得奥斯卡最佳影片奖、最佳男主角奖、最佳导演奖等6项大奖。2014年9月5日，在该片上映20周年之际，《阿甘正传》IMAX版本开始在全美上映",
			Icon:           "http://damowang.test.upcdn.net/video/dy/阿甘正传/agzz.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "p与jk",
			Introduction:   "P是警察Police的首字母，JK是日语女高中生的缩写。高一女生歌子（土屋太凤饰）在好友的邀请下以“22岁女子”的身份参加成年人的联谊，与年轻的警察功太（龟梨和也饰）相遇并且互相产生爱慕之情。然而，在得知歌子还只是高中生后，功太压抑自己的感情，选择对她冷漠处理。不料，歌子因包庇他而受伤，以这个件为契机，功太做出决定，作为为了和歌子在一起的手段，向歌子提出了“不是作为恋人交往，而是结婚”的提议",
			Icon:           "http://damowang.test.upcdn.net/video/dy/阿甘正传/agzz.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "一呼一吸",
			Introduction:   "年仅28岁的罗宾（安德鲁·加菲尔德饰）患上了小儿麻痹症。瞬间，这位生龙活虎、充满运动细胞的年轻人就被击倒了，他脖子以下全部瘫痪，不用机器就无法呼吸，甚至有段时间都没法说话。未来太惨淡，罗宾甚至哀求自己的妻子黛安娜（克莱尔·福伊饰）让他结束生命。此时的黛安娜已经怀孕了，她拒绝向罗宾的负能量屈服。尽管医生诊断的结果是罗宾只能再活3个月，罗宾却不信命，和妻子一起做出了出院的决定。",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/images/taobisuirankechi.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "关原之战",
			Introduction:   "羽柴秀吉即丰臣秀吉解决织田家内部问题后，控制了织田氏旧家臣的实权，有意统一日本。他想利用武力屈服德川家康，可是在小牧·长久手之战中挫败，转而采取与德川家康达成议和的外交屈服手段，先有石川数正突然投靠秀吉，再利用婚姻使家康臣服。秀吉在1587年就任关白一职",
			Icon:           "http://damowang.test.upcdn.net/video/images/3e2t51w4qpi19yo72u5r.png",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "功夫",
			Introduction:   "20世纪40年代的中国广东，有一名无可药救的小混混阿星（周星驰饰），此人能言善道、擅耍嘴皮，但意志不坚，一事无成。他一心渴望加入手段冷酷无情、恶名昭彰的斧头帮，并梦想成为黑道响叮当的人物。此时斧头帮正倾全帮之力急欲铲平唯一未收入势力范围的地头，未料该地卧虎藏龙，痴肥的恶霸女房东肥婆加上与其成对比的懦弱丈夫二叔公，率领一班深藏不漏的武林高手，大展奇功异能，对抗恶势力",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "契克",
			Introduction:   "该片改编自德国作家沃尔夫冈·赫恩多夫的同名小说《我的好友契克》，讲述了住在柏林郊区的迈克，与东方脸孔的转校生契克一见如故，两人于是展开了一段夏日冒险旅程",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "血战钢锯岭",
			Introduction:   "在1942年的太平洋战场，军医戴斯蒙德·道斯（安德鲁·加菲尔德饰）不愿意在前线举枪射杀任何一个人，他因自己的和平理想遭受着其他战士们的排挤。尽管如此，他仍坚守信仰及原则，孤身上阵，无惧枪林弹雨和凶残日军，誓死拯救即使一息尚存的战友。数以百计的同胞在敌人的土地上伤亡惨重，他一人冲入枪林弹雨，不停地祈祷，乞求以自己的绵薄之力尽再救一人，75名受伤战友最终被奇迹般的运送至安全之地，得以生还。",
			Icon:           "",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "建军大业",
			Introduction:   "1927年,北伐战争刚取得重大成果之际,国民党“右派”为夺权电影《建军大业》电影《建军大业》叛变革命,发动了疯狂的“清共”行动,短短数月,近31万进步同胞遭到残酷杀害,全国震惊,刚刚看到希望的中国即将再次陷入军阀混战和独裁专制的深渊。由于没有自己的武装力量,成立不足7年的中国共产党在国民党“右派”的疯狂进攻下,几乎遭遇毁灭性打击。血的教训使毛泽东、周恩来等党内进步分子认识到了“枪杆子里出政权”的硬道理。生死存亡之际,他们临危受命,冒着生命危险分赴湖南和南昌等地,联合朱德、贺龙、叶挺、刘伯承等一批爱国发动起义,誓要组建一支真正属于人民的军队。铁血铸军魂,舍己保家国。",
			Icon:           "http://damowang.test.upcdn.net/video/dy/建军大业/jianjun.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     3,
			Name:           "银魂真人版",
			Introduction:   "江户时代末期，被称为“天人”的异星人来袭。地球人与天人之间的战争瞬即爆发，为数众多的武士和攘夷派志士都参与与天人的战斗。幕府见识到天人强大的实力后，擅自与天人签订不平等条约，准许他们入国。其后更颁布了“废刀令”，夺走了武士的刀。自此，天人横行霸道，幕府成为了“傀儡政权”。在这个时代，武士坂田银时（小栗旬饰）与伙伴们过着异想天开的生活。",
			Icon:           "http://damowang.test.upcdn.net/video/dy/银魂真人版/yhzr.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}
	return models
}

func getTV() []model.Video {
	models := []model.Video{
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     4,
			Name:           "逃避虽然可耻但有用",
			Introduction:   "逃避虽然可耻但有用",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/images/taobisuirankechi.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     4,
			Name:           "骑士陨落",
			Introduction:   "由DonHandfield及Richard Rayner主创的10集历史剧集《骑士陨落 Knightfall》讲述的是十字军中，圣殿骑士团的传奇故事。剧情描绘了骑士团在战场上的信念，忠诚，还有兄弟情，也会谈及为什么13号的黑色星期五会成为世界上最忌讳的日子。另外Jeremy Renner亦是该剧执行制片，曾有传他会客串演出（但未有肯定）。 \n　　Tom Cullen饰演Landry，是十字军东征的老将，圣殿骑士团的领导，他一直在探索基督教的最珍贵的宝物：圣杯。身手厉害﹑勇猛的他对上帝及自己的信心因在圣地战败而深深动摇，不过在圣杯失去多年踪影但突然再次传出它存在的消息后，令到他再次燃起斗志，并决心带领圣殿骑士团进入神秘﹑信仰﹑流血的冒险。 ",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/images/taobisuirankechi.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     4,
			Name:           "大秦帝国之崛起",
			Introduction:   "秦昭襄王王当政不久,列国展开一系列兼并战争。秦国看准时机,先与齐连横,攻略魏国大片土地。由此,战国到了秦齐赵三国并强阶段。此时秦国国内,穰侯魏冉擅权,宣太后掌政,秦昭襄王王权旁落。此时,局势对秦国十分有利：纵横家苏秦与燕王秘约赴齐间齐,以灭齐为目标“死间”齐国。面对国内政治困局和国外有利机遇,秦昭王毅然起用范雎,逐魏冉、废太后之权而恭养之。同时,与苏秦合谋,举兵攻破齐国。秦昭王遣将白起,相继在伊阙、郢都、华阳、长平发动四大战役,歼灭了韩魏齐楚赵百万以上兵力,使秦与列国消长发生了根本转折。自此,秦国崛起.",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/大秦帝国之崛起/daqin.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     4,
			Name:           "隋唐演义",
			Introduction:   "该剧讲述北隋唐演义隋唐演义齐名将秦彝之后秦琼(严屹宽饰)自小与程咬金一起长大,凭借一身武功被官府提拔为下级小官,某日秦琼在临潼山遇到李渊被金蛇卫截杀,单枪匹马救出了李渊全家。而后来到潞州由于受伤不得不当锏卖马来维持生计,幸被单雄信得知加以接济。程咬金在长叶林劫夺了杨林四十八万两白银皇纲。秦母大寿当天程咬金被抓,为搭救程咬金,贾柳店结拜的众兄弟策划造反劫狱,机缘巧合程咬金做了瓦岗寨的首领混世魔王。隋炀帝(富大龙饰)去扬州看琼花,瓦岗寨联合其他起义军汇聚在四明山意图杀死隋炀帝,推翻大隋朝。隋炀帝为灭反王摆下铜旗阵。李世民因隋炀帝下旨调兵十万前去守阵,被逼起兵造反。李密逃到瓦岗被众人拥戴为主并称魏王。隋炀帝被杀后。魏王用玉玺换来萧美娘,朝政日渐腐败,瓦岗山上众将作鸟兽散。秦琼遂率程咬金、罗成投于李世民麾下,屡立战功,最终助李唐王朝平定天下。",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/隋唐演义/suitang.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: 1,
			},
			CategoryId:     4,
			Name:           "谈判官",
			Introduction:   "童薇凭借扎实的专业功底和胆大心细的谈判风格在商务谈判桌上无往而不利是中美经贸协会中最年轻耀眼的谈判专家。一次商业并购项目让她结识了谢晓飞。谢家是美国社会隐秘而富裕的华人世家谢晓飞是家族的继承人。从小在美国长大的谢晓飞心高气傲因理想得不到父亲和家族的支持所以在谈判中频频和童薇作对。但随着交往的深入两人逐渐从互不认同到彼此欣赏最终情定谈判桌。此时谢父遭遇背叛谢晓飞尝遍人情冷暖。童薇挺身而出用自己的谈判专长替谢家拿下重要项目谢家最终收回股份。此时童薇发现当年父母的离世竟与谢家有关她和谢晓飞因上代人的恩怨最终分道扬镳。多年后重逢两个曾经相爱的人却要在谈判桌上兵戎相见。童薇和谢晓飞在相爱相虐中确定了彼此的爱谢晓飞也感慨中国之繁盛说服家族将业务重心移回中国",
			Icon:           "http://damowang.test.upcdn.net/video/dsj/谈判官/tanpanguan.jpg",
			HorizontalIcon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}
	return models
}
