package throughputrule

import (
	"context"
	"database/sql"
	"gobrm/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CREATE TABLE throughput_rule (
//   id INT(10) NOT NULL AUTO_INCREMENT,
//   mx_domain VARCHAR(255) NOT NULL UNIQUE,
//   max_connections INT(11) NOT NULL,
//   messages_per_connection INT(11) NOT NULL,
//   connection_ttl_millis INT(11) NOT NULL,
//   PRIMARY KEY(id)
// );

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

func createThroughputRule(db *sql.DB, throughputRule models.ThroughputRule) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = throughputRule.Insert(ctx, db, boil.Infer())

	if err != nil {
		return err
	}

	throughputRuleChange := models.ThroughputRuleChange{
		Action:                "created",
		ThroughputRuleID:      throughputRule.ID,
		MXDomain:              throughputRule.MXDomain,
		MaxConnections:        throughputRule.MaxConnections,
		MessagesPerConnection: throughputRule.MessagesPerConnection,
		ConnectionTTLMillis:   throughputRule.ConnectionTTLMillis,
	}

	err = throughputRuleChange.Insert(ctx, db, boil.Infer())

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func updateThroughputRule(db *sql.DB, throughputRule models.ThroughputRule) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	currentThroughputRule, err := models.FindThroughputRule(ctx, db, throughputRule.ID)

	if err != nil {
		return err
	}

	currentThroughputRule.MXDomain = throughputRule.MXDomain
	currentThroughputRule.MaxConnections = throughputRule.MaxConnections
	currentThroughputRule.MessagesPerConnection = throughputRule.MessagesPerConnection
	currentThroughputRule.ConnectionTTLMillis = throughputRule.ConnectionTTLMillis

	_, err = currentThroughputRule.Update(ctx, db, boil.Infer())

	if err != nil {
		return err
	}

	throughputRuleChange := models.ThroughputRuleChange{
		Action:                "updated",
		ThroughputRuleID:      currentThroughputRule.ID,
		MXDomain:              currentThroughputRule.MXDomain,
		MaxConnections:        currentThroughputRule.MaxConnections,
		MessagesPerConnection: currentThroughputRule.MessagesPerConnection,
		ConnectionTTLMillis:   currentThroughputRule.ConnectionTTLMillis,
	}

	err = throughputRuleChange.Insert(ctx, db, boil.Infer())

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func deleteThroughputRule(db *sql.DB, id int) error {
	ctx := context.Background()
	throughputRule, err := models.FindThroughputRule(ctx, db, id)
	if err != nil {
		return err
	}

	_, err = throughputRule.Delete(ctx, db)
	if err != nil {
		return err
	}
	return nil
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

func getThroughputRuleChanges(db *sql.DB) (models.ThroughputRuleChangeSlice, error) {
	ctx := context.Background()
	throughputRuleChanges, err := models.ThroughputRuleChanges().All(ctx, db)

	if err != nil {
		return nil, err
	}

	return throughputRuleChanges, nil
}

func getThroughputRuleChangesForThroughputRule(db *sql.DB, id int) (models.ThroughputRuleChangeSlice, error) {
	ctx := context.Background()
	throughputRuleChanges, err := models.ThroughputRuleChanges(qm.Where("throughput_rule_id=?", id)).All(ctx, db)

	if err != nil {
		return nil, err
	}

	return throughputRuleChanges, nil
}
