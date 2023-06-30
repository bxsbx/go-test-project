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
	StudentID   string  // 学生id
	QuestionID  string  // 题目id
	CorrectNum  int     // 正确次数
	ErrorNum    int     // 错误次数
	ScoreRate   float64 // 得分/正确率
	MasterLevel int     // 题目掌握程度
	SchoolID    int     // 学校id
	ClassID     string  // 班级id
	SubjectID   int     // 学科id
	From        int     // 来源,目前有作业系统/阅卷系统
	WrongTag    int     // 错题标签(0.暂无标签,1.概念模糊,2.思路错误,3.审题错误,4.粗心大意,5.完全不会,6.其他原因)
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsCollect   int
	IsWrong     int
	//用于区分 错题1 和 学科网2 题目
	//Origin int `gorm:"column:origin" json:"-"`
	EbagTag int // 智慧课堂错题标签(1.课前,2.课中,3.课后)
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
	if err != nil {
		sqlRecord := fmt.Sprintf(strings.ReplaceAll(sql, "?", "%v"), args...)
		redis.RedisObj.Set(redis.SQL_INSERT+questions.StudentID+questions.QuestionID, sqlRecord)
	}
	return
}

// 更新学生错题记录
func UpdateWrongQs(studentID, questionID string, createTime time.Time, db *gorm.DB) (err error) {
	sql := "update student_wrong_questions set error_num = error_num + ?, create_at = ?, updated_at = ? where student_id = ? and question_id = ?"
	err = db.Exec(sql, 1, createTime, time.Now(), studentID, questionID).Error
	if err != nil {
		sqlRecord := fmt.Sprintf(strings.ReplaceAll(sql, "?", "%v"), 1, studentID, questionID)
		redis.RedisObj.Set(redis.SQL_UPDATE+studentID+questionID, sqlRecord)
	}
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

func BatchHandle(list []StuWrongQuestionItem) {
	var wg sync.WaitGroup
	for i := range list {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			studentId := list[j].F_student_id
			questionId := strconv.Itoa(list[j].F_resource_id)
			oneQs, _ := FindOneWrongQs(studentId, questionId, mysql.MysqlDB)
			if oneQs.StudentID != "" {
				question := StudentWrongQuestions{
					StudentID:   studentId,
					QuestionID:  questionId,
					CorrectNum:  0,
					ErrorNum:    1,
					ScoreRate:   0,
					MasterLevel: 0,
					SchoolID:    list[j].F_school_id,
					ClassID:     list[j].F_class_id,
					SubjectID:   list[j].F_subject_id,
					From:        4,
					WrongTag:    0,
					EbagTag:     list[j].F_moment + 1,
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

				questionDate, _ := time.Parse("2006-01-02", list[j].F_date)
				question.CreatedAt = questionDate
				err := CreateWrongQs(question, mysql.MysqlDB)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				questionDate, _ := time.Parse("2006-01-02", list[j].F_date)
				if oneQs.CreatedAt.Before(questionDate) {
					questionDate = oneQs.CreatedAt
				}
				err := UpdateWrongQs(studentId, questionId, questionDate, mysql.MysqlDB)
				if err != nil {
					fmt.Println(err)
				}
			}
		}(i)
	}
	wg.Wait()
}

// 重新执行未成功的sql
func ExecSqlAgain(keys []string) {
	var wg sync.WaitGroup
	tx := mysql.MysqlDB.Begin()
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			wg.Done()
			sql, _ := redis.RedisObj.GetString(k)
			err := tx.Exec(sql).Error
			if err == nil {
				redis.RedisObj.Remove(k)
			}
		}(key)
	}
	wg.Wait()
	tx.Commit()
}
