package wk

import "time"

type Assignments struct {
	StartedAt time.Time `json:"started_at"`
}

type Reviews struct {
	Reviews struct {
		AssignmentId      int       `json:"assignment_id"`
		IncorrectMeanings int       `json:"incorrect_meaning_answers"`
		IncorrectReadings int       `json:"incorrect_reading_answers"`
		CreatedAt         time.Time `json:"created_at"`
	} `json:"reviews"`
}

type StudyMaterial struct {
	StudyMaterial struct {
		SubjectId       int      `json:"subject_id"`
		MeaningNote     string   `json:"meaning_note"`
		ReadingNote     string   `json:"reading_note"`
		MeaningSynonyms []string `json:"meaning_synonyms"`
	} `json:"study_material"`
}

type User struct {
	User struct {
		Preferences struct {
			LessonsAutoplayAudio       bool   `json:"lessons_autoplay_audio"`
			LessonsBatchSize           int    `json:"lessons_batch_size"`
			LessonsPresentationOrder   string `json:"lessons_presentation_order"`
			ReviewsAutoplayAudio       bool   `json:"reviews_autoplay_audio"`
			ReviewsDisplaySrsIndicator bool   `json:"reviews_display_srs_indicator"`
		} `json:"preferences"`
	} `json:"user"`
}
