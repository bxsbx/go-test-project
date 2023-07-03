package questions

import (
	"StandardProject/sync/ebag/mongodb"
	"StandardProject/sync/ebag/redis"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

const (
	HOMEWORK = "t_cache_question_homework_student"
	LESSON   = "t_cache_question_lesson_student"
)

type StuWrongQuestionItem struct {
	F_id                  string `json:"_id"          bson:"_id"`                    //id
	F_student_id          string `json:"F_student_id"          bson:"F_student_id"`  //学生id
	F_resource_id         int    `json:"F_resource_id"         bson:"F_resource_id"` //资源id
	F_subject_id          int    `json:"F_subject_id"          bson:"F_subject_id"`  //科目id
	F_class_id            string `json:"F_class_id"            bson:"F_class_id"`    //班级id
	F_school_id           int    `json:"F_school_id"           bson:"F_school_id"`   //学校id
	F_moment              int    `json:"F_moment"              bson:"F_moment"`
	F_type_detail         int    `json:"F_type_detail"         bson:"F_type_detail"`         //资源类型
	F_student_right_times int    `json:"F_student_right_times" bson:"F_student_right_times"` //学生正确次数
	F_student_wrong_times int    `json:"F_student_wrong_times" bson:"F_student_wrong_times"` //学生错误次数
	F_question_type       int    `json:"F_question_type"       bson:"F_question_type"`       //题目类型（目前仅试卷题目有值，其余资源该值为0）
	F_date                string `json:"F_date"                bson:"F_date"`                //日期
}

// 从大数据中获取数据
func GetQsFromBigData(moment, page, size int) ([]StuWrongQuestionItem, error) {
	collName := HOMEWORK
	if moment == 1 {
		collName = LESSON
	}
	coll := mongodb.MongoDB.Collection(collName)

	where := bson.D{
		{"F_wrong_times", bson.D{{"$gt", 0}}},
		{"F_is_conquer_deleted", bson.D{{"$ne", 1}}},
		{"F_date", bson.D{{"$gte", "2023-01-01"}}},
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$match", where}},
		bson.D{{"$skip", (page - 1) * size}},
		bson.D{{"$limit", size}},
	}

	var list = make([]StuWrongQuestionItem, 0)

	cur, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		fmt.Println("GetQsFromBigData——err:", err)
		return nil, err
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		item := StuWrongQuestionItem{}
		err = cur.Decode(&item)
		if err != nil {
			fmt.Println("GetQsIdFromBigData——err:", err)
			continue
		}
		list = append(list, item)
	}
	return list, nil
}

// 从大数据中获取所有题目id
func GetQsIdFromBigData(moment int) ([]int, error) {
	collName := HOMEWORK
	if moment == 1 {
		collName = LESSON
	}
	coll := mongodb.MongoDB.Collection(collName)

	where := bson.D{
		{"F_wrong_times", bson.D{{"$gt", 0}}},
		{"F_is_conquer_deleted", bson.D{{"$ne", 1}}},
		{"F_date", bson.D{{"$gte", "2023-01-01"}}},
		{"F_student_id", bson.D{{"$regex", "^[0-9]-[0-9]{4}-[0-9]-"}}}, //匹配智慧课堂的id
	}
	group := bson.D{
		{"_id", "$F_resource_id"},
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$match", where}},
		bson.D{{"$group", group}},
	}

	var list = make([]int, 0)

	cur, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		fmt.Println("GetQsIdFromBigData——err:", err)
		return nil, err
	}
	type temp struct {
		F_resource_id int `json:"_id"          bson:"_id"` //资源id
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		item := temp{}
		err = cur.Decode(&item)
		if err != nil {
			fmt.Println("GetQsIdFromBigData——err:", err)
			continue
		}
		list = append(list, item.F_resource_id)
	}
	return list, nil
}

// 将中山小题id与智慧课堂小题id分离开
func GetZsAndEbagQsIds(qIds []int) (zsQsIds []string, ebagQsIds []string) {
	for _, v := range qIds {
		qId := strconv.Itoa(v)
		if len(qId) < 16 {
			zsQsIds = append(zsQsIds, qId)
		} else {
			ebagQsIds = append(ebagQsIds, qId)
		}
	}
	return
}

// 过滤掉已存在的id
func FilerExistIds(ids []string, key string) (newIds []string) {
	for _, id := range ids {
		exists, _ := redis.RedisObj.Exists(key + id)
		if !exists {
			newIds = append(newIds, id)
		}
	}
	return
}

// 获取学生schoolId为0的数据
func GetStuIdsFromBigDataErrSchoolId(moment int) ([]string, error) {
	collName := HOMEWORK
	if moment == 1 {
		collName = LESSON
	}
	coll := mongodb.MongoDB.Collection(collName)

	where := bson.D{
		{"F_school_id", bson.D{{"$lte", 0}}},
		{"F_student_id", bson.D{{"$regex", "^[0-9]-[0-9]{4}-[0-9]-"}}}, //匹配智慧课堂的id
	}
	group := bson.D{
		{"_id", "$F_student_id"},
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$match", where}},
		bson.D{{"$group", group}},
	}

	var list = make([]string, 0)

	cur, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		fmt.Println("GetQsIdFromBigData——err:", err)
		return nil, err
	}
	type temp struct {
		F_student_id string `json:"_id"          bson:"_id"` //学生id
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		item := temp{}
		err = cur.Decode(&item)
		if err != nil {
			fmt.Println("GetQsIdFromBigData——err:", err)
			continue
		}
		list = append(list, item.F_student_id)
	}
	return list, nil
}
