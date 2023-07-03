package questions

import (
	"StandardProject/sync/ebag/mysql"
	"StandardProject/sync/ebag/redis"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
	"time"
)

type StudentWrongQuestions struct {
	StudentID   string    `gorm:"column:student_id;primaryKey" json:"student_id"`   // 学生id
	QuestionID  string    `gorm:"column:question_id;primaryKey" json:"question_id"` // 题目id
	CorrectNum  int       `gorm:"column:correct_num" json:"correct_num"`            // 正确次数
	ErrorNum    int       `gorm:"column:error_num" json:"error_num"`                // 错误次数
	ScoreRate   float64   `gorm:"column:score_rate" json:"score_rate"`              // 得分/正确率
	MasterLevel int       `gorm:"column:master_level" json:"master_level"`          // 题目掌握程度
	SchoolID    int       `gorm:"column:school_id" json:"school_id"`                // 学校id
	ClassID     string    `gorm:"column:class_id" json:"class_id"`                  // 班级id
	SubjectID   int       `gorm:"column:subject_id" json:"subject_id"`              // 学科id
	From        int       `gorm:"column:from" json:"from"`                          // 来源,目前有作业系统/阅卷系统
	WrongTag    int       `gorm:"column:wrong_tag" json:"wrong_tag"`                // 错题标签(0.暂无标签,1.概念模糊,2.思路错误,3.审题错误,4.粗心大意,5.完全不会,6.其他原因)
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
	IsCollect   int       `gorm:"column:is_collect;default:-1" json:"is_collect"`
	IsWrong     int       `gorm:"column:is_wrong;default:-1" json:"is_wrong,omitempty"`
	//用于区分 错题1 和 学科网2 题目
	//Origin int `gorm:"column:origin" json:"-"`
	EbagTag int `gorm:"column:zh_tag" json:"-"` // 智慧课堂错题标签(1.课前,2.课中,3.课后)
}

type RecordSql struct {
	ExecSql string        `json:"exec_sql"`
	Sql     string        `json:"sql"`
	Args    []interface{} `json:"args"`
}

func genRecordSql(sql string, args []interface{}) string {
	sqlRecord := fmt.Sprintf(strings.ReplaceAll(sql, "?", "%v"), args...)
	record := RecordSql{
		sqlRecord,
		sql,
		args,
	}
	marshal, _ := json.Marshal(record)
	return string(marshal)
}

// 创建学生错题记录
func CreateWrongQs(questions StudentWrongQuestions, db *gorm.DB) (err error) {
	fileds := []string{"student_id", "question_id", "correct_num", "error_num",
		"score_rate", "master_level", "school_id", "class_id", "subject_id",
		"`from`", "wrong_tag", "created_at", "updated_at", "is_collect", "is_wrong", "zh_tag"}

	questionMark := strings.TrimSuffix(strings.Repeat("?,", len(fileds)), ",")
	sql := fmt.Sprintf("insert into student_wrong_questions (%s) values (%s)", strings.Join(fileds, ","), questionMark)

	args := []interface{}{questions.StudentID, questions.QuestionID, questions.CorrectNum, questions.ErrorNum,
		questions.ScoreRate, questions.MasterLevel, questions.SchoolID, questions.ClassID, questions.SubjectID,
		questions.From, questions.WrongTag, questions.CreatedAt, questions.UpdatedAt, questions.IsCollect, questions.IsWrong, questions.EbagTag}

	err = db.Exec(sql, args...).Error
	//if err != nil {
	//	redis.RedisObj.Set(redis.SQL_INSERT+questions.StudentID+questions.QuestionID, genRecordSql(sql, args))
	//}
	return
}

// 更新学生错题记录
func UpdateWrongQs(studentID, questionID string, error_num int, createTime time.Time, db *gorm.DB) (err error) {
	sql := "update student_wrong_questions set error_num = error_num + ?, create_at = ?, updated_at = ? where student_id = ? and question_id = ?"
	args := []interface{}{error_num, createTime, time.Now(), studentID, questionID}
	err = db.Exec(sql, args...).Error
	//if err != nil {
	//	redis.RedisObj.Set(redis.SQL_UPDATE+studentID+questionID, genRecordSql(sql, args))
	//}
	return
}

// 查找学生错题记录
func FindOneWrongQs(studentID, questionID string, db *gorm.DB) (question StudentWrongQuestions, err error) {
	err = db.Table("student_wrong_questions").Where("student_id = ? and question_id = ?", studentID, questionID).First(&question).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return question, nil
}

