package config

type PushLanguageCfg struct {
	Id       int32
	HandleId int32
	Lang     int32
	Msg      string
}

type PushLanguageTable struct {
	data                                map[int32]*PushLanguageCfg
	pushLanguageHandleIdIndexMap        map[int32][]int32
	pushLanguageHandleIdAndLangIndexMap map[pushLanguageHandleIdAndLangIndex][]int32
}

var PushLanguageData = &PushLanguageTable{
	data: map[int32]*PushLanguageCfg{},
	pushLanguageHandleIdIndexMap: map[int32][]int32{
		1: {1, 2, 3, 4, 5, 6, 7, 8},
		2: {9, 10, 11, 12, 13, 14, 15, 16},
		3: {17, 18, 19, 20, 21, 22, 23, 24},
		4: {25, 26, 27, 28, 29, 30, 31, 32},
		5: {33, 34, 35, 36, 37, 38, 39, 40},
		6: {41, 42, 43, 44, 45, 46, 47, 48},
	},
	pushLanguageHandleIdAndLangIndexMap: map[pushLanguageHandleIdAndLangIndex][]int32{
		{1, 0}:  {2},
		{1, 11}: {6},
		{1, 13}: {8},
		{1, 15}: {7},
		{1, 1}:  {1},
		{1, 5}:  {5},
		{1, 8}:  {4},
		{1, 9}:  {3},
		{2, 0}:  {10},
		{2, 11}: {14},
		{2, 13}: {16},
		{2, 15}: {15},
		{2, 1}:  {9},
		{2, 5}:  {13},
		{2, 8}:  {12},
		{2, 9}:  {11},
		{3, 0}:  {18},
		{3, 11}: {22},
		{3, 13}: {24},
		{3, 15}: {23},
		{3, 1}:  {17},
		{3, 5}:  {21},
		{3, 8}:  {20},
		{3, 9}:  {19},
		{4, 0}:  {26},
		{4, 11}: {30},
		{4, 13}: {32},
		{4, 15}: {31},
		{4, 1}:  {25},
		{4, 5}:  {29},
		{4, 8}:  {28},
		{4, 9}:  {27},
		{5, 0}:  {34},
		{5, 11}: {38},
		{5, 13}: {40},
		{5, 15}: {39},
		{5, 1}:  {33},
		{5, 5}:  {37},
		{5, 8}:  {36},
		{5, 9}:  {35},
		{6, 0}:  {42},
		{6, 11}: {46},
		{6, 13}: {48},
		{6, 15}: {47},
		{6, 1}:  {41},
		{6, 5}:  {45},
		{6, 8}:  {44},
		{6, 9}:  {43},
	},
}

func (table *PushLanguageTable) Get(id int32) *PushLanguageCfg {
	return table.data[id]
}

func (table *PushLanguageTable) GetAll() []int32 {
	return pushLanguageKeys
}

func (table *PushLanguageTable) GetAllData() []*PushLanguageCfg {
	return pushLanguageValues
}

