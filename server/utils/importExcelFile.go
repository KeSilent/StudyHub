/*
 * @Author: Yang
 * @Date: 2023-05-24 18:08:16
 * @Description: 解析Excel文件
 */
package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KeSilent/study-hub/server/model/edu_organization"
	"github.com/KeSilent/study-hub/server/model/system/request"
	"github.com/tealeg/xlsx"
)

/**
 * @Description: 解析Excel
 * @param {string} excelFileName
 * @return {*}
 */
func EnalysisExcel(excelFileName string, courses []edu_organization.EduCourse) (rs []request.Register) {
	// 打开Excel文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历所有的Sheet
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet名称: %s\n", sheet.Name)
		findCourses := FindCoursesByCourseName(courses, strings.TrimSpace(sheet.Name))
		if len(findCourses) == 0 {
			continue
		}

		// 遍历Sheet中的行
		for _, row := range sheet.Rows {
			var r request.Register

			r.NickName = row.Cells[0].String()
			r.Username = row.Cells[0].String()
			r.Course = int(findCourses[0].ID)

			if row.Cells[1].String() == "" {
				r.TotalSessions = 0
				r.RemainingSessions = 0
			} else {
				r.TotalSessions, _ = strconv.Atoi(row.Cells[1].String())
				r.RemainingSessions = r.TotalSessions
			}
			rs = append(rs, r)
		}
	}

	return rs
}

func FindCoursesByCourseName(courses []edu_organization.EduCourse, courseName string) []edu_organization.EduCourse {
	var matchingCourses []edu_organization.EduCourse

	for _, course := range courses {
		if course.CourseName == courseName {
			matchingCourses = append(matchingCourses, course)
		}
	}

	return matchingCourses
}
