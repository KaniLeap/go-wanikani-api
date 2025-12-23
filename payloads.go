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
		SubjectID       int      `json:"subject_id"`
		MeaningNote     string   `json:"meaning_note"`
		ReadingNote     string   `json:"reading_note"`
		MeaningSynonyms []string `json:"meaning_synonyms"`
	} `json:"study_material"`
}

type Preferences struct {
	DefaultVoiceActorID        int    `json:"default_voice_actor_id"`
	ExtraStudyAutoplayAudio    bool   `json:"extra_study_autoplay_audio"`
	LessonsAutoplayAudio       bool   `json:"lessons_autoplay_audio"`
	LessonsBatchSize           int    `json:"lessons_batch_size"`
	LessonsPresentationOrder   string `json:"lessons_presentation_order"`
	ReviewsAutoplayAudio       bool   `json:"reviews_autoplay_audio"`
	ReviewsDisplaySrsIndicator bool   `json:"reviews_display_srs_indicator"`
	ReviewsPresentationOrder   string `json:"reviews_presentation_order"`
}

type Users struct {
	User struct {
		Preferences Preferences `json:"preferences"`
	} `json:"user"`
}
