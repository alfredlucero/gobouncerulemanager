package bouncerule

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// CREATE TABLE bounce_rule (
//   id SMALLINT NOT NULL AUTO_INCREMENT,
//   response_code SMALLINT NOT NULL DEFAULT 0,
//   enhanced_code VARCHAR(16) NOT NULL,
//   regex VARCHAR(255) NOT NULL,
//   priority TINYINT NOT NULL DEFAULT 0,
//   description VARCHAR(255) NOT NULL,
//   bounce_action VARCHAR(255) NOT NULL,
//   PRIMARY KEY(id)
// );

type BounceRule struct {
	ID           int    `json:"id"`
	ResponseCode int    `json:"response_code"`
	EnhancedCode string `json:"enhanced_code"`
	Regex        string `json:"regex"`
	Priority     int    `json:"priority"`
	Description  string `json:"description"`
	BounceAction string `json:"bounce_action"`
}

const bounceRuleTable = "bounce_rule"

func getBounceRules(db *sql.DB) ([]BounceRule, error) {
	statement := fmt.Sprintf("SELECT id, response_code, enhanced_code, regex, priority, description, bounce_action FROM %s", bounceRuleTable)
	log.Printf("Getting bounce rules with this query: %s", statement)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bounceRules := []BounceRule{}

	for rows.Next() {
		var br BounceRule
		if err := rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction); err != nil {
			return nil, err
		}
		bounceRules = append(bounceRules, br)
	}

	return bounceRules, nil
}

func (br *BounceRule) getBounceRule(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, response_code, enhanced_code, regex, priority, description, bounce_action FROM %s WHERE id=%d", bounceRuleTable, br.ID)
	log.Printf("Getting bounce rule with this query: %s", statement)
	return db.QueryRow(statement).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
}

func (br *BounceRule) createBounceRule(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO %s (response_code, enhanced_code, regex, priority, description, bounce_action) VALUES (%d, '%s', '%s', %d, '%s', '%s')", bounceRuleTable, br.ResponseCode, br.EnhancedCode, br.Regex, br.Priority, br.Description, br.BounceAction)
	log.Printf("Creating bounce rule with this query: %s", statement)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&br.ID)

	if err != nil {
		return err
	}

	return nil
}

func (br *BounceRule) updateBounceRule(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE %s SET response_code=%d, enhanced_code='%s', regex='%s', priority=%d, description='%s', bounce_action='%s' WHERE id=%d", bounceRuleTable, br.ResponseCode, br.EnhancedCode, br.Regex, br.Priority, br.Description, br.BounceAction, br.ID)
	log.Printf("Updating bounce rule with this query: %s", statement)
	_, err := db.Exec(statement)
	return err
}

func (br *BounceRule) deleteBounceRule(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM %s WHERE id=%d", bounceRuleTable, br.ID)
	log.Printf("Deleting bounce rule with this query: %s", statement)
	_, err := db.Exec(statement)
	return err
}

// CREATE TABLE bounce_rule_change (
//   id SMALLINT NOT NULL AUTO_INCREMENT,
//   action VARCHAR(16) NOT NULL,
//   bounce_rule_id SMALLINT NOT NULL,
//   response_code SMALLINT NOT NULL DEFAULT 0,
//   enhanced_code VARCHAR(16) NOT NULL,
//   regex VARCHAR(255) NOT NULL,
//   priority TINYINT NOT NULL DEFAULT 0,
//   description VARCHAR(255) NOT NULL,
//   bounce_action VARCHAR(255) NOT NULL,
//   updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (id)
// );

type BounceRuleChange struct {
	ID           int       `json:"id"`
	Action       string    `json:"action"`
	BounceRuleID int       `json:"bounce_rule_id"`
	ResponseCode int       `json:"response_code"`
	EnhancedCode string    `json:"enhanced_code"`
	Regex        string    `json:"regex"`
	Priority     int       `json:"priority"`
	Description  string    `json:"description"`
	BounceAction string    `json:"bounce_action"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const bounceRuleChangeTable = "bounce_rule_change"

func getBounceRuleChanges(db *sql.DB) ([]BounceRuleChange, error) {
	statement := fmt.Sprintf("SELECT id, action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action, updated_at FROM %s", bounceRuleChangeTable)
	log.Printf("Getting bounce rule changes with this query: %s", statement)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bounceRuleChanges := []BounceRuleChange{}

	for rows.Next() {
		var brc BounceRuleChange
		if err := rows.Scan(&brc.ID, &brc.Action, &brc.BounceRuleID, &brc.ResponseCode, &brc.EnhancedCode, &brc.Regex, &brc.Priority, &brc.Description, &brc.BounceAction, &brc.UpdatedAt); err != nil {
			return nil, err
		}
		bounceRuleChanges = append(bounceRuleChanges, brc)
	}

	return bounceRuleChanges, nil
}

func getBounceRuleChangesForBounceRule(db *sql.DB, bounceRuleID int) ([]BounceRuleChange, error) {
	statement := fmt.Sprintf("SELECT id, action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action, updated_at FROM %s WHERE bounce_rule_id = %d", bounceRuleChangeTable, bounceRuleID)
	log.Printf("Getting bounce rule changes for bounce rule with this query: %s", statement)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bounceRuleChanges := []BounceRuleChange{}

	for rows.Next() {
		var brc BounceRuleChange
		if err := rows.Scan(&brc.ID, &brc.Action, &brc.BounceRuleID, &brc.ResponseCode, &brc.EnhancedCode, &brc.Regex, &brc.Priority, &brc.Description, &brc.BounceAction, &brc.UpdatedAt); err != nil {
			return nil, err
		}
		bounceRuleChanges = append(bounceRuleChanges, brc)
	}

	return bounceRuleChanges, nil
}
