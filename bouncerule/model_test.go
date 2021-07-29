package bouncerule

import (
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBounceRules(t *testing.T) {
	log.Print("Testing model's getBounceRules")
	// Open new mock database
	db, mock, err := sqlmock.New()
	assert.NoErrorf(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "response_code", "enhanced_code", "regex", "priority", "description", "bounce_action"}).
		AddRow(1, 450, "4.7.1", "regex1", 1, "description1", "suppress").
		AddRow(2, 451, "4.7.2", "regex2", 2, "description2", "no_action")

	query := fmt.Sprintf("SELECT (.+) FROM %s", bounceRuleTable)
	mock.ExpectQuery(query).WillReturnRows(rows)

	actualBounceRules, err := getBounceRules(db)

	assert.Len(t, actualBounceRules, 2, "should have 2 bounce rules")

	expectedBounceRules := []BounceRule{
		{
			ID:           1,
			ResponseCode: 450,
			EnhancedCode: "4.7.1",
			Regex:        "regex1",
			Priority:     1,
			Description:  "description1",
			BounceAction: "suppress",
		},
		{
			ID:           2,
			ResponseCode: 451,
			EnhancedCode: "4.7.2",
			Regex:        "regex2",
			Priority:     2,
			Description:  "description2",
			BounceAction: "no_action",
		},
	}

	assert.Equalf(t, expectedBounceRules, actualBounceRules, "bounce rules do not match %v", expectedBounceRules)

	mockErr := mock.ExpectationsWereMet()
	assert.NoErrorf(t, mockErr, "did not pass expectations such as %s", mockErr)
}

func TestGetBounceRule(t *testing.T) {
	log.Print("Testing model's getBounceRule")
	db, mock, err := sqlmock.New()
	assert.NoErrorf(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "response_code", "enhanced_code", "regex", "priority", "description", "bounce_action"}).
		AddRow(1, 450, "4.7.1", "regex1", 1, "description1", "suppress")

	id := 1
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	expectedBounceRule := BounceRule{
		ID:           id,
		ResponseCode: 450,
		EnhancedCode: "4.7.1",
		Regex:        "regex1",
		Priority:     1,
		Description:  "description1",
		BounceAction: "suppress",
	}
	bounceRule := BounceRule{ID: id}
	bounceRuleErr := bounceRule.getBounceRule(db)

	assert.NoError(t, bounceRuleErr, "should not receive an error when getting bounce rule")
	assert.Equalf(t, expectedBounceRule, bounceRule, "bounce rule does not match %v", expectedBounceRule)

	mockErr := mock.ExpectationsWereMet()
	assert.NoErrorf(t, mockErr, "did not pass expectations such as %s", mockErr)
}

// TODO: BOUNCE RULE NOT FOUND ERROR

func TestCreateBounceRule(t *testing.T) {
	log.Print("Testing model's createBounceRule")
	db, mock, err := sqlmock.New()
	assert.NoErrorf(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	expectedBounceRule := BounceRule{
		ID:           1,
		ResponseCode: 450,
		EnhancedCode: "4.7.1",
		Regex:        "regex1",
		Description:  "description1",
		Priority:     1,
		BounceAction: "suppress",
	}
	bounceRule := BounceRule{
		ResponseCode: expectedBounceRule.ResponseCode,
		EnhancedCode: expectedBounceRule.EnhancedCode,
		Regex:        expectedBounceRule.Regex,
		Description:  expectedBounceRule.Description,
		Priority:     expectedBounceRule.Priority,
		BounceAction: expectedBounceRule.BounceAction,
	}

	bounceRuleErr := bounceRule.createBounceRule(db)
	assert.NoError(t, bounceRuleErr, "should not receive an error when creating bounce rule")
	assert.Equalf(t, expectedBounceRule, bounceRule, "bounce rule does not match %v", expectedBounceRule)

	mockErr := mock.ExpectationsWereMet()
	assert.NoErrorf(t, mockErr, "did not pass expectations such as %s", mockErr)

}

func TestUpdateBounceRule(t *testing.T) {
	log.Print("Testing model's updateBounceRule")
	assert.True(t, true, "True is true!")
}

func TestDeleteBounceRule(t *testing.T) {
	log.Print("Testing model's deleteBounceRule")
	assert.True(t, true, "True is true!")
}
