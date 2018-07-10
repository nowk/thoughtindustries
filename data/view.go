package data

import (
	"time"
)

type View struct {
	CompanyId    string    `json:"companyId,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	CourseId     string    `json:"courseId,omitempty"`
	CourseTitle  string    `json:"courseTitle,omitempty"`
	CourseSku    string    `json:"courseSku,omitempty"`
	LessonId     string    `json:"lessonId,omitempty"`
	LessonTitle  string    `json:"lessonTitle,omitempty"`
	LessonSlug   string    `json:"lessonSlug,omitempty"`
	SectionId    string    `json:"sectionId,omitempty"`
	SectionTitle string    `json:"sectionTitle,omitempty"`
	SectionSlug  string    `json:"sectionSlug,omitempty"`
	TopicId      string    `json:"topicId,omitempty"`
	TopicTitle   string    `json:"topicTitle,omitempty"`
	User         string    `json:"user,omitempty"`

	UserDetail UserDetail `json:"userDetail"`
}

/*
 * {
 *   "companyId": "66b131f7-33b7-4d5a-b0fa-bed64acd4c4e",
 *   "timestamp": "2018-02-05T02:49:17.000Z",
 *   "courseId": "4830c1e4-5b79-4745-90ac-3533e3e64a7b",
 *   "courseTitle": "Test Course",
 *   "courseSku": null,
 *   "lessonId": "1d9b510e-bd53-4172-b21f-5001cdd4da1c",
 *   "lessonTitle": "Test Lesson",
 *   "lessonSlug": "test-lesson",
 *   "sectionId": "92c056f7-6e22-4429-89e5-bdf89b1db8c1",
 *   "sectionTitle": "Test Section",
 *   "sectionSlug": "test-section",
 *   "topicId": "b7696a47-d9e6-41a8-be24-e431881ddfc0",
 *   "topicTitle": "Test Page",
 *   "user": "test@thoughtindustries.com",
 *   "userDetail": {
 *     "id": "41756c5f-5a2f-477e-80f7-086cb7146f59",
 *     "email": "test@thoughtindustries.com",
 *     "firstName": "Test",
 *     "lastName": "User",
 *     "externalCustomerId": null,
 *     "sfContactId": null,
 *     "sfAccountId": null,
 *     "ref1": null,
 *     "ref2": null,
 *     "ref3": null,
 *     "ref4": null,
 *     "ref5": null,
 *     "ref6": null,
 *     "ref7": null,
 *     "ref8": null,
 *     "ref9": null,
 *     "ref10": null
 *   }
 * }
 *
 */
