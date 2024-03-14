package wk

import "time"

type Assignments struct {
	StartedAt time.Time `json:"started_at"`
}

type Reviews struct {
	Review ReviewBase `json:"review"`
}

type StudyMaterials struct {
	StudyMaterial struct {
		SubjectId       int      `json:"subject_id"`
		MeaningNote     string   `json:"meaning_note"`
		ReadingNote     string   `json:"reading_note"`
		MeaningSynonyms []string `json:"meaning_synonyms"`
	} `json:"study_material"`
}

type Users struct {
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
