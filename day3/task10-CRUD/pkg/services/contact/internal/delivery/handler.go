package delivery

import (
    "encoding/json"
    "net/http"

    "go/pkg/services/contact/internal/usecase"
)

type ContactHandler struct {
    useCase usecase.ContactUseCase
}

// NewContactHandler creates a new instance of ContactHandler
func NewContactHandler(useCase usecase.ContactUseCase) *ContactHandler {
    return &ContactHandler{
        useCase: useCase,
    }
}

// HandleHTTP handles HTTP requests for contact operations
func (h *ContactHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getContact(w, r)
	case http.MethodPost:
		h.createContact(w, r)
	case http.MethodPut:
		h.updateContact(w, r)
	case http.MethodDelete:
		h.deleteContact(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getContact handles HTTP GET requests to retrieve a contact
func (h *ContactHandler) getContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.useCase.GetContact()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// createContact handles HTTP POST requests to create a new contact
func (h *ContactHandler) createContact(w http.ResponseWriter, r *http.Request) {
	var contact usecase.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the contact
	err = h.useCase.CreateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// Return the created contact
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// updateContact handles HTTP PUT requests to update an existing contact
func (h *ContactHandler) updateContact(w http.ResponseWriter, r *http.Request) {
	var contact usecase.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.UpdateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

// deleteContact handles HTTP DELETE requests to delete a contact
func (h *ContactHandler) deleteContact(w http.ResponseWriter, r *http.Request) {
	err := h.useCase.DeleteContact()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

type GroupHandler struct {
    useCase usecase.GroupUseCase
}

func NewGroupHandler(useCase usecase.GroupUseCase) *GroupHandler {
    return &GroupHandler{
        useCase: useCase,
    }
}

func (h *GroupHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getGroup(w, r)
	case http.MethodPost:
		h.createGroup(w, r)
	case http.MethodPut:
		h.updateGroup(w, r)
	case http.MethodDelete:
		h.deleteGroup(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GroupHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	group, err := h.useCase.GetGroup()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	var group usecase.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateGroup(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) updateGroup(w http.ResponseWriter, r *http.Request) {
	var group usecase.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.UpdateGroup(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	err := h.useCase.DeleteGroup()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