// 查看是否存在学生错题记录
func IsExistWrongQs(studentID, questionID string, db *gorm.DB) (exist bool, err error) {
	var count int
	err = db.Table("student_wrong_questions").Where("student_id = ? and question_id = ?", studentID, questionID).Count(&count).Error
	if count > 0 {
		exist = true
	}
	return
}

func Handle(item StuWrongQuestionItem) (err error) {
	tx := mysql.MysqlDB.Begin()
	studentId := item.F_student_id
	questionId := strconv.Itoa(item.F_resource_id)
	oneQs, err := FindOneWrongQs(studentId, questionId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	if oneQs.StudentID != "" {
		question := StudentWrongQuestions{
			StudentID:   studentId,
			QuestionID:  questionId,
			CorrectNum:  0,
			ErrorNum:    1,
			ScoreRate:   0,
			MasterLevel: 0,
			SchoolID:    item.F_school_id,
			ClassID:     item.F_class_id,
			SubjectID:   item.F_subject_id,
			From:        4,
			WrongTag:    0,
			EbagTag:     item.F_moment + 1,
			IsCollect:   -1,
			IsWrong:     -1,
			UpdatedAt:   time.Now(),
		}
		if question.SchoolID == 0 {
			jsonData, _ := redis.RedisObj.GetString(redis.STUDENT_SCHOOL + question.StudentID)
			if jsonData != "" {
				var student Student
				json.Unmarshal([]byte(jsonData), &student)
				if student.SchoolId <= 0 {
					return
				}
				question.SchoolID = student.SchoolId
			}
		}
		if len(questionId) < 16 {
			jsonData, _ := redis.RedisObj.GetString(redis.ZS_QUESTIONS + questionId)
			if jsonData != "" {
				var zsQs ZSQuestion
				json.Unmarshal([]byte(jsonData), &zsQs)
				if zsQs.Subject > 0 {
					question.SubjectID = zsQs.Subject
				}
			}
		} else {
			jsonData, _ := redis.RedisObj.GetString(redis.EBAG_QUESTIONS + questionId)
			if jsonData != "" {
				var ebagQs EbagQuestion
				json.Unmarshal([]byte(jsonData), &ebagQs)
				if ebagQs.Subject > 0 {
					question.SubjectID = ebagQs.Subject
				}
			}
		}

		questionDate, _ := time.Parse("2006-01-02", item.F_date)
		question.CreatedAt = questionDate
		err := CreateWrongQs(question, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	} else {
		// 由于之后可能要同步后面时间的数据，得将创建时间提前
		questionDate, _ := time.Parse("2006-01-02", item.F_date)
		if oneQs.CreatedAt.Before(questionDate) {
			questionDate = oneQs.CreatedAt
		}
		err := UpdateWrongQs(studentId, questionId, 1, questionDate, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*
1、取消事务执行，事务是顺序执行的，没办法充分利用连接数，一个事务只用到了一个连接
2、开启事务的方式，防止多并发导致插入不成功
*/
func BatchHandle(list []StuWrongQuestionItem, moment int) {
	var wg sync.WaitGroup
	for i := range list {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			err := Handle(list[j])
			if err != nil {
				fmt.Println(err)
				marshal, _ := json.Marshal(list[j])
				redis.RedisObj.Set(redis.FAIL_MONGO_QS+strconv.Itoa(moment)+"::"+list[j].F_id, string(marshal))
			}
		}(i)
	}
	wg.Wait()
}

//// 重新执行未成功的sql
//func ExecSqlAgain(keys []string) {
//	var wg sync.WaitGroup
//	tx := mysql.MysqlDB.Begin()
//	for _, key := range keys {
//		wg.Add(1)
//		go func(k string) {
//			wg.Done()
//			sql, _ := redis.RedisObj.GetString(k)
//			err := tx.Exec(sql).Error
//			if err == nil {
//				redis.RedisObj.Remove(k)
//			}
//		}(key)
//	}
//	wg.Wait()
//	tx.Commit()
//}

// 重新执行未成功保存记录
func ExecTxAgain(moment int) {
	var wg sync.WaitGroup
	keys, _ := redis.RedisObj.GetKeys(redis.FAIL_MONGO_QS + strconv.Itoa(moment) + "::*")
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			wg.Done()
			str, _ := redis.RedisObj.GetString(k)
			var item StuWrongQuestionItem
			json.Unmarshal([]byte(str), &item)
			if err := Handle(item); err == nil {
				redis.RedisObj.Remove(k)
			}
		}(key)
	}
	wg.Wait()
}

func BatchHandle2() {

}
