package throughputrule

import (
	"context"
	"database/sql"
	"gobrm/models"
)

// CREATE TABLE throughput_rule (
//   id INT(10) NOT NULL AUTO_INCREMENT,
//   mx_domain VARCHAR(255) NOT NULL UNIQUE,
//   max_connections INT(11) NOT NULL,
//   messages_per_connection INT(11) NOT NULL,
//   connection_ttl_millis INT(11) NOT NULL,
//   PRIMARY KEY(id)
// );

const throughputRuleTable = "throughput_rule"

func getThroughputRules(db *sql.DB) (models.ThroughputRuleSlice, error) {
	ctx := context.Background()
	throughputRules, err := models.ThroughputRules().All(ctx, db)

	if err != nil {
		return nil, err
	}

	return throughputRules, nil
}

func getThroughputRule(db *sql.DB, id int) (*models.ThroughputRule, error) {
	ctx := context.Background()
	throughputRule, err := models.FindThroughputRule(ctx, db, id)

	if err != nil {
		return nil, err
	}

	return throughputRule, nil
}

func createThroughputRule(db *sql.DB) {

}

func updateThroughputRule(db *sql.DB) {

}

func deleteThroughputRule(db *sql.DB) {

}

// CREATE TABLE throughput_rule_change (
//   id INT NOT NULL AUTO_INCREMENT,
//   action VARCHAR(16) NOT NULL,
//   throughput_rule_id INT(10) NOT NULL,
//   mx_domain VARCHAR(255) NOT NULL,
//   max_connections INT(11) NOT NULL,
//   messages_per_connection INT(11) NOT NULL,
//   connection_ttl_millis INT(11) NOT NULL,
//   updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (id),
//   FOREIGN KEY (throughput_rule_id) REFERENCES throughput_rule(id) ON DELETE CASCADE
// );

const throughputRuleChangeTable = "throughput_rule_change"

func getThroughputRuleChanges(db *sql.DB) {

}

func getThroughputRuleChangesForThroughputRule(db *sql.DB) {

}