var pushLanguageValues = []*PushLanguageCfg{
	{
		Id:       1,
		HandleId: 1,
		Lang:     1,
		Msg:      "亲爱的大师，铁砧噪鸣，亟待你锻造出绝世装备！",
	},
	{
		Id:       2,
		HandleId: 1,
		Lang:     0,
		Msg:      "Dear Master, the anvil is calling you to forge legendary equipment!",
	},
	{
		Id:       3,
		HandleId: 1,
		Lang:     9,
		Msg:      "かなとこが呼んでいるぞ！より強力な装備を鍛冶しよう！",
	},
	{
		Id:       4,
		HandleId: 1,
		Lang:     8,
		Msg:      "모루가 울려 대고 있습니다. 더욱 강력한 장비를 단조해 보세요!",
	},
	{
		Id:       5,
		HandleId: 1,
		Lang:     5,
		Msg:      "¡Querido maestro, el yunque te llama a que forjes equipo legendario!",
	},
	{
		Id:       6,
		HandleId: 1,
		Lang:     11,
		Msg:      "Mestre, a bigorna está convocando para forjar equipamentos lendários!",
	},
	{
		Id:       7,
		HandleId: 1,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก ทั่งตีเหล็กกำลังเรียกร้องให้คุณไปหลอมอุปกรณ์เลเจนด์นะ!",
	},
	{
		Id:       8,
		HandleId: 1,
		Lang:     13,
		Msg:      "親愛的大師，鐵砧噪鳴，亟待你鍛造出絕世裝備！",
	},
	{
		Id:       9,
		HandleId: 2,
		Lang:     1,
		Msg:      "亲爱的大师，神器藏宝阁收益已满，快回来领取吧~",
	},
	{
		Id:       10,
		HandleId: 2,
		Lang:     0,
		Msg:      "Dear Master, earnings in the Treasure House are full, come back and claim them now!",
	},
	{
		Id:       11,
		HandleId: 2,
		Lang:     9,
		Msg:      "神器宝物庫の収益が最大になったぞ！早く受領しよう！",
	},
	{
		Id:       12,
		HandleId: 2,
		Lang:     8,
		Msg:      "아티팩트 보물각 수익이 가득 찼습니다. 어서 수령하세요.",
	},
	{
		Id:       13,
		HandleId: 2,
		Lang:     5,
		Msg:      "Querido maestro, las ganancias de la Casa del tesoro llegaron al límite. ¡Regresa y reclámalas ahora!",
	},
	{
		Id:       14,
		HandleId: 2,
		Lang:     11,
		Msg:      "Mestre, os ganhos na Casa do Tesouro estão cheios. Volte para resgatar!",
	},
	{
		Id:       15,
		HandleId: 2,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก รางวัลตอบแทนในบ้านสมบัติเต็มแล้ว กลับมารับตอนนี้เลยสิ!",
	},
	{
		Id:       16,
		HandleId: 2,
		Lang:     13,
		Msg:      "親愛的大師，神器藏寶閣收益已滿，快回來領取吧~",
	},
	{
		Id:       17,
		HandleId: 3,
		Lang:     1,
		Msg:      "亲爱的大师，猩红回廊内的宝箱可以领取了~",
	},
	{
		Id:       18,
		HandleId: 3,
		Lang:     0,
		Msg:      "Dear Master, chest(s) in the Scarlet Corridor is ready for claim!",
	},
	{
		Id:       19,
		HandleId: 3,
		Lang:     9,
		Msg:      "緋色の回廊の宝箱が解錠されたぞ！早速開けてみよう！",
	},
	{
		Id:       20,
		HandleId: 3,
		Lang:     8,
		Msg:      "선홍빛 회랑의 보물 상자를 수령할 수 있습니다.",
	},
	{
		Id:       21,
		HandleId: 3,
		Lang:     5,
		Msg:      "¡Querido maestro, puedes reclamar cofre(s) en el Pasillo rojo!",
	},
	{
		Id:       22,
		HandleId: 3,
		Lang:     11,
		Msg:      "Mestre, os baús no Corredor Rubi estão prontos para serem resgatados!",
	},
	{
		Id:       23,
		HandleId: 3,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก หีบในทางเดินแดงฉานพร้อมให้รับได้แล้วนะ!",
	},
	{
		Id:       24,
		HandleId: 3,
		Lang:     13,
		Msg:      "親愛的大師，猩紅回廊內的寶箱可以領取了~",
	},
	{
		Id:       25,
		HandleId: 4,
		Lang:     1,
		Msg:      "亲爱的大师，寻宝灯油已满，快回来挖掘更多强大的宝石吧！",
	},
	{
		Id:       26,
		HandleId: 4,
		Lang:     0,
		Msg:      "Dear Master, kerosene for the Trove Hunt is full. Come back and dig for more powerful gems!",
	},
	{
		Id:       27,
		HandleId: 4,
		Lang:     9,
		Msg:      "宝探しの灯油が満タンになったぞ！早く宝石を掘り出そう！",
	},
	{
		Id:       28,
		HandleId: 4,
		Lang:     8,
		Msg:      "보물찾기 등유가 가득 찼습니다. 더욱 강력한 보석을 발굴해 보세요!",
	},
	{
		Id:       29,
		HandleId: 4,
		Lang:     5,
		Msg:      "Querido maestro, el querosén para la Caza de tesoro llegó al límite. ¡Regresa y extrae gemas más poderosas!",
	},
	{
		Id:       30,
		HandleId: 4,
		Lang:     11,
		Msg:      "Mestre, a querosene para a Caça-tesouro está cheia. Volte e procure mais gemas!",
	},
	{
		Id:       31,
		HandleId: 4,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก น้ำมันก๊าดสำหรับล่าขุมทรัพย์เต็มแล้ว กลับมาขุดหาอัญมณีที่ทรงพลังเพิ่มอีกเถอะ!",
	},
	{
		Id:       32,
		HandleId: 4,
		Lang:     13,
		Msg:      "親愛的大師，尋寶燈油已滿，快回來挖掘更多強大的寶石吧！",
	},
	{
		Id:       33,
		HandleId: 5,
		Lang:     1,
		Msg:      "亲爱的大师，锻造铁砧升级完成，可以锻造出更强大的装备了！快回来看看吧~",
	},
	{
		Id:       34,
		HandleId: 5,
		Lang:     0,
		Msg:      "Dear Master, the forging anvil has been successfully upgraded and more powerful equipment is waiting ahead! Check it out now!",
	},
	{
		Id:       35,
		HandleId: 5,
		Lang:     9,
		Msg:      "かなとこがレベルアップしたぞ！より強い装備を鍛冶しよう！",
	},
	{
		Id:       36,
		HandleId: 5,
		Lang:     8,
		Msg:      "단조 모루가 업그레이드되었습니다. 더욱 강력한 장비를 단조해 보세요!",
	},
	{
		Id:       37,
		HandleId: 5,
		Lang:     5,
		Msg:      "¡Querido maestro, el yunque de forja se ha mejorado y ahora te espera equipo más poderoso! ¡Compruébalo ahora!",
	},
	{
		Id:       38,
		HandleId: 5,
		Lang:     11,
		Msg:      "Mestre, a bigorna de forja foi aprimorada com sucesso e equipamentos mais poderosos podem ser seus! Confira!",
	},
	{
		Id:       39,
		HandleId: 5,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก ทั่งตีเหล็กอัปเกรดสำเร็จแล้วนะ อุปกรณ์ที่ทรงพลังขึ้นอีกกำลังรออยู่! ไปดูกันเถอะ!",
	},
	{
		Id:       40,
		HandleId: 5,
		Lang:     13,
		Msg:      "親愛的大師，鍛造鐵砧升級完成，可以鍛造出更強大的裝備了！快回來看看吧~",
	},
	{
		Id:       41,
		HandleId: 6,
		Lang:     1,
		Msg:      "亲爱的大师，宝石藏馆升级完成，可以收藏更多宝石了！快回来看看吧~",
	},
	{
		Id:       42,
		HandleId: 6,
		Lang:     0,
		Msg:      "Dear Master, the Gem Collection has been successfully upgraded and you can collect more gems! Check it out now!",
	},
	{
		Id:       43,
		HandleId: 6,
		Lang:     9,
		Msg:      "宝石収蔵庫がレベルアップしたぞ！より多くの宝石を収納しよう！",
	},
	{
		Id:       44,
		HandleId: 6,
		Lang:     8,
		Msg:      "보석 창고가 업그레이드되었습니다. 어서 와 보세요~",
	},
	{
		Id:       45,
		HandleId: 6,
		Lang:     5,
		Msg:      "Querido maestro, la lista de gemas se ha mejorado y ahora puedes tener más gemas! ¡Compruébalo ahora!",
	},
	{
		Id:       46,
		HandleId: 6,
		Lang:     11,
		Msg:      "Mestre, a Coleção Gemas foi aprimorada com sucesso e você pode coletar mais gemas! Confira já!",
	},
	{
		Id:       47,
		HandleId: 6,
		Lang:     15,
		Msg:      "มาสเตอร์ที่รัก คอลเลกชันอัญมณีอัปเกรดสำเร็จแล้วนะ คุณจะรวบรวมอัญมณีได้มากขึ้น! ไปดูกันเถอะ!",
	},
	{
		Id:       48,
		HandleId: 6,
		Lang:     13,
		Msg:      "親愛的大師，寶石藏館升級完成，可以收藏更多寶石了！快回來看看吧~",
	},
}

var pushLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48,
}

func init() {
	PushLanguageData.data = make(map[int32]*PushLanguageCfg)
	for i := 0; i < len(pushLanguageKeys); i++ {
		PushLanguageData.data[pushLanguageKeys[i]] = pushLanguageValues[i]
	}
}

func (table *PushLanguageTable) GetByHandleId(HandleId int32) (res []*PushLanguageCfg) {
	for _, i := range table.pushLanguageHandleIdIndexMap[HandleId] {
		res = append(res, table.Get(i))
	}
	return
}

func (table *PushLanguageTable) GetByHandleIdAndLang(HandleId int32, Lang int32) (res []*PushLanguageCfg) {
	for _, i := range table.pushLanguageHandleIdAndLangIndexMap[pushLanguageHandleIdAndLangIndex{HandleId, Lang}] {
		res = append(res, table.Get(i))
	}
	return
}

type pushLanguageHandleIdAndLangIndex struct {
	handleId int32
	lang     int32
}
