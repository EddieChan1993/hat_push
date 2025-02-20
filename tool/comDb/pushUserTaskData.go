package comDb

/**
玩家阵容
*/
import (
	"context"
	"fmt"
	"git.dhgames.cn/svr_comm/gcore/glog"
	"git.dhgames.cn/svr_comm/gcore/gmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hat_push/pbgo/pbpush"
	"hat_push/tool/consul/static"
)

const limitMax = 500
const pushUserTasks = "push_user_tasks"

type PushTaskDbFileterTyp = bson.M

func initPushTaskDb() {
	modles := make([]mongo.IndexModel, 0)
	modles = append(modles, mongo.IndexModel{Keys: bson.M{"account": 1}})
	modles = append(modles, mongo.IndexModel{Keys: bson.M{"handletype": 1}})
	modles = append(modles, mongo.IndexModel{Keys: bson.M{"sendat": 1}})
	_, err := gmongo.CreateIndexes(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), modles)
	if err != nil {
		glog.Panic(fmt.Sprintf("url %s db %s table %s err %v", static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), err))
	}
}

func PushTaskAccount(account int32) PushTaskDbFileterTyp {
	return bson.M{"account": account}
}

func PushTaskFileter(account, taskType int32) PushTaskDbFileterTyp {
	f1 := bson.M{"account": account}
	f2 := bson.M{"handletype": taskType}
	filter := bson.M{"$and": []bson.M{f1, f2}}
	return filter
}

func PushTaskIsSameFileter(account, taskType int32, sendAt int64) PushTaskDbFileterTyp {
	f1 := bson.M{"account": account}
	f2 := bson.M{"handletype": taskType}
	f3 := bson.M{"sendat": sendAt}
	filter := bson.M{"$and": []bson.M{f1, f2, f3}}
	return filter
}

func FindTasksBeforeAt(endAt int64) []*pbpush.PushUserTask {
	data := make([]*pbpush.PushUserTask, 0, limitMax)
	fAll := bson.M{"sendat": bson.M{"$lt": endAt}}
	option := options.Find().SetSort(bson.D{{"sendat", -1}}).SetLimit(limitMax)
	cur, err := gmongo.Find(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), fAll, option)
	if err != nil {
		glog.Errorf("FindAll Find fail err %v", err)
		return nil
	}
	if err = cur.All(context.TODO(), &data); err != nil {
		glog.Errorf("FindAll cur All fail err %v", err)
		return nil
	}
	return data
}

// SetTopDigestBotCamp 存储阵营信息
func SetTopDigestBotCamp(data *pbpush.PushUserTask) error {
	filter := PushTaskFileter(data.Account, data.HandleType)
	res, err := findOneTopDigestBotCamp(filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = insertPushTask(data)
			if err != nil {
				glog.Error(err)
				return err
			}
		} else {
			glog.Error(err)
			return err
		}
	} else {
		if res != nil {
			err = UpdatePushTask(filter, data)
			if err != nil {
				glog.Error(err)
				return err
			}
		}
	}
	return err
}

func findOneTopDigestBotCamp(filter interface{}) (*mongo.SingleResult, error) {
	return gmongo.FindOne(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), filter)
}

func insertPushTask(data *pbpush.PushUserTask) error {
	_, err := gmongo.InsertOne(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), data)
	if err != nil {
		return fmt.Errorf("insert err %v", data)
	}
	return nil
}

func UpdatePushTask(where interface{}, set interface{}) error {
	_, err := gmongo.UpdateOne(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), where, bson.M{"$set": set})
	if err != nil {
		return fmt.Errorf("UpdatePushTask err %v", err)
	}
	return nil
}

func DelPushTask(filter PushTaskDbFileterTyp) {
	_, err := gmongo.DeleteMany(static.StaticCommonDBUrl(), static.StaticCommonDb(), collection(pushUserTasks), filter)
	if err != nil {
		glog.Errorf("DelPushTask fail err %v", err)
		return
	}
	return
}
