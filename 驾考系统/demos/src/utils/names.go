package utils


var(
	//姓氏          切片
	familyNames = []string{"赵","钱","孙","李","周","吴","郑","王","冯","陈","楚","卫","蒋","沈","韩","杨","张","欧阳","独孤","司徒","上官","诸葛","南宫"}
	//辈分(宗的永其光) keytype   valtype(切片)
	middleNamesMap = map[string][]string{}
	//名字
	lastNames = []string{"金","章","玉","简","麟","凤","征","祥","丰","功","伟","业","积","厚","流","光","克","昌","厥","后","兰","桂","芬","芳"}
)

//init函数用来初始化变量会执行在main函数之前
func init() {
	//for key,val :=range array{}
	for _,val :=range familyNames{
		//姓氏
		if val != "陈"{
			//辈分
			middleNamesMap[val] = []string{"大","朝","春","应","文","学","克","成","家","传","忠","世","友","贤","能"}
		}else{
			middleNamesMap[val] = []string{"宗","得","永","其","光"}
		}
	}
}

//定义一个函数,生成随机的名字
func GetRandomName()(name string){
	familyName := familyNames[GetRandomInt(0,len(familyNames)-1)]//随机姓氏
	middleName := middleNamesMap[familyName][GetRandomInt(0,len(middleNamesMap[familyName])-1)]//随机辈分
	lastName   := lastNames[GetRandomInt(0,len(lastNames)-1)]//随机名字
	return familyName+middleName+lastName
}
