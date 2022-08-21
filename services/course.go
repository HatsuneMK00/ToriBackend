package services

import (
	"ToriBackend/entity"
	"ToriBackend/global"
)

type CourseService interface {
	GetAllCourses() ([]entity.Course, bool)
	GetCourse(id uint) (*entity.Course, bool)
	EnrollCourse(userId, courseId uint) bool
	GetEnrolledCourses(userId uint) ([]entity.Course, bool)
}

type courseService struct {
}

func (c courseService) GetAllCourses() ([]entity.Course, bool) {
	var courses []entity.Course
	err := global.MysqlDB.Find(&courses).Error
	ok := true
	if err != nil {
		global.Logger.Errorf("%v", err)
		ok = false
	}
	return courses, ok
}

func (c courseService) GetCourse(id uint) (*entity.Course, bool) {
	var course entity.Course
	course.ID = id
	err := global.MysqlDB.First(&course).Error
	ok := true
	if err != nil {
		global.Logger.Errorf("%v", err)
		ok = false
	}
	return &course, ok
}

func (c courseService) EnrollCourse(userId, courseId uint) bool {
	var user entity.User
	var course entity.Course
	user.ID = userId
	course.ID = courseId
	err := global.MysqlDB.Model(&user).Association("Courses").Append(&course)
	if err != nil {
		global.Logger.Errorf("enroll course failed: %v", err)
		return false
	}
	return true
}

func (c courseService) GetEnrolledCourses(userId uint) ([]entity.Course, bool) {
	var courses []entity.Course
	err := global.MysqlDB.
		Table("users").
		Where("users.id = ?", userId).
		Joins("INNER JOIN user_courses uc ON users.id = uc.user_id").
		Joins("INNER JOIN courses on uc.course_id = courses.id").
		Select("courses.*, uc.created_at as enrolled_at").
		Order("enrolled_at desc").
		Find(&courses).Error
	ok := true
	if err != nil {
		global.Logger.Errorf("get enrolled courses failed: %v", err)
		ok = false
	}
	return courses, ok
}
