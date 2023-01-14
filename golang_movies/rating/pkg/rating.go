package model

// ===================== Type-Aliases ======================

// RecordID defines a record id. Together with RecordType
// identifies unique records across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID
// identifies unique records across all types.
type RecordType string

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int

// ===================== Constant Variables ======================

// RecordTypeMovies Existing record types.
const (
	RecordTypeMovies = RecordType("movie")
)

// ===========================================

// Rating defines an individual rating created by a user for  // some record.
type Rating struct {
	RecordID   string      `json:"recordID"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userID"`
	Value      RatingValue `json:"ratingValue"`
}

// ===========================================
