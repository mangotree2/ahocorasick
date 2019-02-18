### AC自动机过滤器

----

```
func prepareData() string {
	dict := make(map[int]string, 1115)
	f, err := os.OpenFile("dict.txt", os.O_RDONLY, 0660)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	i := 0
	for {
		l, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		piece := bytes.Split(bytes.TrimSpace(l), []byte("\t"))
		key := string(piece[0])
		dict[i] = key
		i++
	}

	str:= ""
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0 ;i<100;i++{
		r := rnd.Intn(len(dict)*2)
		if r >= len(dict) {
			str += "#"
		} else {
			str += dict[r]
		}


	}
	return str
}

func BenchmarkAC_Filter(b *testing.B) {

	b.StopTimer()
	input := prepareData()
	b.Log("len:",len(input),"input:",input)
	b.StartTimer()

	ac := FromFile("dict.txt")

	//input := "lkfjglkfj我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够lddf我的忑是，测试，让字数够lgkjfoidjgoifjdgljflkjgfdji我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够lsogjfdkljglkjflk jglkfjl kgjfdlijglk dfjkgljfdljg lkfjdklgjflkjglk 衣服价格i都放假哦i攻击对方结果发夫尼克流年不利百分比法你看老夫的内控规范了的感觉舒服就看过了就发生了快过节了富士康辜负了时间过来看是否结果来看是飞机离开关键时刻发了几个离开房间观看了就反过来看手机分隔符" +
	//	"一辈子的孤单是得呀，候鸟de 倒霉命运；特工小子就是你"

	output := ac.Filter(input)

	b.Log("len:",len(output),"out:",output)
}

```
```

goos: windows
goarch: amd64
pkg: github.com/mangotree2/ahocorasick
BenchmarkAC_Filter-4   	2000000000	         0.00 ns/op
--- BENCH: BenchmarkAC_Filter-4
    ac_test.go:110: len: 703 input: ###新海神号##英烈的岁月康纳斯的战争丑女贝蒂####人生几度秋凉我不是黄蓉一言为定深呼吸#绿茶##斗鱼#你怎么舍得我难过#蠢蛋进化论一米阳光#买办之家###十分爱#他还是不懂###那一夜命中注定#狼吞虎咽十分爱#宝莱坞生死恋#一千年以后总统千金欧游记##沧海一声笑#桃色交易男儿当自强后窗惊魂##间谍游戏###死心的理由辛瑞那#终极土匪兵人#飞一般爱情小说#世界自然奇观如果有一天###沧海一声笑##与魔鬼同行云淡天高#三叶草绅士大盗#机械公敌#我唾弃你的坟墓####开往春天的地铁恋爱高手#云水谣##一辈子的孤单疯狂木乃伊魔鬼末日#
    ac_test.go:120: len: 100 out: ###*##***####****#*##*#*#**#*###*#*###**#**#*#**##*#***##*###**#**#*#**###*##**#**#*#*####**#*##***#
    ac_test.go:110: len: 621 input: #不能说的秘密####天网追击空战英豪####不良主妇##偏不离婚#穿牛仔裤的十字军#东成西就U-571关中匪事轨迹#####女神陷阱#保护证人组刺青##切肤只爱####野兽之瞳#三男三女#战略阴谋大雨成灾##该隐的记号阿宝的故事#一手托两家无法阻挡的婚姻放学后的屋顶##双面冷血杀手###留级之王芥末不辣杰西警探犯罪档案###宣言###第十大道和狼#虫洞效应##忘不了#珍妮朱诺斗鱼#王者无敌伤心太平洋#真实银行抢劫案欢迎来到东莫村#哆啦A梦最后死战京城四少###谍影重重值得###亲密有罪###
    ac_test.go:120: len: 100 out: #*####**####*##*#*#****#####*#**##*####*#*#**##**#***##*###***###*###*#*##*#**#**#**#***###**###*###
    ac_test.go:110: len: 630 input: #百万富翁的初恋#######噢!北鼻##红鞋淑女手册我和春天有个约会#我爱天上人间顶尖对决##新金瓶梅##末日浩劫失恋阵线联盟####爱的躯壳毕业生青春爱人事件轰天潜舰#毁灭战士人鱼小姐#哭泣的内衣###黑衣人孟菲斯美女号#我们的存在致命魔术###军阀趣史挪威的森林青春泉湿度爱情#恋爱小说天网追击我说你做#####铁血丹心#一直很安静#新娘与偏见#赤裸天使##图书馆员:寻找命运之矛####天黑黑###时间####超人奥特曼####台北晚9朝5恋爱假期杀人性机械#三男三女##地狱男爵#O记实录一个人
    ac_test.go:120: len: 100 out: #*#######*##**#**##*##**####****#**#*###**#**###****#***#####*#*#*#*##*####*###*####*####***#*##*#**
    ac_test.go:110: len: 808 input: #无懈可击##人龙传说那么爱你为什么红日#惊惧黑书#憨豆特派员#一路同行##三叶草#超时空接触四大名女之李香君红颜#切肤只爱###情欲九歌百个女生一个男生来生缘##背起爸爸上学##我恨我痴心恋爱高手#党员马大姐爱情的牙齿#魔高一丈不溶性侵犯独立日##书剑恩仇录三岔口说不出的告别好色妖姬杨贵妃玩命关头##星座全金属外壳候鸟#轰天潜舰###泳池情杀案仙女与骗子世界第一等关于世界,你知道个X#####因为是女子不弃今生乞丐王子叶子#不再回头恋爱小说六人行不由自主我等我在爆破爱情呼叫转移##18禁不禁海神波塞冬号#蜗牛#电台情歌一级戒备#水晶伤心太平洋爱笑的眼睛买办之家四十岁的老处男##春天在哪里##
    ac_test.go:120: len: 100 out: #*##***#*#*#*##*#***#*###***##*##**#**#***##*****##***#*###****#####****#*******##**#*#**#*****##*##
    ac_test.go:110: len: 694 input: #说不出的告别##花样中年灌篮高手#从海底出击原来你也在这里#阿宝的故事不良主妇日记###欲望桃花#新长征路上的摇滚性昏迷沧海一声笑百万富翁的初恋有完没完哥斯拉小狼狗##乌鸦#阳光总在风雨后勾魂地堡#花与剑##七剑下天山感恩的心关于世界,你知道个X#全面包围###好男人##泳池诱惑##追逐自由#孽恋焚情#时光骇客就是我#正义前锋加速度无法阻挡的婚姻#不溶性侵犯#####好想你琥珀之歌#美人计六面埋伏#一级病毒无地自容#80天环游世界#为你燃烧####8MM#一起喝彩黑鹰坠落老师不是人哥斯拉#晴天#血战阿拉曼###想谈恋爱的鱼#####
    ac_test.go:120: len: 100 out: #*##**#**#**###*#*******##*#**#*##***#*###*##*##*#*#**#***#*#####**#**#**#*#*####*#****#*#*###*#####
	... [output truncated]
PASS

```

