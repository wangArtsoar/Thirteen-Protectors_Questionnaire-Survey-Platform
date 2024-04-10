package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestEditServerById tests the EditServerById method of the ServerRepo
func TestEditServerById(t *testing.T) {
	// Create a mock session using xorm.NewSession
	session := persistence.NewXorm().NewSession()
	defer session.Close()

	// Create a mock server with some initial values
	server := &models.Server{
		Id:         1,
		Name:       "Test Server",
		CreateAt:   time.Now(),
		OwnerId:    "e7b6780d-7e2b-4859-acc9-caad91bb62e9",
		OwnerEmail: "test@example.com",
	}
	labels, err := json.Marshal([]string{"test", "server"})
	server.Labels = labels

	// Insert the mock server into the session
	_, err = session.Insert(server)
	assert.NoError(t, err)

	// Create a new ServerRepo instance
	repo := NewServerRepo()

	labels, err = json.Marshal([]string{"test", "server", "updated"})
	// Create a new server with some updated values
	newServer := &models.Server{
		Labels:   labels,
		UpdateAt: time.Now(),
	}

	// Call the EditServerById method with the mock session, the server ID, and the new server
	err = repo.EditServerById(session, server.Id, newServer)
	assert.NoError(t, err)

	// Query the updated server from the session
	updatedServer := &models.Server{}
	_, err = session.ID(server.Id).Get(updatedServer)
	assert.NoError(t, err)

	var newLabels, updatedLabels []string
	_ = json.Unmarshal(newServer.Labels, &newLabels)
	_ = json.Unmarshal(updatedServer.Labels, &updatedLabels)
	// Assert that the updated server has the expected values
	assert.Equal(t, newLabels, updatedLabels)
	assert.Equal(t, newServer.UpdateAt.Format(time.DateOnly), updatedServer.UpdateAt.Format(time.DateOnly))

	sql := `delete from server where id = ?`
	_, err = session.Exec(sql, server.Id)
	assert.NoError(t, err)
}
